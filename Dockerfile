FROM golang:latest AS builder
WORKDIR /go/src/

ADD . .

RUN GIN_MODE=release go build -ldflags="-s -w" -o release/cyclonedx-enrich .

FROM cgr.dev/chainguard/wolfi-base:latest AS runtime

COPY --from=builder --chown=nonroot:nonroot /go/src/release/ /app
WORKDIR /app

ENTRYPOINT [ "/app/cyclonedx-enrich"]
