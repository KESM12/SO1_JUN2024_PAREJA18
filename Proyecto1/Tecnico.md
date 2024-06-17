# MANUAL TECNICO

## kvm
KVM (Kernel-based Virtual Machine) es una tecnología de virtualización de código abierto integrada en el kernel de Linux. Permite ejecutar múltiples sistemas operativos invitados (máquinas virtuales) en un único ordenador físico.

![kvm](data:image/jpeg;base64,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)


 Esto ofrece varias ventajas, incluyendo:

**Eficiencia:** KVM aprovecha las capacidades de virtualización del hardware del sistema, lo que le permite ofrecer un alto rendimiento y una baja latencia para las máquinas virtuales.

**Escalabilidad:** KVM puede ejecutar un gran número de máquinas virtuales en un solo ordenador, lo que lo hace ideal para entornos de servidor y computación en la nube.

**Aislamiento:** Las máquinas virtuales KVM están aisladas entre sí y del sistema operativo host, lo que proporciona seguridad y protección contra fallos en una máquina virtual que afecten a las demás.

**Portabilidad:** Las máquinas virtuales KVM son portátiles y se pueden mover fácilmente entre diferentes ordenadores físicos.


**KVM es una tecnología de virtualización potente y versátil que ofrece una amplia gama de beneficios para la virtualización de servidores, escritorios y aplicaciones.**

## Modulos

Este código define un módulo del kernel de Linux que recopila información sobre la CPU y los procesos del sistema, y la muestra en un archivo accesible a través de procfs

```c
static int escribir_archivo(struct seq_file *archivo, void *v) {
    for_each_process(cpu) {
        seq_printf(archivo, "PID%d", cpu->pid);
        seq_printf(archivo, ",");
        seq_printf(archivo, "%s", cpu->comm);
        seq_printf(archivo, ",");
        seq_printf(archivo, "%lu", cpu->__state);
        seq_printf(archivo, ",");

        if (cpu->mm) {
            rss = get_mm_rss(cpu->mm) << PAGE_SHIFT;
            seq_printf(archivo, "%lu", rss);
        } else {
            seq_printf(archivo, "%s", "");
        }
        seq_printf(archivo, ",");

        seq_printf(archivo, "%d", cpu->cred->user->uid);
        seq_printf(archivo, ",");

        list_for_each(lstProcess, &(cpu->children)) {
            child = list_entry(lstProcess, struct task_struct, sibling);
            seq_printf(archivo, "Child:%d", child->pid);
            seq_printf(archivo, ".");
            seq_printf(archivo, "%s", child->comm);
            seq_printf(archivo, ".");
            seq_printf(archivo, "%d", child->__state);
            seq_printf(archivo, ".");

             if (child->mm) {
                rss = get_mm_rss(child->mm) << PAGE_SHIFT;
                seq_printf(archivo, "%lu", rss);
            } else {
                seq_printf(archivo, "%s", "");
            }
            seq_printf(archivo, ".");

            seq_printf(archivo, "%d", child->cred->user->uid);
        }
    }

    return 0;
}

//Funcion que se ejecutara cada vez que se lea el archivo con el comando CAT
static int al_abrir(struct inode *inode, struct file *file)
{
    return single_open(file, escribir_archivo, NULL);
}

//Si el kernel es 5.6 o mayor se usa la estructura proc_ops
static struct proc_ops operaciones =
{
    .proc_open = al_abrir,
    .proc_read = seq_read
};

//Funcion a ejecuta al insertar el modulo en el kernel con insmod
static int _insert(void)
{
    proc_create("cpu_so1_1s2024", 0, NULL, &operaciones);
    printk(KERN_INFO "Laboratorio Sistemas Operativos 1\n");
    return 0;
}

//Funcion a ejecuta al remover el modulo del kernel con rmmod
static void _remove(void)
{
    remove_proc_entry("cpu_so1_1s2024", NULL);
    printk(KERN_INFO "Laboratorio Sistemas Operativos 1\n");
}

module_init(_insert);
module_exit(_remove);

```
**Este código del módulo del kernel de Linux permite crear un archivo en /proc que muestra información detallada sobre la CPU, los procesos del sistema y sus procesos hijos, incluyendo el PID, nombre, estado, memoria residente y usuario propietario.**

Ahora el de la ram:
```c
static void init_meminfo(void) {
    si_meminfo(&si);
}

MODULE_LICENSE("GPL");
MODULE_DESCRIPTION("Modulo de RAM, Laboratorio Sistemas Operativos 1");
MODULE_AUTHOR("Grupo16");

//Funcion que se ejecutara cada vez que se lea el archivo con el comando CAT
static int escribir_archivo(struct seq_file *archivo, void *v)
{
    init_meminfo();
    //Se escribe en el archivo la informacion de la memoria RAM en MB
    seq_printf(archivo, "%lu,%lu,%lu\n", (si.freeram * si.mem_unit) / (1024 * 1024), ((si.totalram - si.freeram) * si.mem_unit) / (1024 * 1024), (si.totalram * si.mem_unit) / (1024 * 1024));
    return 0;
}

//Funcion que se ejecutara cada vez que se lea el archivo con el comando CAT
static int al_abrir(struct inode *inode, struct file *file)
{
    return single_open(file, escribir_archivo, NULL);
}

//Si el kernel es 5.6 o mayor se usa la estructura proc_ops
static struct proc_ops operaciones =
{
    .proc_open = al_abrir,
    .proc_read = seq_read
};

//Funcion a ejecuta al insertar el modulo en el kernel con insmod
static int _insert(void)
{
    proc_create("ram_so1_jun2024", 0, NULL, &operaciones);
    printk(KERN_INFO "Laboratorio Sistemas Operativos 1\n");
    return 0;
}

//Funcion a ejecuta al remover el modulo del kernel con rmmod
static void _remove(void)
{
    remove_proc_entry("ram_so1_jun2024", NULL);
    printk(KERN_INFO "Laboratorio Sistemas Operativos 1\n");
}

module_init(_insert);
module_exit(_remove);
```

## Backend 

### DataController.go

Este código define un conjunto de funciones en el lenguaje Go para interactuar con una base de datos MongoDB, como parte de un sistema de monitoreo de recursos del sistema.

Las funciones para insertar RAM, CPU Y PROCESOS es la siguiente

```go
func InsertRam(nameCol string, Total int, Enuso int, Libre int, Porcentaje int) {
	collection := Instance.Mg.Db.Collection(nameCol)
	doc := Model.Ram{Total: Total, En_uso: Enuso, Libre: Libre, Porcentaje: Porcentaje}
	_, err := collection.InsertOne(context.TODO(), doc)
	if err != nil {
		log.Fatal(err)
	}
}
```

Los parametros que requieren las funciones varia segun sea el proceso, una explicación de los parametros del metodo InsertaRam son los siguientes:

Parámetros:
- nameCol (string): Nombre de la colección en la que se insertarán los datos (colección de RAM).
- Total (int): Capacidad total de RAM del sistema en unidades enteras (por ejemplo, bytes).
- Enuso (int): Cantidad de RAM actualmente en uso en unidades enteras.
- Libre (int): Cantidad de RAM disponible en unidades enteras.
- Porcentaje (int): Porcentaje de uso de RAM calculado (0-100).


Elimina (borra) una colección específica de MongoDB.

```go
func ResetCollection(nameCol string) error {
	collection := Instance.Mg.Db.Collection(nameCol)
	err := collection.Drop(context.TODO())
	if err != nil {
		return err
	}
	return nil
}
```


Parámetros:
- nameCol (string): Nombre de la colección que se eliminará.

Retorno:
- error: Si ocurre un error durante la eliminación, se devuelve un objeto error. De lo contrario, se devuelve nil.

### Conn.go

Este código define funciones para conectarse a una base de datos MongoDB y realizar operaciones básicas en Go.

**Estructura**

```go
type MongoInstance struct {
    Client *mongo.Client
    Db     *mongo.Database
}
```
La función Connect() establece una conexión con la base de datos MongoDB.

```go
func Connect() error {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	server := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	dbName := os.Getenv("DB_NAME")
	var mongoUri = "mongodb://" + server + ":" + port + "/" + dbName

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

1. Carga las variables de entorno del archivo .env usando godotenv.Load(). El archivo .env probablemente contiene la información de conexión de la base de datos (host, puerto, nombre de la base de datos).
2. Recupera los valores de las variables de entorno DB_HOST, DB_PORT, y DB_NAME.
3. Construye una cadena de conexión URI de MongoDB usando la información de entorno.
4. Crea un nuevo cliente de MongoDB usando mongo.NewClient y aplica la cadena URI de conexión.
5. Establece un contexto de tiempo de espera para la conexión usando context.WithTimeout.
6. Intenta conectar el cliente a la base de datos usando client.Connect.
7. Obtiene una referencia a la base de datos específica usando client.Database(dbName).
8. Maneja errores de conexión registrándolos con log.Fatal.
9. Crea una instancia de MongoInstance con el cliente y la base de datos conectados. 1

### Data.go

Este código define estructuras Go para representar los datos que se almacenan en la base de datos MongoDB y se intercambian con el front-end.

```go
type Ram struct {
	Total      int `json:"totalRam"`
	En_uso     int `json:"memoriaEnUso"`
	Libre      int `json:"libre"`
	Porcentaje int `json:"porcentaje"`
}
```
Este es un ejemplo de las estrucuras, en total el backend tiene 5 estructuras. 

### main.go

Este código principal implementa un servidor web que monitorea y administra procesos en el sistema, y se comunica con una base de datos MongoDB.

```go
func main() {
	app := fiber.New()

	// Conectar a la base de datos
	if err := Database.Connect(); err != nil {
		log.Fatal("Error connecting to the database:", err)
	} else {
		log.Println("Successfully connected to the database.")
	}
	// Habilitar CORS
	app.Use(cors.New(cors.Config{
		AllowOrigins: "*",
		AllowMethods: "GET,POST,DELETE",
	}))

	// Definir rutas
	app.Get("/cpuyram", getPorcentajeRamyCpu)
	app.Get("/cpu", getCPUInfo)
	app.Get("cpu/iniProc/crear", StartProcess)
	app.Post("/cpu/killProc", KillProcess)

	if err := app.Listen(":3000"); err != nil { //aqui no se porque da error al poner una ip
		fmt.Println("Error en el servidor")
	}
}
```

**Flujo principal**

1. Inicialización del servidor:

- Se crea una instancia de la aplicación Fiber (app).
- Se establece la conexión con la base de datos MongoDB usando Database.Connect().
- Se habilita CORS para permitir solicitudes de diferentes orígenes.
- Se definen las rutas para el servidor web.
- Se inicia el servidor escuchando en el puerto 3000.

2. Rutas del servidor:

- ***/cpuyram*** : Obtiene información de la CPU y la RAM y la devuelve en formato JSON.
- ***/cpu*** : Obtiene información detallada de la CPU, incluyendo el uso y los procesos en ejecución, y la devuelve en formato JSON.
/cpu/iniProc/crear: Inicia un nuevo proceso en el sistema con el comando sleep infinity.
- ***/cpu/killProc***: Termina un proceso en ejecución mediante su PID (Process ID) recibido como parámetro.
3. Funciones para obtener información del sistema:

- ***getRAMdata()*** : Ejecuta un comando para obtener datos de la RAM en formato JSON y los decodifica en una estructura Model.Ram.
- ***getPorcentajeRamyCpu()*** : Obtiene el porcentaje de uso de la CPU y la RAM, los combina en un mapa y lo devuelve en formato JSON.
- ***getRAMInfo1()*** : Ejecuta un comando para obtener datos de la RAM en formato JSON y los decodifica en una estructura Model.Ram.
- ***getMem()*** : Obtiene datos de la RAM, los convierte de bytes a Megabytes, los guarda en la base de datos usando Controller.InsertRam y devuelve la información convertida.
- ***getCPUInfo()*** : Ejecuta un comando para obtener información de la CPU en formato JSON, la decodifica en una estructura Model.Cpu, calcula el porcentaje de uso libre y lo devuelve en formato JSON.


4. Funciones para interacturar con la base de datos.

- ***getCPU(cpuInfo *Model.Process)*** : Inserta información de un proceso en la base de datos usando Controller.InserProcess, incluyendo su PID, nombre, estado e ID del proceso padre (si existe).
- ***getCpuPercentage(nameCol string)*** : Ejecuta un comando para obtener el porcentaje de uso de la CPU, lo calcula, lo inserta en la base de datos usando Controller.InsertCpu y lo devuelve.
5. Funciones para gestionar procesos:

- ***StartProcess(c *fiber.Ctx)*** : Inicia un nuevo proceso con el comando sleep infinity y devuelve su PID.
- ***KillProcess(c *fiber.Ctx)***: Recibe el PID de un proceso por parámetro, envía una señal SIGKILL para terminarlo y devuelve un mensaje de confirmación.

Este código proporciona una base para un servidor web que monitorea el uso de CPU y RAM, y permite iniciar y detener procesos en el sistema

## Frontend

### main.jsx
Este código configura el entorno de desarrollo de React con Bootstrap e iconos, y luego renderiza el componente principal App en el elemento "root" del HTML, dando vida a la aplicación React en el navegador

```jsx
import React from 'react'
import ReactDOM from 'react-dom/client'
import App from './App'
import './styles/index.css'
import "bootstrap/dist/css/bootstrap.min.css"
import "bootstrap-icons/font/bootstrap-icons.css"

ReactDOM.createRoot(document.getElementById('root')).render(
  <React.StrictMode>
    <App />
  </React.StrictMode>,
)

```

### app.jsx
Este código define el componente principal de una aplicación React, probablemente usada para visualizar información del sistema en tiempo real.

```jsx

function App() {
  
  let component
  switch (window.location.pathname) {
    case "/":
      component = <RealTimeCharts />
      break;
    case "/cpuyram":
      component = <RealTimeCharts />
      break;
    case "/cpu":
      component = <ProcessTable />
      break;
    default:
  }

  return (
    <>
    <Head />
    {component}
    </>
  )
  
}
```

Este código implementa un sistema de navegación básica en una aplicación React. En función de la ruta actual, se muestra un componente específico para visualizar gráficos en tiempo real de CPU y RAM, o una tabla de procesos.

### Encabezado.jsx

Este código define un componente React llamado Cabeza que representa el encabezado (header) de la aplicación. Utiliza librerías de React Bootstrap para crear una barra de navegación.

```jsx
function Cabeza() {
  return (
    <Navbar bg="light" expand="lg" className="shadow-sm">
      <Container>
        <Navbar.Brand href="/cpuyram" className="d-flex align-items-center">
          <span className="ms-2">SO1-JUN 2024</span>
        </Navbar.Brand>
        <Navbar.Toggle aria-controls="basic-navbar-nav" />
        <Navbar.Collapse id="basic-navbar-nav">
          <Nav className="ms-auto">
            <Nav.Link href="/cpuyram">Monitoreo</Nav.Link>
            <Nav.Link href="/cpu">Tabla de Procesos</Nav.Link>
          </Nav>
        </Navbar.Collapse>
      </Container>
    </Navbar>
  );
}
```
El componente Cabeza renderiza una barra de navegación con el logo de la aplicación y dos enlaces: "Monitoreo" y "Tabla de Procesos". Utiliza estilos de React Bootstrap para una apariencia limpia y adaptable.

### Gradicas.jsx
Este código define un componente React (RealTimeCharts) que muestra gráficos circulares (pie charts) en tiempo real para el uso de CPU y RAM.

```jsx
function RealTimeCharts() {
  const [cpuUsage, setCpuUsage] = useState(null);
  const [ramUsage, setRamUsage] = useState(null);
  const url = "/back"; // Cambiar por la URL de tu API

  useEffect(() => {
    const fetchUsageData = () => {
      fetch(url + '/cpuyram') // Reemplaza con tu endpoint 
        .then(response => response.json())
        .then(data => {
          setCpuUsage(data.cpu_percentage);
          setRamUsage(data.ram_percentage);
          console.log('Datos recibidos:', data); 
        })
        .catch(error => console.error('Error fetching data:', error));
    };

    fetchUsageData(); 

    const interval = setInterval(() => {
      fetchUsageData(); 
    }, 500);

    return () => clearInterval(interval); 
  }, []);
```

**Generación de datos para el gráfico:**

const generatePieData = (label, percentage) => { ... }: Define una función generatePieData que toma el nombre del recurso (CPU o RAM) y su porcentaje de uso como parámetros, y devuelve los datos necesarios para el componente Pie de Chart.js.


### procesos.jsx
Este código define un componente React (ProcessTable) que muestra una tabla con información de los procesos del sistema, permitiendo matarlos y crear nuevos procesos básicos.

**Funciones para interactuar con la API:**

- ***fetchProcesses*** : Recupera la lista de procesos y la información del sistema desde la API.
- ***handleRefreshProcesses*** : Actualiza la lista de procesos.
- ***handleCreateProcess***: Envía una petición a la API para crear un nuevo proceso básico.
- ***handleKillProcess***: Envía una petición a la API para matar un proceso identificado por su PID.
- ***showAlert***: Muestra una alerta temporal en la pantalla.

## Docker
Docker es una plataforma de código abierto para desarrollar, implementar y ejecutar aplicaciones de forma segura y rápida.

### ¿Qué hace Docker?

- **Virtualización a nivel de contenedor**: Docker crea contenedores ligeros y aislados, llamados "contenedores Docker", que encapsulan una aplicación y todas sus dependencias.
- **Aislamiento**: Cada contenedor tiene su propio sistema de archivos, espacio de nombres de red y recursos de CPU y memoria, lo que los hace aislados entre sí y del sistema host.
- **Portabilidad**: Las aplicaciones empaquetadas en contenedores Docker se pueden ejecutar de manera consistente en cualquier entorno, ya sea en una computadora local, en la nube o en un servidor remoto.
- **Facilidad de uso**: Docker proporciona una interfaz de línea de comandos (CLI) y una API REST para crear, administrar y ejecutar contenedores.
Automatización: Docker se integra con herramientas de automatización como Kubernetes para orquestar la implementación y administración de aplicaciones a gran escala.

### Beneficios de usar Docker:

- **Aislamiento y seguridad**: Los contenedores aíslan las aplicaciones entre sí y del sistema host, lo que mejora la seguridad y reduce las dependencias.
- **Portabilidad**: Las aplicaciones se ejecutan de manera consistente en cualquier entorno, lo que facilita la implementación en diferentes plataformas.
- **Agilidad**: Los contenedores se pueden crear, implementar y escalar rápidamente, lo que acelera el desarrollo y la entrega de software.
- **Eficiencia de recursos**: Los contenedores comparten el kernel del sistema operativo, lo que reduce el uso de recursos y mejora la eficiencia.
- **Repetibilidad**: Las aplicaciones se ejecutan de manera predecible en cualquier entorno, lo que garantiza resultados consistentes.

### Casos de uso de Docker:

Desarrollo y pruebas de software: Docker permite a los desarrolladores trabajar en aplicaciones aisladas y probarlas en diferentes entornos.
Implementación de aplicaciones: Docker facilita la implementación de aplicaciones en entornos de producción, ya sea en la nube o en servidores locales.
Microservicios: Docker es ideal para desarrollar y ejecutar aplicaciones de microservicios, que son pequeñas y modulares.
Operaciones DevOps: Docker se integra con herramientas DevOps para automatizar la implementación, la administración y la entrega continua de software.


### docker-compose.yaml
Este código define un archivo de configuración para ejecutar una aplicación con múltiples contenedores utilizando Docker Compose. A continuación, se ofrece una descripción detallada de las secciones:
```jsx
version: '3'

services:
  database:
    image: mongo:latest
    container_name: mongo-container
    restart: always
    environment:
      - MONGO_INITDB_DATABASE=DB
    volumes:
      - mongo-data:/data/db
    ports:
      - '27017:27017'

  backend:
    image: kesm12/backend:latest
    container_name: backend-container
    environment:
      - DB_HOST=database
      - DB_PORT=27017
      - DB_NAME=DB
    ports:
      - '3000:3000'
    volumes:
      - type: bind
        source: /proc
        target: /proc
    restart: always
    depends_on:
      - database

  frontend:
    image: kems12/frontend:latest
    container_name: frontend-container
    ports:
      - '80:80'
    restart: always
    depends_on:
      - backend

volumes:
  mongo-data:
    external: false
```

En resumen, esta configuración de Docker Compose define tres servicios que trabajan juntos:

Un servicio de base de datos MongoDB.
Un servicio de aplicación backend que se conecta a la base de datos.
Un servicio de aplicación frontend que interactúa con el backend.
La configuración asegura que los servicios se inicien en el orden correcto y define cómo se comunican entre sí y acceden a los datos persistentes.

### Dockerfile Backend
Este código define una imagen Docker para una aplicación Go que utiliza una base de datos MongoDB.

- Base: FROM golang:alpine - La imagen base es golang:alpine, que proporciona un entorno mínimo de Go con Alpine Linux.

- Directorio de trabajo: WORKDIR /back - Establece el directorio de trabajo dentro del contenedor a /back.

- Copiado del código fuente: COPY . . - Copia todo el código fuente del proyecto desde la máquina host al directorio de trabajo del contenedor (/back).

- Inicialización de módulos: RUN go mod init main - Inicializa un archivo go.mod para gestionar las dependencias del proyecto.

- Instalación de dependencias:

1. RUN go get github.com/gorilla/mux

2. RUN go get github.com/gorilla/handlers

3. RUN go get go.mongodb.org/mongo-driver/mongo

4. RUN go get go.mongodb.org/mongo-driver/mongo/options

5. RUN go get github.com/gofiber/fiber/v2

6. RUN go get github.com/joho/godotenv


- Variables de entorno

- Exposición de puerto: EXPOSE 3000 - Expone el puerto 3000 del contenedor, que probablemente será utilizado por la aplicación Go para escuchar peticiones.

- Comando: CMD [ "go", "run", "main.go"] - Define el comando que se ejecuta al iniciar el contenedor.

**go run main.go ejecuta el programa principal (main.go) de la aplicación Go.**

### Dockerfile Frontend

Este código define una imagen Docker multi-etapa para una aplicación frontend basada en Node.js y servida por Nginx.

**Etapa 1 (builder):**

- ***Base***: node:20-alpine (imagen base de Node.js versión 20 con Alpine Linux)
- ***Directorio de trabajo***: /frontend
- ***Copiado de dependencias***: package.json y package-lock.json
- ***Instalación de dependencias***: npm install
- ***Copiado del código fuente: Se copia todo el código fuente del proyecto al directorio de trabajo.**

**Etapa 2 (final):**

- ***Base**: nginx:1.21-alpine (imagen base de Nginx versión 1.21 con Alpine Linux)
- ***Copiado de configuración***: nginx.conf de la carpeta nginx se copia a la configuración de Nginx.
- ***Copiado del build***: Se copia el directorio dist (que presumiblemente contiene los archivos estáticos compilados de la aplicación frontend) desde la etapa builder a la carpeta raíz del servidor web Nginx.
- ***Exposición del puerto***: Se expone el puerto 80 (puerto HTTP estándar)

## Nginx

Nginx (pronunciado "engine-ex") es un servidor web, proxy inverso, balanceador de carga, proxy de correo y proxy genérico TCP/UDP de código abierto, gratuito y de alto rendimiento. Es conocido por su estabilidad, eficiencia, escalabilidad y conjunto de funciones.

![img2](data:image/png;base64,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)

**Servidor web:** Nginx puede servir directamente contenido estático (HTML, CSS, JavaScript, imágenes) y contenido dinámico generado por otras aplicaciones a través de FastCGI o SCGI.

**Proxy inverso:** Puede actuar como un proxy inverso, redirigiendo solicitudes a servidores backend al tiempo que proporciona funciones como equilibrio de carga, almacenamiento en caché y seguridad.

**Balanceador de carga:** Nginx puede distribuir el tráfico entrante entre varios servidores backend para mejorar el rendimiento y la escalabilidad.

**Proxy de correo:** Puede funcionar como servidor proxy de correo, redirigiendo correos electrónicos a otros servidores de correo o realizando filtrado de contenido.

**Proxy genérico TCP/UDP:** Nginx también puede actuar como un servidor proxy TCP/UDP genérico, redirigiendo el tráfico entre diferentes protocolos de red.

### Ventajas clave de Nginx:

**Alto rendimiento:** Nginx es conocido por su capacidad para manejar una gran cantidad de conexiones simultáneas de manera eficiente, lo que lo hace adecuado para sitios web de alto tráfico.

**Bajo uso de memoria:** Tiene una huella de memoria mínima, lo que le permite ejecutarse en sistemas con recursos limitados.

**Escalabilidad:** Nginx se puede escalar fácilmente de forma horizontal agregando más servidores para manejar el aumento del tráfico.

**Flexibilidad:** Ofrece una amplia gama de funciones y se puede personalizar para diversos casos de uso a través de archivos de configuración.

**Código abierto:** Al ser de código abierto, Nginx es gratuito de usar y modificar, con una gran comunidad que brinda soporte y recursos.

### Casos de uso comunes de Nginx:

**Servidor web:** Muchos sitios web populares utilizan Nginx como su servidor web principal debido a su rendimiento y estabilidad.

**Balanceo de carga:** A menudo se utiliza para distribuir el tráfico entre varios servidores web o instancias de aplicación para lograr alta disponibilidad y escalabilidad.

**Proxy inverso:** Nginx se puede usar como un proxy inverso para descargar la carga de servir contenido estático y proporcionar medidas de seguridad adicionales para las aplicaciones backend.

**Almacenamiento en caché:** Puede almacenar en caché el contenido al que se accede con frecuencia para mejorar el rendimiento del sitio web y reducir la carga en los servidores backend.

**Transmisión de medios:** Nginx se puede utilizar para transmitir contenido multimedia de manera eficiente, como videos y audio.

**En general, Nginx es un servidor web potente y versátil que se puede utilizar para diversas tareas relacionadas con la web. Su eficiencia, escalabilidad y amplio conjunto de funciones lo convierten en una opción popular para impulsar aplicaciones web modernas.**

### nginx.conf
```conf
worker_processes 1;

events {
  worker_connections  1024;
}

http {
    server {
        listen 80;
        server_name localhost;

        root   /usr/share/nginx/html;
        index  index.html index.htm;
        include /etc/nginx/mime.types;

        gzip on;
        gzip_min_length 5;
        gzip_proxied expired no-cache no-store private auth;
        gzip_types text/plain text/css application/json application/javascript application/x-javascript text/xml application/xml application/xml+rss text/javascript;

        location / {
            try_files $uri $uri/ /index.html;
        }

        location /cpuyram/ {
            proxy_pass http://localhost:3000/cpuyram;
        }

        location /cpu/ {
            proxy_pass http://localhost:3000/cpu;
        }

    }
}
```


Este código configura un servidor web Nginx para servir contenido estático y actuar como proxy inverso para dos aplicaciones backend:

**Características principales:**

- Servidor web: Sirve contenido estático desde la carpeta /usr/share/nginx/html.

- Compresión Gzip: Comprime archivos estáticos para reducir el tamaño y mejorar la velocidad de carga.
- Proxy inverso:
Redirige las solicitudes a /cpuyram/ a la aplicación en http://localhost:3000/cpuyram.
Redirige las solicitudes a /cpu/ a la aplicación en http://localhost:3000/cpu.

**En resumen, este servidor Nginx proporciona:
Servidor web estático básico.
Balanceo de carga simple entre dos aplicaciones backend.**