# To build the archgst image, just run:
# > docker build -t archgst .
#
# In order to work properly, this Docker container needs to have a volume that:
# - as source points to a directory which contains a config.toml and firebase-config.toml files
# - as destination it points to the /home folder
#
# Simple usage with a mounted data directory (considering ~/.archgst/config as the configuration folder):
# > docker run -it -v ~/.archgst/config:/home archgst archgst parse config.toml firebase-config.json
#
# If you want to run this container as a daemon, you can do so by executing
# > docker run -td -v ~/.archgst/config:/home --name archgst archgst
#
# Once you have done so, you can enter the container shell by executing
# > docker exec -it archgst bash
#
# To exit the bash, just execute
# > exit
FROM golang:1.17-alpine3.15 AS build-env
ARG arch=x86_64

# Install dependencies
RUN apk update && \
    apk add --no-cache curl make git libc-dev bash gcc linux-headers eudev-dev py-pip ca-certificates 

# Set working directory for the build
WORKDIR /archgst

# Add source files
COPY . .

# See https://github.com/CosmWasm/wasmvm/releases
ADD https://github.com/CosmWasm/wasmvm/releases/download/v1.0.0-beta10/libwasmvm_muslc.aarch64.a /lib/libwasmvm_muslc.aarch64.a
ADD https://github.com/CosmWasm/wasmvm/releases/download/v1.0.0-beta10/libwasmvm_muslc.x86_64.a /lib/libwasmvm_muslc.x86_64.a
RUN sha256sum /lib/libwasmvm_muslc.aarch64.a | grep 5b7abfdd307568f5339e2bea1523a6aa767cf57d6a8c72bc813476d790918e44
RUN sha256sum /lib/libwasmvm_muslc.x86_64.a | grep 2f44efa9c6c1cda138bd1f46d8d53c5ebfe1f4a53cf3457b01db86472c4917ac

# Copy the library you want to the final location that will be found by the linker flag `-lwasmvm_muslc`
RUN cp /lib/libwasmvm_muslc.${arch}.a /lib/libwasmvm_muslc.a

# force it to use static lib (from above) not standard libgo_cosmwasm.so file
RUN LEDGER_ENABLED=false BUILD_TAGS=muslc LINK_STATICALLY=true make build

# Final image
FROM alpine:3.13

# Install ca-certificates
RUN apk add --update ca-certificates
WORKDIR /home

# Install bash
RUN apk add --no-cache bash

# Copy over binaries from the build-env
COPY --from=build-env /archgst/build/archgst /usr/bin/archgst
