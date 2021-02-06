build:
		go build -o getitdone cmd/GetItDone/main.go

test:
		go test -cover ./...

build-test:
		go test -c -cover ./cmd/GetItDone
		go test -c -cover ./pkg/tareas

install:
		go install ./...

run:
		go run cmd/GetItDone/main.go
