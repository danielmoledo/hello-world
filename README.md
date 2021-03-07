# Gataca Test

````
The goal of this mini project is to build a micro-service in GoLang.
Basically, you should create a new project, using Go as a code language, with a service called hello-world. This endpoint should trigger the process to store in a database (whatever you want) the log of each access (ID, request timestamp). So, for each request to that service, there will be an entry in the database.
You can use whatever additional technology you want to use:
Services (Echo, Gin, Mux, …)
Database (Gorm, …)
````

Para arrancar la aplicación:
````
$ go run main.go
````

Para probar el endpoint:
````
$ curl -X POST http://localhost:8080/hello-world
````

### Estructura de paquetes

Dado que no existe un estándar oficial, o por lo menos yo no lo he encontrado, que defina la estructura de paquetes, 
aunque es cierto que existe un [estándar no oficial](https://github.com/golang-standards/project-layout) popular, he optado por la siguiente distribución:

````
hello-world/
├── controllers
│   └── helloworld.go
├── go.mod
├── go.sum
├── main.go
├── model
│   └── helloworld.go
├── README.md
├── repository
│   └── helloworld.go
├── services
│   └── helloworld.go
````

He elegido esta estructura por legibilidad, ya que indica que voy a utilizar los patrones "repository" y "service" en esta tarea. 

### Diseño
Aunque el código podría ser mucho más sucinto ya que se trata de una prueba pequeña, he intentado añadir un poco de profundidad añadiendo las capas repository y service.
La finalidad de la primera, es servir como punto único de acceso a los datos, con el objetivo de centralizar en un mismo punto el código que no puede ser abstralizado de
la elección de la base de datos y que por lo tanto solo haría falta cambiar la implementación concreta si se decidiera cambiar de motor de bbdd. 
El motivo de la segunda, es encapsular la lógica de negocio del resto de factores que no sean puramente reglas de negocio (acceso a datos, gestión de peticiones, DTOs, ...).

Para la bbdd he usado [GORM](https://gorm.io/docs/) y [GIN](https://github.com/gin-gonic/gin) como framework web.
 

