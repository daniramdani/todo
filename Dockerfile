FROM golang:1.10-alpine AS build

MAINTAINER Devops Team of Todo 

ARG git_tag
ARG git_commit
ENV project=github.com/daniramdani/todo

# INSTALL Godep
RUN apk update && apk add ca-certificates git curl wget zip build-base
RUN curl https://raw.githubusercontent.com/golang/dep/master/install.sh | sh

WORKDIR /go/src/${project}
COPY . .
RUN dep ensure -v

# Build APP
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags "-s -X $project/cmd.Version=$git_tag -X $project/cmd.GitHash=$git_commit" -a -installsuffix cgo -o /bin/app

# Cleanup all pre-build data
FROM alpine:latest
RUN apk update && apk add ca-certificates curl zip wget
ENV project=github.com/daniramdani/todo

WORKDIR /bin
COPY --from=build /go/src/${project}/.env .env
COPY --from=build /bin/app app

ENTRYPOINT ["app"]
EXPOSE 3015