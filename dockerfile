FROM golang:1.14-buster

ARG APP=${APP}

WORKDIR /go/src/${APP}
COPY . .
RUN go get -u ./... \
    && go build -o ./bin/myapp ./cmd/myapp/main.go
EXPOSE 3000
CMD ./bin/myapp