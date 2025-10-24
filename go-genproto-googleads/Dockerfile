FROM ubuntu:22.04

RUN apt-get update && \
    apt-get install -y build-essential git curl python3 unzip openjdk-11-jre-headless zip

RUN curl -L https://github.com/bazelbuild/bazelisk/releases/download/v1.19.0/bazelisk-linux-amd64 \
    -o /usr/local/bin/bazelisk && \
    chmod +x /usr/local/bin/bazelisk && \
    ln -s /usr/local/bin/bazelisk /usr/local/bin/bazel

WORKDIR /workspace
