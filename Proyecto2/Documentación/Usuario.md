# MANUAL DE USUARIO 

**Graficas (Grafana)**
Grafana es una plataforma de código abierto para la visualización, monitoreo y análisis de datos en tiempo real. Fue diseñada para ayudar a los usuarios a entender y analizar grandes volúmenes de datos provenientes de diversas fuentes.

![Grafica](../Documentación/img/Graficas.jpg)

1. En la grafica de pastel se ven los porcertajes de los  tweets por pais

2. El numero **2441** que vemos en la pantalla es la catidad de  tweets 

3. El numero **87** Indica los países participantes

4. La grafica que esta debajo de los numero es la cantidad total de tweets de forma individual.

5. La grafica que esta en la parte inferior es la cantidad de tweets por paise

**Locust**
Locust es una herramienta de código abierto utilizada para pruebas de carga. Está diseñada para simular el comportamiento de múltiples usuarios accediendo a una aplicación web al mismo tiempo, con el fin de evaluar su rendimiento bajo condiciones de alto tráfico. A continuación, se detallan algunos aspectos clave de Locust:

![img1](../Documentación/img/img3.jpg)

En la imagen de arriba podemos ver que para poder usar locust primero debemos de poner nuestras quedenciales.

A continuación una explicación de lo que es locust

¿Qué es Locust?

1. Locust es un framework de pruebas de carga que permite crear scripts de prueba en Python, donde los usuarios virtuales pueden realizar diversas acciones en una aplicación web.
2. Es un proyecto de código abierto, lo que significa que está disponible de forma gratuita y puede ser modificado y mejorado por la comunidad.

![img2](../Documentación/img/img11.jpg)


¿Para qué sirve Locust?
1. Pruebas de Carga: Locust se utiliza principalmente para realizar pruebas de carga en aplicaciones web, simulando un gran número de usuarios concurrentes.
2. Evaluación de Rendimiento: Ayuda a identificar cuellos de botella y problemas de rendimiento en la aplicación al observar cómo se comporta bajo diferentes niveles de carga.
3. Escalabilidad: Permite simular miles o incluso millones de usuarios mediante la distribución de la carga en múltiples máquinas.


![img3](../Documentación/img/img22.jpg)

En este proyecto el uso de Locust sería simular el acceso de miles de usuarios a un sitio que creamos para observar cómo se comporta el servidor y asegurarse de que puede manejar el tráfico sin fallar.

**Base de datos**

MongoDB Compass es una herramienta gráfica de administración y análisis para bases de datos MongoDB. Fue desarrollada por MongoDB Inc. y proporciona una interfaz visual intuitiva para interactuar con las bases de datos MongoDB. 

![img4](../Documentación/img/img4.jpg)

¿Para qué sirve MongoDB Compass?
1. Exploración de Datos
2. Consultas Visuales
3. Análisis de Datos
4. Gestión de Índices
5. Validación de Esquemas
6. Operaciones CRUD
7. Seguridad y Autenticación
8. Exportación e Importación de Datos

Para este proyecto utilizamos compa para revisar los datos que nos estan entrando a la base de datos, en ***DB2*** encontramos una carpeta llamada register, en esta carpeta estan guardados todos los datos que entran, la estructura de los datos son las siguiente.

1. _id: Es el identificador
2. texto: Viene una cita, del comportamiento del clima en el pais
3. pais: En el viene el nombre del pais del cual se esta enviando la cita
4. fecha: Conctiene la fecha y hora exacta en que entro a la base de datos el archivo con la informacion ("La petición al servidor")

![img5](../Documentación/img/img5.jpg)

En la imagen superior podemos ver la cantidad de registros dentro del servidor, y otros datos, que nos serviran para conocer cuanto pesa nuestra base de datos, cuantos datos tiene guardados.

