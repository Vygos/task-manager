FROM golang:1.23.5 as builder

WORKDIR /opt/

COPY . .

RUN GOOS=linux GOARCH=amd64 go build -o ./bin/app ./cmd/main.go

FROM alpine

COPY --from=builder /opt/bin/app /app
COPY --from=builder /opt/dev.yaml /

ENTRYPOINT ["/app"]