## Construcción
Necesitas tener `go`, `make`, y `git` instalados para poder compilar el proyecto.
No es necesaria ninguna otra dependencia, `go` se encargará de traer las todo que necesite.

Para ejecutar el programa sin compilar, puedes utilizar el siguiente comando:
```
make run
```

Pará obtener un ejecutable e instalarlo donde quieras:
```
make build
```

Si tienes $GOPATH en tu $PATH, puedes instalar directamente el binario con:
```
make install
```

[Aquí](tools.md) puedes leer porqué hemos elegido `make`.
