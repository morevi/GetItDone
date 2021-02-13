
# Integración continua.
Se han configurado GH Actions, Travis y CircleCI.
Para los 3 sistemas de CI se utilizan badges para visualizar el estado del código desde el README.md.

## Actions
Es difícil organizar varios servicios de CI de forma que se esperen unos a otros, así que este sistema se encargará de realizar

Se han configurado el workflow [docker-build-test](../github/workflows/docker-build-test.yml), que realiza 3 _jobs_:
- `check`. Combrueba si han habido cambios en el Dockerfile
- `build`. Si han habido cambios, reconstruye y sube la nueva imagen de docker; si no: no hace nada.
- `test`. Espera a que acabe `build`, y ejecuta los tests en ese contenedor.
![actions-log](images/ci/actions-log.png)

## Travis
Se ha [configurado](../.travis.yml) de forma que realizaremos tests con diferentes versiones de `Go` y en los sistemas operativos Linux y OSX. 
Se realizaran 2 tests: 2 versiones x 2 os.
Las versiones que utilizamos es _tip_ (la última), y la _1.15.8_ que es la última en este momento, pero algun día dejará de serlo, y el proyecto seguirá programado con las carácteristicas de esta versión.
Para ello, no se utiliza nuestra imagen docker para los tests, si no que Travis utiliza las suyas para los SO y versiones indicados. Podemos ver en una captura que en cada ejecución de Travis, se realizan 6 pruebas.
![travis-log](images/ci/travis-log.png)

## Circle CI
Este sistema es bastante veloz, así que la utilizaremos para visualizar los resultados de los tests unitarios y agilizar el desarrollo del proyecto.
Como problema, es que no he conseguido sincronizarlo con actions, de forma que CircleCI espere a Actions y se sigan mostrando todas las acciones CI en el status en Github (me refiero a [esto](https://github.com/morevi/GetItDone/issues/61#issuecomment-778606513)), por tanto, es posible que los tests se realizen con una version desactualizada de la imagen.

Su configuración es [esta](../.circleci/config.yml). Podemos verlo funcionando:
![circleci-log](images/ci/circleci-log.png)

## Uso del gestor de tareas
Se han añadido algunas entradas al [Makefile](../Makefile) para trabajar con docker de forma más sencilla ejecutando más facilmente los comandos más recurrentes, permitiendonos utilizar siempre el Taskfile para trabajar.

```
docker-test:
		docker run -t -v `pwd`:/test morevi/getitdone:latest

docker-build:
		docker build . --tag morevi/getitdone:latest

docker-push:
		docker push morevi/getitdone:latest
```

De este modo tambien se simplifican los ficheros de configuración de integracion continua, al usar estas sentencias reducidas en lugar de la version original.
Y en caso de querer hacer un cambio (por ejemplo el nombre o la version de la imagen), solo es necesario editar este fichero!

## Aprovechamiento del docker
Tanto GithubAction como Circle CI, utilizan la misma imagen docker, especificada a través del Makefile.

## Avance del proyecto

