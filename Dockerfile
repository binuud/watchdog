#
# This is a multi target docker file
# The first state - builder - builds all the binaries
# Subsequent targets are used to build the final docker image, with the prebuild binaries
# This docker file is used to create multiple docker images
#

FROM --platform=$BUILDPLATFORM golang:alpine AS builder
ARG TARGETOS
ARG TARGETARCH

LABEL maintainer="Binu Udayakumar"

RUN mkdir /build
COPY ./ /build/

## For mounting config files on distroless containers
RUN mkdir /configs

WORKDIR /build

ENV GO111MODULE=auto

# Build WatchDog - Watch for Certificate expiry, connectivity, domain expiry
RUN CGO_ENABLED=0 GOOS=${TARGETOS} GOARCH=${TARGETARCH} go build  -o watchDog cmd/watchdog/main.go
RUN CGO_ENABLED=0 GOOS=${TARGETOS} GOARCH=${TARGETARCH} go build  -o watchDogServer cmd/watchdogServer/main.go

FROM gcr.io/distroless/static

LABEL maintainer="Binu Udayakumar"
USER nonroot:nonroot


## HTTP Port
EXPOSE 9080

## GRPC Port
EXPOSE 9090 

ARG BUILD_VER
ENV BUILD_VER=$BUILD_VER

## create a directory called dataVolume, using busybox
## since scratch images do not allow for directory creation
COPY --from=builder --chown=nonroot:nonroot /configs /configs

# copy compiled app
COPY --from=builder --chown=nonroot:nonroot /build/watchDog /watchDog
COPY --from=builder --chown=nonroot:nonroot /build/watchDogServer /watchDogServer


ENTRYPOINT ["./watchDogServer"]