ARG VERSION=

FROM --platform=$BUILDPLATFORM golang:1.23-alpine AS builder
ARG VERSION
WORKDIR /src
RUN --mount=type=cache,target=/go/pkg/mod \
    --mount=type=bind,target=/src \
<<EOT
    set -xeuo pipefail
    GOLDFLAGS="-s -w"
    GOLDFLAGS="${GOLDFLAGS} -X 'github.com/soramitsukhmer-lab/vault-plugin-catalog/version.Version=${VERSION:-unspecified}'"
    for GOOS in darwin linux; do
        for GOARCH in amd64 arm64; do
            export GOOS GOARCH
            go build -ldflags="${GOLDFLAGS}" -o /out/vault-plugin-catalog-${GOOS}-${GOARCH} .
        done
    done
EOT

FROM scratch AS binaries
COPY --from=builder --link /out/vault-plugin-catalog-* /

FROM scratch
ARG VERSION
ARG TARGETARCH
COPY --from=binaries --link /vault-plugin-catalog-linux-${TARGETARCH} /vault-plugin-catalog
