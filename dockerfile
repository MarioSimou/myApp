FROM golang:1.14-buster
WORKDIR /go/src/app
COPY . .
RUN go get -u ./... \
    && go build -o ./bin/myapp ./cmd/myapp/main.go
EXPOSE 3000
CMD ./bin/myapp