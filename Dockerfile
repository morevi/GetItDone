# Usamos una imagen optimizada de go sobre alpine para que sea lo más ligera
# posible durante la compilacion.
FROM golang:1.15.7-alpine AS build
WORKDIR /test
COPY . .

# Reutilizamos un unico RUN para evitar la creación de layers
# Ademas, usamos --no-cache para que no se guarde la caché del manejador de paquetes. Aunque igualmente, al ser una máquina de construcción, no afectará a la máquina final.
RUN \
  apk update && apk add --no-cache git make \
  && make build-test

# Usamos una máquina mucho más ligera alpine, para copiar los ejecutables en ella.
# Y sea más rápido trabajar con ella, compartirla o subirla a las diferentes plataformas
FROM alpine
WORKDIR /test
COPY --from=build /test /test
CMD ./GetItDone.test && ./tareas.test
