# Primera actividad

El objetivo de esta actividad es ejecutar una aplicación en forma local utilizando `docker`, y en aws utilizando `elasticbeanstack`. Se espera que los equipos cumplan con el desafío técnico al mismo tiempo que reflexionan sobre cuestiones de implementación y despliegue relacionadas con los temas vistos en el curso.

A lo largo de esta guía aparecen preguntas y pedidos de captura de pantalla. El entregable deberá ser un documento que incluya las respuestas a las preguntas y capturas pedidas, que se debe entregar en la actividad creada a tal fin en aulas.

Los equipos deben resolver las posibles dificultades que surjan a lo largo del proceso por su cuenta. En caso de no poder cumplir con alguno de los pasos, se deberá dejar captura del error obtenido y una explicación de porque entienden que se genera.

Cada paso de esta guía tiene un tiempo máximo sugerido. Si no lo completa pasado ese tiempo se recomienda dejar la explicación correspondiente en el documento y pasar al siguiente.

# Paso 0: Pre-requisitos

1. Contar con una máquina en la que poder ejecutar `docker` y `docker-compose`
2. Contar con un cliente web que permita realizar peticiones http (postman, curl o cualquier otro)
3. Clonar este repo en alguna carpeta en el disco

## Paso 1: Ejecución local (15 minutos)

Acceda a la carpeta `todo` del repo clonado y dedique unos minutos a entender el código de la aplicación (archivo `main.go`) y el `Dockerfile`.

__Pregunta 1: Explique los pasos del archivo `Dockerfile`. ¿Porque cree que hay dos directivas `FROM` en el mismo?__

Construya la imagen especificada en el `Dockerfile` y ejecute un contenedor con la misma siguiendo estos pasos en una consola:

```
docker build -t todo .
docker run -p 8080:8080 todo
```

__Captura 1: Muestre el resultado de la consola al ejecutar el contenedor__

Con el contenedor funcionando, utilizar un cliente _http_ para realizar los siguientes _requests_:

```
GET     http://localhost:8080/tasks
POST    http://localhost:8080/tasks  BODY: { "description": "my first task"}
POST    http://localhost:8080/tasks  BODY: { "description": "my second task"}
GET     http://localhost:8080/tasks
```

__Captura 2: Muestre el resultado del último request__

## Paso 2: Ejecución en aws elastic beanstalck

Cree un archivo llamado `todo.zip` con todos los archivos del directorio `todo`. Importante: el zip debe tener todos los archivos en la raíz, no un directorio dentro. La forma más fácil es seleccionar todos los archivos y crear el zip en base a eso (no en base a la carpeta).

Acceda a su cuenta en el lab de aws y dentro de la consola a Elastic Beanstalk:

![beanstalk](./images/uno.png)

Seleccionar la opción de "Crear Aplicación":

![beanstalk](./images/dos.png)
