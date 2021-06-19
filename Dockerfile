# build stage
FROM golang:1.16-alpine AS builder
WORKDIR /go/src/github.com/atajsic/sca-wishlist-notifier
ADD . ./
RUN go build -o scawln

# final stage
FROM alpine:latest
WORKDIR /app/
COPY --from=builder /go/src/github.com/atajsic/sca-wishlist-notifier/scawln /app/scawln
ENTRYPOINT ["./scawln"]