FROM golang:1.14-buster

ARG APP=${APP}

WORKDIR /go/src/${APP}
COPY . .

EXPOSE 3000

ENTRYPOINT [ "/bin/bash", "./install_ssh_keys.sh" ]
CMD [ "./bin/myapp" ] 
