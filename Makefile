build:
		go build -o getitdone cmd/getitdone/main.go

test:
		go test -cover ./...

build-test:
		go test -c -cover ./cmd/getitdone
		go test -c -cover ./pkg/tareas

install:
		go install ./...

run:
		go run cmd/getitdone/main.go
