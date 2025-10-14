FROM golang:1.25.3 as build

RUN mkdir -p /app
RUN apt-get update

WORKDIR /app

COPY ./ ./

RUN go build -v -ldflags="-X 'main.version=v0.1.1'" moose.go

FROM ubuntu:24.04

RUN apt-get update
RUN apt-get install curl -y

RUN mkdir -p /app/configs
RUN mkdir -p /app/var/logs

WORKDIR /app

COPY --from=build /app/moose /app/moose
COPY --from=build /app/config.dist.yml /app/configs/config.dist.yml

EXPOSE 8000

VOLUME /app/configs
VOLUME /app/var

RUN ./moose version

CMD ["./moose", "server", "--sse", "--url", "http://127.0.0.1:8000"]
