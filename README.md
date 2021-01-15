# GetItDone

## Descripción.
El objetivo es crear una API que permita gestionar, mediante operaciones CRUD, listas de tareas por etiquetas, y deadlines. De forma que puedas estar siempre organizado y productivo. Se podrá levantar en un servidor y pueda actuar de microservicio, y pueda ser utilizada desde otras aplicaciones.

## ¿Por qué este proyecto?
Personalmente, siempre me encuentro olvidando algunas de las cosas que tengo pendientes y termino recordandola el último día para realizarla de forma estresada. He elegido este proyecto porque pareció algo útil (aunque algo específico), y, con respecto al tema de la asignatura de Infraestructura Virtual de la UGR, me permite desarrollar un microservicio desplegable en la web, que sea de utilidad a alguna otra aplicación (es decir, el proyecto solo cubre el microservicio/API para su posible uso).

## Construcción
Necesitas tener `go` y `make` instalados para compilar el proyecto.
No necesitas instalar ninguna otra dependencia, `go` se engargara de traer las dependencias que no esten ya instaladas en el sistema.

Si quieres obtener un ejecutable:
```
make build
```

Si tienes  $GOPATH en tu $PATH, puedes instalar directamente el binario con:
```
make install
```

Para ejecutarlo puedes utilizar el ejecutable generado o puedes realizar:
```
make run
```

[Aqui](docs/tools.md) puedes leer porqué hemos elegido `make`

## Tests
Para ejecutar todos los tests usamos:
```
make test
```

Este comando ejecutará los tests de las clases *Project* y *Task*

Puedes leer más sobre el entorno de test [aqui](docs/tools.md)

## Historias de usuario
Se ha creado la [etiqueta]() `user-stories` para poder visualizarlas de forma sencilla en la ventana de issues.

Cada issue del tipo `user-stories`, contiene un título conciso, y una frase de historia de usuario en la descripción.

Aquí enlaces a cada una de las historias de usuario:
 - [HU01](https://github.com/morevi/GetItDone/issues/1) Informacion de las proyectos.
 - [HU02](https://github.com/morevi/GetItDone/issues/2) Gestion de proyectos.
 - [HU03](https://github.com/morevi/GetItDone/issues/3) Busqueda y organizacion de proyectos.
 - [HU04](https://github.com/morevi/GetItDone/issues/4) Visualizacion del proyecto.
 - [HU05](https://github.com/morevi/GetItDone/issues/19) Informacion de la tarea.
 - [HU06](https://github.com/morevi/GetItDone/issues/5) Gestion de tareas.

## Clases.
### ![Task](pkg/tareas/task.go)
Representa una tarea o "todo", incluye una descripción, una fecha limite y su estado (completada o no). Su implementación es un paso adelante en la HU05 y la HU06.

### ![Project](pkg/tareas/project.go)
Representa nuestra colección de tareas, incluye tags para organizarlos, una descripción y la serie de tareas que se le asignen. Es un paso hacia HU01, HU02 y HU04.

## Pasos a realizar.
 - Elegir un contenedor base para las pruebas.
 - Crear un dockerfile
 - Crear una clase `escritorio` desde la que gestionar los proyectos y sus tareas
 - Terminar una historia de usuario.
 - Seguir documentando issues. [Tercer milestone](https://github.com/morevi/GetItDone/milestone/3).

## Pasos realizados.
Puedes leer sobre ellos ![aqui](docs/pasos.md).

## Más información.
 - [El problema a resolver](docs/problemDescription.md)
 - [Configuración de `git`](docs/git.md)
 - [Herramientas](docs/tools.md)

## Autor.
Francisco José Moreno Vílchez [@morevi](https://github.com/morevi)

