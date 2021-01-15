
# Herramientas
Este fichero describe que herramientas han utilizado para desarrollar este proyecto.
  - Lenguaje: `Go`, este lenguaje ofrece buen rendimiento, es compilado, por lo que podremos generar un binario para su ejecucion. Tiene una sintaxis relativamente sencilla, lo que lo hace sencillo de escribir. Además una de sus princiaples características es que hace el paralelismo y la concurrencia más sencillas mediante el uso de `goroutines`. Además el compilado es bastante rápido. 
  - Construcción: `make`, es una herramienta flexible, independiente del lenguaje, permitirá simplificar las operaciones de construcción y mantenimiento del proyecto. Aunque `Go` viene con su propia herramienta de construcción, el uso de un Makefile permitirá simplificar los comandos, y reducirá nuestra carga al realizar operaciones de este tipo.
  - Test: El propio lenguaje `Go` incluye el uso de test. Para ejecutarlos hace falta crear un fichero *_test.go*, y luego realizar *go test* en el directorio con los ficheros *.go*. Además usando opciones como `-cover` podremos ver el porcentaje de codigo cubierto por los tests.
