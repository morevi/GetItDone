## CONTENEDORES
Para facilitar la automatización de los tests, se va a crear una imagen de docker, de forma que se pueda lanzar tanto en local como en diferentes servicios y que el resultado sea siempre el mismo.

### Contenedor base.
En la mayoría de casos, una imagen oficial de un lenguaje, `golang`,será más eficiente en cuanto a velocidad como en peso que una imagen de uso genérico como podría ser la de `ubuntu` con *go* instalado. 

Por tanto, vamos a centrarnos en las imágenes oficiales de `golang`. La más reducida en tamaño podemos ver que son las basadas en Alpine, una distribución Linux que busca tener un tamaño mínimo. Para empezar a trabajar, esta será nuestra mejor opción.

### El dockerfile.
Se crea la imagen a partir de `golang:1.15.7-alpine` y se configura el Dockerfile de forma que:
- Utiliza pocas sentencias RUN para evitar capas.
- Utiliza un nuevo usuario para ejecutar los tests sin privilegios,
- Utiliza el taskfile para realizar las acciones.
Échale un [vistazo](../Dockerfile)!

### Automatización y Docker Hub.
Se ha creado [.github/workflows/docker-build.yml](.github/workflows/docker-build.yml), de forma que utilizaremos `github actions` para automátizar la **construcción**, testeo del proyecto y **subida a los registros**.
![docker-test](images/docker/workflows.png)

Este workflow se ejecutará en cualquier tipo de push (o pull) al repositorio, de forma que reconstruye la imagen si se ha modificado el Dockerfile y lo sube a los registros.

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
Se ha intentado.

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

![Comparación](images/docker/full-comparison.png)


