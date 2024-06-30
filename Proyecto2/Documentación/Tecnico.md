# MANUAL TECNICO

### Carpetas
- Consumer
- Documentacion
- gRPC
- locust
- Plantillas
- redis-rust

### Consumer

**Conn.go**

Importaciónes:

```go
import (
	"Consumer/Instance"
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)
```

son las importaciónes que se utilizan en esta plantilla, y a continuacion veremos las escructuras:

```go
type MongoInstance struct {
	Client *mongo.Client
	Db     *mongo.Database
}
```

'MongoInstance' es una estructura que encapsula las referencias al cliente MongoDB (Client) y a la base de datos específica (Db) a la que se conecta la aplicación.

```go
func Connect() error {
	server := "mongo-service"
	port := "27017"
	dbName := "DB2"
	mongoUri := "mongodb://" + server + ":" + port

	client, err := mongo.NewClient(options.Client().ApplyURI(mongoUri))
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	err = client.Connect(ctx)
	db := client.Database(dbName)

	if err != nil {
		log.Fatal(err)
	}

	Instance.Mg = Instance.MongoInstance(MongoInstance{
		Client: client,
		Db:     db,
	})

	return nil
}
```

Connect() es una función que realiza la conexión a la instancia de MongoDB especificada (mongo-service:27017) y configura la estructura MongoInstance con el cliente y la base de datos.

Variables de conexión:

- server: Nombre del servicio de MongoDB (mongo-service).
- port: Puerto de MongoDB (27017).
- dbName: Nombre de la base de datos (DB2).
- mongoUri: URI de conexión a MongoDB construido usando server y port.

Creación del cliente MongoDB:

- client, err := mongo.NewClient(options.Client().ApplyURI(mongoUri)): Crea un nuevo cliente MongoDB utilizando la URI construida.
- ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second): Crea un contexto con un tiempo de espera de 30 segundos para la operación de conexión.
- defer cancel(): Garantiza que se cancele el contexto al final de la función.
- err = client.Connect(ctx): Intenta conectar el cliente MongoDB al servidor usando el contexto creado.
- db := client.Database(dbName): Obtiene una referencia a la base de datos especificada (DB2).

Manejo de errores:

- if err != nil { log.Fatal(err) }: Si hay algún error en la conexión, se registra el error y se detiene la ejecución del programa.

Configuración de MongoInstance:

- Instance.Mg = Instance.MongoInstance(MongoInstance{ Client: - client, Db: db }): Configura la estructura MongoInstance global (Instance.Mg) con el cliente MongoDB (client) y la base de datos (db) obtenidos.

Retorno de error:
- return nil: Retorna nil si la conexión y configuración fueron exitosas, indicando que no hubo errores.

**Instance.go**

```go
package Instance

import "go.mongodb.org/mongo-driver/mongo"

type MongoInstance struct {
	Client *mongo.Client
	Db     *mongo.Database
}

var Mg MongoInstance

```

Descripción: MongoInstance es una estructura que encapsula dos campos:
- Client: Un puntero a mongo.Client, que representa el cliente de MongoDB.
- Db: Un puntero a mongo.Database, que representa una base de datos específica dentro de MongoDB.

**Data.go**

```go
package model

import "time"

type Data struct {
	Texto string    `json:"texto"`
	Pais  string    `json:"pais"`
	Fecha time.Time `json:"fecha"`
}

```

Descripción: Data es una estructura que encapsula tres campos de datos:

- Texto: Un campo de tipo string que almacena texto.
- Pais: Un campo de tipo string que almacena el nombre del país.
- Fecha: Un campo de tipo time.Time que almacena una fecha y hora.

**Dockerfile**

Es un Dockerfile que se utiliza para construir y ejecutar una aplicación escrita en Go dentro de un contenedor Docker, utilizando un enfoque de multi-etapa (multi-stage build).

```Dockerfile
# Start by building the application.
FROM golang:1.20-buster as build

WORKDIR /go/src/app
ADD . /go/src/app

RUN go get -d -v ./...

RUN go build -o /go/bin/app

# Now copy it into our base image.
FROM gcr.io/distroless/base-debian12
COPY --from=build /go/bin/app /
CMD ["/app"]
```

**main**

```go
package main

func processEvent(event []byte) {
	var data model.Data
	err := json.Unmarshal(event, &data)
	if err != nil {
		log.Fatal(err)
	}

	// Conectar a la base de datos si no está conectado
	if Instance.Mg.Client == nil {
		if err := Database.Connect(); err != nil {
			log.Fatal("Error en", err)
		}
	}

	// Establecer los campos de fecha y hora
	data.Fecha = time.Now()

	collection := Instance.Mg.Db.Collection("register")
	_, err = collection.InsertOne(context.TODO(), data)
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	topic := "mytopic"
	r := kafka.NewReader(kafka.ReaderConfig{
		Brokers:     []string{"my-cluster-kafka-bootstrap.kafka.svc.cluster.local:9092"},
		Topic:       topic,
		Partition:   0,
		MinBytes:    10e3,
		MaxBytes:    10e6,
		StartOffset: kafka.LastOffset,
		GroupID:     uuid.New().String(),
	})

	for {
		m, err := r.ReadMessage(context.Background())
		if err != nil {
			log.Printf("Error reading message: %v", err)
		}
		fmt.Printf("Message at offset %d: %s = %s\n", m.Offset, string(m.Key), string(m.Value))

		processEvent(m.Value)

		err = r.CommitMessages(context.Background(), m)
		if err != nil {
			log.Printf("Error committing message: %v", err)
		}
	}
}
```

Descripción de ***ProcessEvent***: Es una funcion que procesa un evento recibido desde kafka

Descripción de ***Main***: La funcion main es la entrada principal del programa y configura un consumidor de kafka, contiene un bucle infinito que lle llos mensajes de kafka.

Este código muestra cómo consumir mensajes desde un cluster de Kafka utilizando Go, procesar estos mensajes (decodificarlos desde JSON y almacenarlos en MongoDB), y manejar errores de manera adecuada. Es esencialmente un ejemplo de integración entre Kafka y MongoDB utilizando Go, mostrando buenas prácticas como el uso de contextos, manejo de errores y modularización del código.

### gRPC
**Main.go para clientGRPC**

Este programa es un servicio web en Go que recibe datos a través de una solicitud HTTP POST y los envía a dos destinos diferentes: un servidor GRPC y un servicio Rust. A continuación, se detalla el código:

```go
package main

import (
	"bytes"
	pb "clientGRPC/client"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/gofiber/fiber/v2"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var ctx = context.Background()

type Data struct {
	Texto string
	Pais  string
}

func sendToRust(data *Data) {
	jsonData, err := json.Marshal(data)
	if err != nil {
		log.Fatal(err)
	}

	res, err := http.Post("http://rust-redis-service:8000/set", "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		log.Fatal(err)
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(res.Body)

	if res.StatusCode != http.StatusOK {
		log.Fatalf("status code error: %d %s", res.StatusCode, res.Status)
	}
}

func sendData(c *fiber.Ctx) error {
	/* API REST */
	var data map[string]string
	e := c.BodyParser(&data)
	if e != nil {
		return e
	}

	tweet := Data{
		Texto: data["texto"],
		Pais:  data["pais"],
	}

	go sendGrpcServer(tweet)
	go sendToRust(&tweet)

	return nil
}

func sendGrpcServer(tweet Data) {
	/* GRPC Client */
	conn, err := grpc.Dial("localhost:3001", grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}

	cl := pb.NewGetInfoClient(conn)
	defer func(conn *grpc.ClientConn) {
		err := conn.Close()
		if err != nil {
			log.Fatalf("could not close connection: %v", err)
		}
	}(conn)

	ret, err := cl.ReturnInfo(ctx, &pb.RequestId{
		Texto: tweet.Texto,
		Pais:  tweet.Pais,
	})
	if err != nil {
		log.Fatal(err)
	} else {
		fmt.Println("Respuesta del servidor ", ret)
	}
}

func main() {
	app := fiber.New()

	app.Post("/insert", sendData)

	err := app.Listen(":3000")
	if err != nil {
		log.Fatal(err)
	}
}

```

1. Se importan los paquetes necesarios para manejar HTTP, GRPC, JSON, logs, contexto y Fiber (un framework web en Go).

2. Se crea un contexto global que se utilizará para las solicitudes GRPC.

3. Se define una estructura Data que contiene dos campos: Texto y Pais.

4. Función sendToRust: Esta función convierte la estructura Data a JSON y la envía a un servicio Rust a través de una solicitud HTTP POST. Si ocurre algún error durante la serialización o la solicitud HTTP, se registra el error y el programa termina. También se verifica el código de estado de la respuesta para asegurarse de que la solicitud fue exitosa.


5. Función sendData: Esta función maneja la solicitud HTTP POST. Primero, analiza el cuerpo de la solicitud para extraer los datos. Luego, crea una instancia de Data con los datos recibidos. Finalmente, llama a sendGrpcServer y sendToRust en goroutines (ejecutándolas de forma concurrente).

6. Función sendGrpcServer: Esta función establece una conexión GRPC con un servidor en localhost:3001. Crea un cliente GRPC y llama al método ReturnInfo del servidor, pasando los datos recibidos en la solicitud. La respuesta del servidor se imprime en la consola.

7. Función main: La función main inicia una nueva aplicación Fiber, define una ruta POST /insert que maneja sendData, y comienza a escuchar en el puerto 3000.

**Dockerfile para ClientGRPC**

```Dockerfile
FROM golang:alpine AS builder

WORKDIR /go/src/app

COPY . ./

RUN go get -d -v
RUN go build -o /go/bin/app

FROM alpine:latest

RUN apk --no-cache add ca-certificates

WORKDIR /app
COPY --from=builder /go/bin/app /app

ENTRYPOINT ["/app/app"]

EXPOSE 3000
```

Primera Etapa:

- Utiliza una imagen base de Go para construir la aplicación.
- Establece un directorio de trabajo.
- Copia el código fuente.
- Descarga las dependencias y compila el código fuente en un binario.

Segunda Etapa:

- Utiliza una imagen base ligera de Alpine Linux para ejecutar la aplicación.
- Instala certificados CA.
- Establece un directorio de trabajo.
- Copia el binario desde la primera etapa.
- Establece el binario como el punto de entrada del contenedor.
- Expone el puerto 3000.

Este enfoque de construcción en dos etapas permite generar imágenes de Docker más pequeñas y eficientes, ya que la imagen final solo contiene el binario compilado y las dependencias necesarias para ejecutarlo, sin incluir las herramientas de compilación y el código fuente.

**main.go para serverGRPC**

```go
package main

import (
	"context"
	"fmt"
	"net"
	"serverGRPC/kafka"
	"serverGRPC/model"
	pb "serverGRPC/server"

	"google.golang.org/grpc"
)

type server struct {
	pb.UnimplementedGetInfoServer
}

func (s *server) ReturnInfo(ctx context.Context, in *pb.RequestId) (*pb.ReplyInfo, error) {
	tweet := model.Data{
		Texto: in.GetTexto(),
		Pais:  in.GetPais(),
	}

	fmt.Println(tweet)

	kafka.Produce(tweet)

	return &pb.ReplyInfo{Info: "Twitter recibido."}, nil
}

func main() {
	listen, err := net.Listen("tcp", ":3001")
	if err != nil {
		panic(err)
	}
	s := grpc.NewServer()
	pb.RegisterGetInfoServer(s, &server{})

	if err := s.Serve(listen); err != nil {
		panic(err)
	}
}
```

Importaciones
Aquí estamos importando varios paquetes:

- context, fmt, y net son paquetes estándar de Go.
- serverGRPC/kafka y serverGRPC/model son paquetes personalizados, presumiblemente definidos en tu proyecto.
- pb "serverGRPC/server" es el paquete generado por Protocol Buffers para el servicio gRPC.
- google.golang.org/grpc es el paquete de Go para trabajar con gRPC.

Definición del Servidor
Definimos una estructura server que implementa el servicio gRPC GetInfoServer. La estructura incluye una composición de pb.UnimplementedGetInfoServer para cumplir con la interfaz gRPC.

Implementación del Método gRPC

El método ReturnInfo es la implementación del método gRPC definido en el archivo Protocol Buffers.

1. Recibe un contexto ctx y un mensaje de solicitud in de tipo pb.RequestId.
2. Crea una instancia de model.Data con los datos del mensaje de solicitud (Texto y Pais).
3. Imprime la estructura tweet en la consola.
4. Envía el tweet a Kafka usando la función kafka.Produce.
5. Retorna una respuesta pb.ReplyInfo con un mensaje de confirmación.

Función Principal

1. La función main crea un listener en el puerto 3001 para escuchar conexiones TCP.
2. Si hay un error al crear el listener, el programa se detiene con panic.
3. Se crea un nuevo servidor gRPC.
4. Se registra el servicio GetInfoServer en el servidor gRPC.
5. El servidor gRPC empieza a servir y aceptar conexiones en el listener. Si hay un error al servir, el programa se detiene con panic.


**Kafka**

```go
package kafka

import (
	"context"
	"encoding/json"
	"log"
	"serverGRPC/model"
	"time"

	"github.com/segmentio/kafka-go"
)

func Produce(value model.Data) {
	topic := "mytopic"
	partition := 0

	conn, err := kafka.DialLeader(context.Background(), "tcp", "my-cluster-kafka-bootstrap.kafka.svc.cluster.local:9092", topic, partition)
	if err != nil {
		panic(err)
	}

	valueBytes, err := json.Marshal(&value)
	if err != nil {
		panic(err)
	}
	err = conn.SetWriteDeadline(time.Now().Add(10 * time.Second))
	if err != nil {
		return
	}
	_, err = conn.WriteMessages(
		kafka.Message{
			Value: valueBytes,
		})

	if err != nil {
		panic(err)
	}

	if err := conn.Close(); err != nil {
		panic(err)
	}

	log.Printf("Produced message to topic %s at partition %d", topic, partition)
}
```

Importaciones:

- context, encoding/json, log, y time son paquetes estándar de Go.
- serverGRPC/model es un paquete personalizado que define la estructura de datos Data.
- github.com/segmentio/kafka-go es una librería para interactuar con Apache Kafka en Go.


Función Produce:
- Variables topic y partition: Define el tópico y la partición de Kafka a los que se enviará el mensaje.
- Conexión a Kafka: Establece una conexión al líder de la partición del tópico usando kafka.DialLeader. Se conecta al servidor Kafka en la dirección especificada.
- Serialización del Mensaje: Convierte la estructura value de tipo model.Data a un formato JSON para poder enviarlo como mensaje.
- Definición de Tiempo Límite: Establece un tiempo límite de 10 segundos para escribir el mensaje en Kafka.
- Escritura del Mensaje: Escribe el mensaje en Kafka. Si hay un error, lanza un pánico.
- Cierre de la Conexión: Cierra la conexión con Kafka. Si hay un error, lanza un pánico.
- Registro del Mensaje Producido: Imprime en el log un mensaje indicando que se ha producido un mensaje al tópico y partición especificados.

**Dockerfile de serverGRPC**
```Dockerfile
FROM golang:alpine AS builder

WORKDIR /go/src/app

COPY . ./

RUN go get -d -v
RUN go build -o /go/bin/app

FROM alpine:latest

RUN apk --no-cache add ca-certificates

WORKDIR /app
COPY --from=builder /go/bin/app /app

ENTRYPOINT ["/app/app"]

EXPOSE 3001
```

Este Dockerfile construye una imagen Docker optimizada para una aplicación Go. La construcción se realiza en dos etapas:

Builder Stage: Utiliza una imagen de Go para compilar la aplicación y descargar dependencias.
Final Stage: Utiliza una imagen minimalista de Alpine Linux, copia el binario compilado desde la etapa de construcción y establece el entorno para ejecutar la aplicación.
Al utilizar dos etapas, se mantiene la imagen final ligera, ya que solo incluye el binario necesario y las dependencias mínimas.

### Plantillas

**grafaba.yaml**

```yaml
apiVersion: apps/v1
kind: Deployment
metadata:
  name: grafana
  namespace: so1jun2024
spec:
  replicas: 1
  selector:
    matchLabels:
      app: grafana
  template:
    metadata:
      name: grafana
      labels:
        app: grafana
    spec:
      containers:
      - name: grafana
        image: grafana/grafana:8.4.4
        ports:
        - name: grafana
          containerPort: 3010
        resources:
          limits:
            memory: "1Gi"
            cpu: "1000m"
          requests:
            memory: 500M
            cpu: "500m"
```
Especificaciónes Generales
1. apiVersion: apps/v1: Especifica la versión de la API de Kubernetes que se está utilizando.
2. kind: Deployment: Define que este recurso es un Deployment, que gestiona una aplicación replicada.
3. metadata: Contiene datos que identifican el Deployment:
	- name: grafana: Nombre del Deployment.
	- namespace: so1jun2024: El namespace en el que se desplegará el Deployment.

Especificaciones del Deployment
- spec: Define la especificación del Deployment.
	- replicas: 1: Especifica que debe haber 1 réplica del pod ejecutándose.
	- selector: Define cómo identificar los pods que forman parte de este Deployment.
		- matchLabels: Utiliza etiquetas para seleccionar los pods.
			- app: grafana: Solo los pods con esta etiqueta son seleccionados.
	- template: Define el template para los pods que serán creados.
		- metadata: Metadatos del pod.
			- name: grafana: Nombre del pod.
			- labels: Etiquetas aplicadas al pod.
				- app: grafana: Etiqueta aplicada.
		- spec: Especificación de los contenedores dentro del pod.
			- containers: Lista de contenedores en el pod.
				- name: grafana: Nombre del contenedor.
				- image: grafana/grafana:8.4.4: Imagen de Docker que se usará para el contenedor.
			- ports: Lista de puertos expuestos por el contenedor.
				- name: grafana: Nombre del puerto.
				- containerPort: 3010: Puerto expuesto por el contenedor.
			- resources: Define los recursos de CPU y memoria para el contenedor.
				- limits: Límite máximo de recursos que puede usar el contenedor.
				- memory: "1Gi": Límite de 1 Gigabyte de memoria.
				- cpu: "1000m": Límite de 1000 milicores de CPU (equivalente a 1 CPU core).
			-requests: Recursos solicitados por el contenedor al programarse.
				-memory: 500M: Solicitud de 500 Megabytes de memoria.
				- cpu: "500m": Solicitud de 500 milicores de CPU (equivalente a 0.5 CPU cores).

**grpc.yaml**
```yaml
apiVersion: apps/v1
kind: Deployment
metadata:
  name: grpc-producer
  namespace: project
spec:
  replicas: 1
  selector:
    matchLabels:
      app: grpc-producer
  template:
    metadata:
      labels:
        app: grpc-producer
    spec:
      containers:
      - name: grpc-client
        image: kesm12/clientgrpc
        ports:
        - containerPort: 3000
        resources:
          limits:
            cpu: "0.4"
            memory: "500Mi"
      - name: grpc-server
        image: kesm12/servergrpc
        ports:
        - containerPort: 3001
        resources:
          limits:
            cpu: "0.4"
            memory: "500Mi"
---
apiVersion: v1
kind: Service
metadata:
  name: grpc-client-service
  namespace: project
spec:
  selector:
    app: grpc-producer
  ports:
    - protocol: TCP
      port: 3000
      targetPort: 3000
  type: ClusterIP
```
Deployment
1. Deployment: Despliega una aplicación llamada grpc-producer en el namespace project.
2. Replicas: Ejecuta 1 réplica del pod.
3. Containers:
	- grpc-client: Usa la imagen kesm12/clientgrpc, expone el puerto 3000, y tiene límites de recursos (CPU: 0.4, Memoria: 500Mi).
	- grpc-server: Usa la imagen kesm12/servergrpc, expone el puerto 3001, y tiene límites de recursos (CPU: 0.4, Memoria: 500Mi).

Service
- Service: Define un servicio llamado grpc-client-service en el namespace project.
- Selector: Selecciona los pods con la etiqueta app: grpc-producer.
- Ports: Mapea el puerto 3000 del servicio al puerto 3000 del pod.
- Type: ClusterIP, lo que hace que el servicio sea accesible solo dentro del clúster.

**ingress.yaml**
```yaml
apiVersion: networking.k8s.io/v1
kind: Ingress
metadata:
  name: so1-proyecto2-ingress
  namespace: so1jun2024
spec:
  ingressClassName: nginx
  rules:
    - host: 34.122.71.194.nip.io
      http:
        paths:
          - pathType: Prefix
            backend:
              service:
                name: grpc-client-service
                port:
                  number: 3000
            path: /insert


            #34.118.234.94   130.211.209.138
```
Este manifiesto de Kubernetes define un recurso Ingress que utiliza NGINX Ingress Controller para enrutar el tráfico HTTP. Las solicitudes dirigidas a 34.122.71.194.nip.io/insert serán enviadas al servicio grpc-client-service en el puerto 3000 dentro del namespace so1jun2024. El uso de pathType: Prefix asegura que cualquier URL que comience con /insert coincida y se redirija adecuadamente.


**redis.yaml**

```yaml
apiVersion: apps/v1
kind: Deployment
metadata:
  labels:
    run: redis
  name: redis
  namespace: project
spec:
  replicas: 1
  selector:
    matchLabels:
      run: redis
  template:
    metadata:
      labels:
        run: redis
    spec:
      containers:
      - name: redis
        image: redis:6.2
        ports:
        - containerPort: 6379
        resources:
          limits:
            cpu: "0.2"
            memory: "128Mi"
---
apiVersion: v1
kind: Service
metadata:
  labels:
    run: redis
  name: redis
  namespace: project
spec:
  ports:
  - port: 6379
    protocol: TCP
    targetPort: 6379
  selector:
    run: redis
  type: ClusterIP

```

Deployment

1. Deployment: Despliega una instancia de Redis en el namespace project.
2. Replicas: Ejecuta 1 réplica del pod.
3. Container:
	- name: redis
	- image: redis:6.2
	- ports: Expone el puerto 6379.
	- resources: Limita el uso de recursos a 0.2 CPU y 128Mi de memoria

Service
1. Service: Define un servicio llamado redis en el namespace project.
2. Ports:
	- port: 6379 (externo)
	- targetPort: 6379 (interno)
	- protocol: TCP
3. Selector: Selecciona los pods con la etiqueta run: redis.
4. Type: ClusterIP, lo que hace que el servicio sea accesible solo dentro del clúster.



### Redis
```rs
#[macro_use] extern crate rocket;

use rocket::serde::json::Json;
use rocket::serde::{Deserialize, Serialize};
use redis::Commands;

#[derive(Deserialize, Serialize)]
struct Data {
    Pais: String,
    Texto: String
}

#[post("/set", format = "json", data = "<data>")]
async fn set_data(data: Json<Data>) -> Result<&'static str, &'static str> {
    // Crear cliente de redis
    let client = redis::Client::open("redis://redis:6379/")
        .map_err(|_| "Failed to create Redis client")?;

    // Conexion a redis
    let mut con = client.get_connection()
        .map_err(|_| "Failed to connect to Redis")?;
    
    // Insertar hash en redis
    let _: () = con.hincr(&data.Pais, &data.Texto, 1)
    .map_err(|_| "Failed to set data in Redis")?;

    // Insertar hash en redis
    let _: () = con.hincr("paises", &data.Pais, 1)
        .map_err(|_| "Failed to set data in Redis")?;

    // Insertar hash en redis
    let _: () = con.hincr("mensajes", &data.Texto, 1)
        .map_err(|_| "Failed to set data in Redis")?;

    // Incrementar contador total de mensajes
    let _: () = con.incr("cantMensajes", 1)
        .map_err(|_| "Failed to increment total messages in Redis")?;

    Ok("Data set")
}

#[launch]
fn rocket() -> _ {
    rocket::build().mount("/", routes![set_data])
}
```

Importaciones y Configuración
- rocket y rocket::serde se utilizan para manejar la serialización y deserialización de JSON.
- redis::Commands se usa para interactuar con Redis.

Definición de la Estructura de Datos
- Estructura Data: Representa los datos que se recibirán y enviarán en formato JSON, con dos campos: Pais y Texto.

Handler para la Ruta POST

- Ruta /set: Define un endpoint POST que acepta datos en formato JSON.
- Conexión a Redis:
	- Crea un cliente Redis y establece una conexión.
	- Incrementa los contadores en Redis:
		- Por país (data.Pais).
		- Por texto (data.Texto).
		- Total de mensajes (cantMensajes).
- Manejo de Errores: Si alguna operación falla, retorna un mensaje de error.

Lanzamiento de la Aplicación
- Función rocket: Configura y lanza la aplicación Rocket, montando la ruta /set