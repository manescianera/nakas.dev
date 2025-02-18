FROM cgr.dev/chainguard/go:latest AS fetcher
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download

FROM ghcr.io/a-h/templ:latest AS generator
WORKDIR /app
COPY --chown=65532:65532 . .
RUN ["templ", "generate"]

FROM cgr.dev/chainguard/go:latest AS builder
WORKDIR /app
COPY --from=generator /app .
RUN CGO_ENABLED=0 GOOS=linux go build -a -ldflags '-extldflags "-static"' -buildvcs=false -o app

FROM cgr.dev/chainguard/static:latest
WORKDIR /app
COPY --from=builder --chown=nonroot:nonroot /app/app .
COPY --chown=nonroot:nonroot cv.yaml .
EXPOSE 8080
USER nonroot:nonroot
ENTRYPOINT ["/app/app"]
