build:
    CGO_ENABLED=1 GOOS=linux GOARCH=amd64 go build  -ldflags="-w -s" -o ./bin/portfolio ./main.go
run: build
    mkdir -p data
    ./bin/portfolio

generate:
	templ generate

test:
    go test -v ./...


