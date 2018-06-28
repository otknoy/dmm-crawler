FROM golang:latest as builder

ENV CGO_ENABLED=0
ENV GOOS=linux
ENV GOARCH=amd64
WORKDIR /go/src/github.com/otknoy/dmm-crawler
COPY . .
RUN make

# runtime image
FROM alpine

ENV OUTPUT_DIR=/output

RUN apk add --no-cache ca-certificates
COPY --from=builder /go/src/github.com/otknoy/dmm-crawler/dmm-crawler /dmm-crawler

# EXPOSE 8080
ENTRYPOINT ["/dmm-crawler"]