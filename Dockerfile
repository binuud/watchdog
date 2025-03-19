#
# This is a multi target docker file
# The first state - builder - builds all the binaries
# Subsequent targets are used to build the final docker image, with the prebuild binaries
# This docker file is used to create multiple docker images
#

FROM golang:latest AS builder

LABEL maintainer="Binu Udayakumar"

RUN mkdir /build
COPY ./ /build/
WORKDIR /build
ENV GO111MODULE=auto

# Build WatchDog - Watch for Certificate expiry, connectivity, domain expiry
RUN CGO_ENABLED=0 GOOS=linux go build  -o appStoreRepo services/watchdog/cmd/main.go
