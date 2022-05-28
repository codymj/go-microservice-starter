# syntax=docker/dockerfile:1

###
### build
###

FROM golang:1.18.2-stretch AS build1

ADD go.mod /app/go.mod
ADD go.sum /app/go.sum

WORKDIR /app
RUN go mod download

###
### deploy
###

FROM build1 as build2

ENV CGO_ENABLED=0

COPY . /app

WORKDIR /app/cmd/app
RUN go build -o app_binary

###
### run
###

FROM gcr.io/distroless/static

USER nonroot:nonroot

COPY --from=build2 --chown=nonroot:nonroot /app/cmd/app/app_binary /app/app_binary
COPY --from=build2 --chown=nonroot:nonroot /app/settings.yaml /app/settings.yaml

WORKDIR /app
ENTRYPOINT ["/app/app_binary"]
