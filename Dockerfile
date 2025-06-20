FROM --platform=$BUILDPLATFORM golang:1.23-alpine AS builder
ARG VERSION=
WORKDIR /src
RUN --mount=type=cache,target=/go/pkg/mod \
    --mount=type=bind,target=/src \
<<EOT
    set -xeuo pipefail
    for GOOS in darwin linux; do
        for GOARCH in amd64 arm64; do
            export GOOS GOARCH
            go build -ldflags="-s -w" -o /out/vault-plugin-catalog${VERSION:+-${VERSION}}-$GOOS-$GOARCH .
        done
    done
EOT
FROM scratch AS binaries
COPY --from=builder --link /out/vault-plugin-catalog-* /
