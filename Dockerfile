# syntax=docker/dockerfile:1

# builder image
FROM golang:1.16 AS builder
WORKDIR /build
COPY . ./
RUN make

# target image
FROM alpine:latest
RUN apk --no-cache add ca-certificates
WORKDIR /app
COPY --from=builder /build/interpolator .
COPY entrypoint.sh /app
ENTRYPOINT ["/app/entrypoint.sh"]
