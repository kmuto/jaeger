# Copyright (c) 2024 The Jaeger Authors.
# SPDX-License-Identifier: Apache-2.0

FROM golang:bookworm AS builder

FROM scratch
EXPOSE 8080 8081 8082 8083
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY ./hotrod-linux /go/bin/hotrod-linux
ENTRYPOINT ["/go/bin/hotrod-linux"]
CMD ["all"]
