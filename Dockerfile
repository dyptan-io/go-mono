# syntax=docker/dockerfile:1
ARG GO_VERSION=1.26

FROM --platform=$BUILDPLATFORM golang:${GO_VERSION}-alpine AS builder
WORKDIR /repo

COPY . .

ARG TARGETOS TARGETARCH APP
RUN --mount=type=cache,target=/root/.cache/go-build \
    --mount=type=cache,target=/root/go/pkg/mod \
    CGO_ENABLED=0 GOOS=$TARGETOS GOARCH=$TARGETARCH \
    go -C cmd/${APP} build -ldflags="-s -w" -trimpath -o /out/app .

FROM gcr.io/distroless/static
COPY --from=builder /out/app /app
ENTRYPOINT ["/app"]
