# base go image
FROM golang:1.18-alpine as builder

RUN mkdir /app

COPY . /app

WORKDIR /app

RUN GCO_ENABLED=0 go build -o displayApp

RUN chmod +x ./displayApp

# FROM alpine:latest

# RUN mkdir /app 

# COPY --from=builder /app/displayApp /app

CMD [ "/app/displayApp" ]