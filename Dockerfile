# $ sudo docker build -t portfolio:latest  .

FROM golang:1.22.1-alpine3.19 as build

WORKDIR /app
RUN apk add --update build-base
COPY ./go.mod .
COPY ./go.sum .
RUN export GOPROXY=direct
RUN go mod download
COPY . /app
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build  -ldflags="-w -s" -o ./bin/portfolio ./cmd/main.go

FROM alpine:3.19.1

WORKDIR /app

COPY --from=build /app/./bin/portfolio /app
# COPY --from=build /app/tmpl /app/tmpl

EXPOSE 5000

ENTRYPOINT ["/app/portfolio"]
