# MANUAL TECNICO

## kvm
KVM (Kernel-based Virtual Machine) es una tecnología de virtualización de código abierto integrada en el kernel de Linux. Permite ejecutar múltiples sistemas operativos invitados (máquinas virtuales) en un único ordenador físico.

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

Nota: Para utilizar el programa adecuadamente se recomienda instalar los modulos que son generados a partir de los archivos cpu.c y ram.c (sudo insmod cpu.ko y ram.ko) ubicados en el reposotorio, luego de esto ejecutar docker compose up y ya deberia de observarse de forma correcta la gráfica en el frontend.