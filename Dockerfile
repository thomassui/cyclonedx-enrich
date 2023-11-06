FROM golang:latest AS builder
WORKDIR /go/src/
ENV CGO_ENABLED=0

ADD . .

RUN go build -ldflags="-s -w" -o release/cyclonedx-enrich .

FROM scratch AS runtime
WORKDIR /

COPY --from=builder /go/src/release/cyclonedx-enrich .

ENTRYPOINT [ "/cyclonedx-enrich" ]