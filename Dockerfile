FROM golang:1.22rc1-bullseye AS builder

WORKDIR /goapp

COPY go.mod go.sum Makefile ./

COPY .env ./

COPY cmd common internal ./

## SSL Certificates
# RUN mkdir /usr/local/share/ca-certificates/extra
# COPY ./certs /usr/local/share/ca-certificates/extra
# RUN update-ca-certificates
## SSL Certificates

ENV _GOARCH=amd64 _GOOS=linux

# RUN apt update
# RUN apt-get -yqq install build-essential libssl-dev libffi-dev python3-pip python3-dev gnupg

RUN make build-prod
# RUN make build-prod


FROM golang:1.22rc1-bullseye AS runner

WORKDIR /app

COPY --from=builder /goapp/bin/go-crud-linux /app/go-crud

EXPOSE 9090

CMD ["./goapp/go-crud"]
