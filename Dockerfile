FROM golang:1.11.2-alpine3.8

ENV GOPATH /go
ENV CGO_ENABLED 0
WORKDIR /go/src/github.com/Matsushin/go-ecs-deploy/

RUN apk update && apk upgrade && \
    apk add --no-cache git

COPY ./ /go/src/github.com/Matsushin/go-ecs-deploy/

EXPOSE 8080
CMD ["go", "run", "main.go"]