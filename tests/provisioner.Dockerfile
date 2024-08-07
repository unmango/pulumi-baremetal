# syntax=docker/dockerfile:1
FROM --platform=$BUILDPLATFORM ubuntu:noble-20240605

ENV DEBIAN_FRONTEND=noninteractive
RUN apt update \
    && apt install -y \
        wget \
    && rm -rf /var/lib/apt/lists/*

WORKDIR /app
COPY bin/provisioner .
ENTRYPOINT [ "/app/provisioner" ]
