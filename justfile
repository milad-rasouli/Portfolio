build:
    CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build  -ldflags="-w -s" -o ./bin/portfolio ./main.go
run: build
    ./bin/portfolio





