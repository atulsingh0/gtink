FROM golang:1.19 as build
WORKDIR /build
COPY . .
RUN go mod tidy
RUN CGO_ENABLED=0 go build -o gtink gtink/*
RUN chgrp 0 gtink && chmod g+X gtink

FROM alpine as cacerts
RUN apk update
RUN apk add --no-cache ca-certificates
RUN update-ca-certificates

FROM scratch
COPY --from=cacerts /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=build /build/gtink .
USER 1001
ENTRYPOINT ["/gtink"]