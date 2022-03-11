# from the root of project , run : docker buildx build -t release:latest --platform=linux/amd64 --output type=docker -f ./images/controllerimage/Dockerfile .

ARG BASE_IMAGE=weizhoulan/baseimage:bdecfc7c5e21b5341e12a35b7dcb846c5eaaeae2
ARG GOLANG_IMAGE=docker.io/library/golang:1.17.8

# TARGETARCH is an automatic platform ARG enabled by Docker BuildKit.
# like amd64 arm64
ARG TARGETARCH

#======= build bin ==========
FROM --platform=${BUILDPLATFORM} ${GOLANG_IMAGE} as builder

# TARGETOS is an automatic platform ARG enabled by Docker BuildKit.
ARG TARGETOS
# TARGETARCH is an automatic platform ARG enabled by Docker BuildKit.
ARG TARGETARCH

COPY . /src
WORKDIR /src
RUN  make GOARCH=${TARGETARCH}   \
        DESTDIR_BIN=/tmp/install/${TARGETOS}/${TARGETARCH}/bin \
        DESTDIR_BASH_COMPLETION=/tmp/install/${TARGETOS}/${TARGETARCH}/bash-completion \
        build-bin install-bin install-bash-completion


#====== release image =======

FROM ${BASE_IMAGE}

LABEL maintainer="maintainer@helloword"

# TARGETOS is an automatic platform ARG enabled by Docker BuildKit.
ARG TARGETOS
# TARGETARCH is an automatic platform ARG enabled by Docker BuildKit.
ARG TARGETARCH

RUN groupadd -f test \
    && echo ". /etc/profile.d/bash_completion.sh" >> /etc/bash.bashrc

COPY --from=builder /tmp/install/${TARGETOS}/${TARGETARCH}/bin/*   /usr/bin/
COPY --from=builder /tmp/install/${TARGETOS}/${TARGETARCH}/bash-completion/*  /etc/bash_completion.d

CMD ["/usr/bin/sleep","10d"]











