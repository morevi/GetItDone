build:
		go build -o getitdone cmd/getitdone/main.go

test:
		CGO_ENABLED=0 go test -cover ./...

build-test:
		go test -c -cover ./cmd/getitdone
		go test -c -cover ./pkg/tareas

docker-test:
		docker run -t -v `pwd`:/test morevi/getitdone:latest

docker-build:
		docker build . --tag morevi/getitdone:latest

docker-push:
		docker push morevi/getitdone:latest

install:
		go install ./...

run:
		go run cmd/getitdone/main.go
