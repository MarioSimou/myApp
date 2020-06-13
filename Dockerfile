FROM golang:1.14-buster

ARG SSH_PRV_KEY
ARG APP

RUN mkdir -m 700 -p ~/.ssh \
    && echo "${SSH_PRV_KEY}" > ~/.ssh/id_rsa \
    && chmod -R 600 ~/.ssh/* \
    && eval "$(ssh-agent -s)" \
    && git config --global --add url."git@github.com:".insteadOf "https://github.com/" \ 
    && ssh-add ~/.ssh/id_rsa \
    && ssh-keyscan github.com >> ~/.ssh/known_hosts
    
WORKDIR /go/src/${APP}
COPY . .
RUN go get -u ./... \
    && go build -o ./bin/myapp ./cmd/myapp/main.go

EXPOSE 3000
CMD [ "./bin/myapp" ] 
