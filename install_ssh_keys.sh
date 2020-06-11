#!/bin/bash

set -e -v

mkdir -p ~/.ssh
echo "${SSH_PRV_KEY}" > ~/.ssh/id_rsa
chmod 700 ~/.ssh
chmod -R 600 ~/.ssh/*
eval "$(ssh-agent -s)"ls

git config --global --add url."git@github.com:".insteadOf "https://github.com/"
ssh-add ~/.ssh/id_rsa
ssh-keyscan github.com >> ~/.ssh/known_hosts
go get -u ./...
go build -o ./bin/myapp ./cmd/myapp/main.go

ls -al

exec "$@"