FROM golang:1.18-alpine as builder

RUN mkdir /app

COPY . /app

WORKDIR /app

RUN go build -o listenerApp .

RUN chmod +x ./listenerApp

FROM alpine:latest

RUN mkdir /app 

COPY --from=builder /app/listenerApp /app

CMD [ "/app/listenerApp" ]