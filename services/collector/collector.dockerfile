# base go image
FROM golang:1.18-alpine as builder

RUN mkdir /app

COPY . /app

WORKDIR /app

RUN CGO_ENABLED=0 go build -o collectorApp

RUN chmod +x /app/collectorApp

# TODO(@goocarry): why tempplates are not working at final image??
# build a tiny docker image
# FROM alpine:latest

# RUN mkdir /app

# COPY --from=builder /app/templates /app/templates
# COPY --from=builder /app/. /app

CMD ["/app/collectorApp"]