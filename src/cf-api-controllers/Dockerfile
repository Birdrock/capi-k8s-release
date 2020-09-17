ARG BUILD_IMAGE=golang:1.13
ARG RUN_IMAGE=gcr.io/distroless/static:nonroot

FROM ${BUILD_IMAGE} as builder
ARG GIT_SHA
RUN if [ -z "$GIT_SHA" ]; then echo "GIT_SHA not set"; exit 1; else : ; fi

WORKDIR /capi-kpack-watcher

# Copy the go source
COPY go.mod .
COPY go.sum .

WORKDIR /workspace
# Copy the Go Modules manifests
COPY go.mod go.mod
COPY go.sum go.sum
# cache deps before building and copying source so that we don't need to re-download as much
# and so that source changes don't invalidate our downloaded layer
RUN go mod download

# Copy the go source
COPY main.go main.go
COPY controllers/ controllers/
COPY cf/ cf/
COPY image_registry/ image_registry/

# Build
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 GO111MODULE=on go build -a -o manager main.go

# Use distroless as minimal base image to package the manager binary
# Refer to https://github.com/GoogleContainerTools/distroless for more details
FROM ${RUN_IMAGE}
WORKDIR /
COPY --from=builder /workspace/manager .
USER nonroot:nonroot

ENTRYPOINT ["/manager"]

ARG GIT_SHA
ARG BUILD_IMAGE
ARG RUN_IMAGE
LABEL org.opencontainers.image.source=https://github.com/Birdrock/capi-k8s-release.git \
            org.opencontainers.image.revision=$GIT_SHA \
            org.cloudfoundry.project.name=capi-k8s-watcher \
            org.cloudfoundry.project.component.name=capi-kpack-watcher \
            org.cloudfoundry.image.base.build=$BUILD_IMAGE \
            org.cloudfoundry.image.base.run=$RUN_IMAGE
