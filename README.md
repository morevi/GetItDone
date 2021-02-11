# GetItDone
[![Actions](https://github.com/morevi/GetItDone/workflows/Docker%20Build/badge.svg?branch=master)](https://github.com/morevi/GetItDone/actions?query=workflow%3Adocker-build)
[![Travis](https://travis-ci.com/morevi/GetItDone.svg?branch=master)](https://travis-ci.com/morevi/GetItDone)
[![CircleCI](https://circleci.com/github/morevi/GetItDone.svg?style=svg)](https://app.circleci.com/pipelines/github/morevi/GetItDone?branch=circleci-project-setup)
## Descripción.
El objetivo es crear una API que permita gestionar, mediante operaciones CRUD, listas de tareas por etiquetas, y deadlines. De forma que puedas estar siempre organizado y productivo. Se podrá levantar en un servidor y pueda actuar de microservicio, y pueda ser utilizada desde otras aplicaciones.

## DOCKER
Para facilitar la automatización de los tests, se va a crear una imagen de docker, de forma que se pueda lanzar tanto en local como en diferentes servicios y que el resultado sea siempre el mismo.

### Contenedor base.
En la mayoría de casos, una imagen oficial de un lenguaje, `golang`,será más eficiente en cuanto a velocidad como en peso que una imagen de uso genérico como podría ser la de `ubuntu` con *go* instalado. 

Por tanto, vamos a centrarnos en las imágenes oficiales de `golang`. La más reducida en tamaño podemos ver que son las basadas en Alpine, una distribución Linux que busca tener un tamaño mínimo. Para empezar a trabajar, esta será nuestra mejor opción.

### El dockerfile.
Resumidamente, se compilarán los tests en una imagen de `golang`, y luego se copiarán los binarios en una imagen `alpine`
```
# Usamos una imagen optimizada de go sobre alpine para que sea lo más ligera
# posible durante la compilacion.
FROM golang:1.15.7-alpine
WORKDIR /test
COPY . .

# Reutilizamos un unico RUN para evitar la creación de layers
# Ademas, usamos --no-cache para que no se guarde la caché del manejador de paquetes. Aunque igualmente, al ser una máquina de construcción, no afectará a la máquina final.
RUN \
  apk update && apk add --no-cache git make \
  && addgroup -S go && adduser -S go -G go

USER go

CMD make test

```
### Automatización y Docker Hub.
Se ha creado [.github/workflows/docker-build.yml](.github/workflows/docker-build.yml), de forma que utilizaremos `github actions` para automátizar la construcción, testeo del proyecto y subida a Docker Hub.
![docker-test](docs/images/docker/workflows.png)

Este workflow se ejecutará en cualquier tipo de push (o pull), además es posible ejecutarlo manualmente desde la ventana de actions.
Las tareas se realizaran secuencialmente sobre la misma instancia de *ubuntu-latest*.


Básicamente se realiza:
1. Se sitúa sobre el proyecto y realiza el build.
2. Luego ejecuta los test.
3. Posteriormente realiza login en DockerHub usando las credenciales ofrecidas por github (username) o en secretos (password).
4. Procede a publicar en DockerHub la imagen.

Si cualquier paso da lugar a fallo, se indicará error tanto en el action como en el badge al inicio del README.md

### Registros de contenedores.
#### Docker Hub
Es un servicio con una documentación mucho más extensa, y al estar directamenteofrecido por Docker, la compatibilidad es inmediata, sin más configuración que la creación de una cuenta y un repositorio.

Los repositorios en DockerHub tienen el requisito de que deben estar nombrados en minúscula, así que para probar que funciona, puedes hacer:
```
docker run -t -v `pwd`:/test morevi/getitdone
```

[Aquí](https://hub.docker.com/repository/docker/morevi/getitdone) puedes ver el contenedor en el repositorio de DockerHub.

#### Github Docker Registry
La documentación no es suficientemente amplia como para ser útil en la resolución de dudas. Además, es un servicio que aún está en beta, lo que supone que esta sujeto a posibles cambios, que hace más dícil de mantener el workflow.

### Optimización de la imagen.
Las medidas para reducir el tamaño de la imagen final que se publicará en DockerHub:
- Uso de pocas directivas `COPY` y `RUN`:
Cada aparición de estas genera una nueva capa o 'layer' que docker utiliza para optimizar la construcción de futuras imágenes. Pero a cambio se incrementa es espacio que utilizan. Por tanto, reutilizaremos estas sentencias mediante el uso de *&&*.
- Uso de `--no-cache`:
Los 'package manager' utilizan una caché, para optimizar futuras instalaciones de paquetes. Como nuestra imagen no requiere a priori es autocontenida, y no requiere de más paquetes (aunque podría usarse de imagen base para otra imagen), esta opción nos va permitir ahorrar más MB.

#### Comparación
Solo compararé las imágenes más eficientes.

| Base          | Optimización | Tamaño | Tiempo |
| ------------- | ------------- | ------------ | ------------ |
| Golang:alpine  | Pocas capas + --no-cache | 314 MB | 3.25 segundos |
| Golang:alpine  | Ninguna | 314 MB | 3.52 segundos |
| Alpine con golang instalado | pocas capas + --nocache | 457MB | 3.51 segundos |

![Comparación](docs/images/docker/full-comparison.png)

## Pasos a realizar.
 - Continuar por la sexta semana del curso.
 - Comenzar el enrutamiento de la api.
 - Terminar una historia de usuario.

## Pasos realizados.
Puedes leer sobre ellos ![aqui](docs/pasos.md).

## Más información.
 - [¿Por qué este proyecto?](docs/why.md)
 - [Historias de usuario](docs/hu.md)
 - [Clases](docs/classes.md)
 - [Taskfile y builds](docs/builds.md)
 - [Tests](docs/tests.md)
 - [El problema a resolver](docs/problemDescription.md)
 - [Configuración de `git`](docs/git.md)
 - [Herramientas](docs/tools.md)

## Autor.
Francisco José Moreno Vílchez [@morevi](https://github.com/morevi)

