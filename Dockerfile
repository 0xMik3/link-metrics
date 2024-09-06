FROM golang:1.21.4-alpine AS builder

RUN apk add upx git gcc musl-dev


ENV GO111MODULE=on \
  CGO_ENABLED=0 \
  GOOS=linux \
  GOARCH=amd64

WORKDIR /src

COPY . .

RUN go mod download && \
    go build -trimpath -ldflags "-s -w -extldflags '-static'" \
    -installsuffix cgo -tags netgo \
    -o /bin/service cmd/service/main.go && \
    strip /bin/service && \
    upx -q -9 /bin/service

FROM scratch
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=builder /bin/service /bin/service
COPY --from=builder /src/regexes.yaml /bin/regexes.yaml

ENV PORT 8080
EXPOSE 8080

ENTRYPOINT ["/bin/service"]