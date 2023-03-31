FROM golang:1.18-alpine as builder

RUN mkdir app

COPY . /app

WORKDIR /app

RUN GCO_ENABLE=0 go build -o statictvApp ./cmd/api

RUN chmod +x /app/statictvApp

FROM alpine:latest

RUN mkdir /app

COPY --from=builder /app/statictvApp /app
COPY --from=builder /app/static /app/static

CMD ["/app/statictvApp"]
