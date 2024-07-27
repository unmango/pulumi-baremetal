FROM ubuntu:noble-20240605

WORKDIR /app
COPY bin/provisioner .
RUN [ -x provisioner ] || exit 1
