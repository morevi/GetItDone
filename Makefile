build:
		go build -o getitdone cmd/GetItDone/main.go

test:
		go test -cover ./...

install:
		go install ./...

run:
		go run cmd/GetItDone/main.go
