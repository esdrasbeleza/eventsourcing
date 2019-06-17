FROM golang:1.12-alpine as builder
RUN apk add --no-cache git
WORKDIR /tmp/src/app
COPY . .
RUN go build

FROM alpine:latest
EXPOSE 8080
COPY --from=builder /tmp/src/app/backend /tmp/backend
ENTRYPOINT ["/tmp/backend"]
