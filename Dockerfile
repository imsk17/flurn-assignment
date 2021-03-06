# Use Multi Step Dockerfile to Create Tiny Images
FROM alpine:3.13.1 AS base

FROM golang:1.16.3-alpine AS builder
RUN apk update
RUN apk add build-base
RUN mkdir /build
ADD . /build
WORKDIR /build
RUN go build -o assignment -ldflags "-s" api/main.go

FROM base as FINAL
WORKDIR /app
COPY --from=builder /build/assignment .
CMD [ "/app/assignment" ]
