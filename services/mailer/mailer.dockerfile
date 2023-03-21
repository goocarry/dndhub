FROM golang:1.18-alpine as builder

RUN mkdir /app

COPY . /app

WORKDIR /app

RUN GCO_ENABLE=0 go build -o mailerApp ./cmd/api

RUN chmod +x /app/mailerApp

# build a tiny docker image
# FROM alpine:latest

# RUN mkdir /app

# COPY --from=builder /app/mailerApp /app

CMD ["/app/mailerApp"]