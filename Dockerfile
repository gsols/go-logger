# syntax = docker/dockerfile:1.2

# get modules, if they don't change the cache can be used for faster builds
FROM golang:1.18@sha256:02c05351ed076c581854c554fa65cb2eca47b4389fb79a1fc36f21b8df59c24f AS base
ENV GO111MODULE=on
ENV CGO_ENABLED=0
ENV GOOS=linux
ENV GOARCH=amd64
WORKDIR /src
COPY go.* .
RUN --mount=type=cache,target=/go/pkg/mod \
    go mod download

# build th application
FROM golang:1.18@sha256:02c05351ed076c581854c554fa65cb2eca47b4389fb79a1fc36f21b8df59c24f AS build
# temp mount all files instead of loading into image with COPY
# temp mount module cache
# temp mount go build cache
RUN --mount=target=. \
    --mount=type=cache,target=/go/pkg/mod \
    --mount=type=cache,target=/root/.cache/go-build \
    go build -ldflags="-w -s" -o /app/main ./cmd/go/*.go

# Import the binary from build stage
FROM gcr.io/distroless/static:nonroot@sha256:2556293984c5738fc75208cce52cf0a4762c709cf38e4bf8def65a61992da0ad as prd
COPY --from=build /app/main /
# this is the numeric version of user nonroot:nonroot to check runAsNonRoot in kubernetes
USER 65532:65532
ENTRYPOINT ["/main"]
