# Test GO

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
 

