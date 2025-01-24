# docker login
# docker buildx build --platform linux/amd64,linux/arm64 -t kenshimuto/hotrod:latest  --push .
FROM --platform=$BUILDPLATFORM golang:bookworm AS builder
ARG TARGETOS
ARG TARGETARCH
WORKDIR /app
COPY go.mod .
COPY go.sum .
RUN go mod download
COPY . .
RUN cd examples/hotrod && CGO_ENABLED=0 GOOS=$TARGETOS GOARCH=$TARGETARCH go build -o hotrod-linux -buildvcs=false .

FROM scratch
LABEL maintainer="kmuto@kmuto.jp"
EXPOSE 8080 8081 8082 8083 8084
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=builder /app/examples/hotrod/hotrod-linux /go/bin/hotrod-linux
ENTRYPOINT ["/go/bin/hotrod-linux"]
CMD ["all"]
