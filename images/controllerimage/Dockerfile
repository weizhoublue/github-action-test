# from the root of project , run :
# VERSION=v1.0.0
# GIT_COMMIT_VERSION=$( git show -s --format='format:%H')
# GIT_COMMIT_TIME=$( git show -s --format='format:%aI')
# docker buildx build -t release:latest --build-arg VERSION=${VERSION} --build-arg GIT_COMMIT_TIME=${GIT_COMMIT_TIME} --build-arg GIT_COMMIT_VERSION=${GIT_COMMIT_VERSION} --platform=linux/amd64,linux/arm64 --output type=docker -f ./images/controllerimage/Dockerfile .

ARG BASE_IMAGE=ghcr.io/weizhoublue/github-action-test/baseimage:8700413ad57ead614f734b1311cf83231941713c@sha256:118ffa12150f4e6572a694250e7d7ce6fd6a555d28231f751c0089e2aa6f9f86
ARG GOLANG_IMAGE=docker.io/library/golang:1.17.8@sha256:f675106e44f205a7284e15cd75c41b241329f3c03ac30b0ba07b14a6ea7c99d9

# TARGETARCH is an automatic platform ARG enabled by Docker BuildKit.
# like amd64 arm64
ARG TARGETARCH

#======= build bin ==========
FROM --platform=${BUILDPLATFORM} ${GOLANG_IMAGE} as builder

#below args could be from workflow docker/build-push-action of build-args
# TARGETOS is an automatic platform ARG enabled by Docker BuildKit.
ARG TARGETOS
# TARGETARCH is an automatic platform ARG enabled by Docker BuildKit.
ARG TARGETARCH
#cunstome args
#data race and lock debug
ARG RACE
#strip binary
ARG NOSTRIP
#no optimize  for binary
ARG NOOPT
#quiet make
ARG QUIET_MAKE

COPY . /src
#go to the controller directory and make
WORKDIR /src/cmd/controller
RUN  make GOARCH=${TARGETARCH}   \
        RACE=${RACE} NOSTRIP=${NOSTRIP} NOOPT=${NOOPT} QUIET_MAKE=${QUIET_MAKE} \
        DESTDIR_BIN=/tmp/install/${TARGETOS}/${TARGETARCH}/bin \
        DESTDIR_BASH_COMPLETION=/tmp/install/${TARGETOS}/${TARGETARCH}/bash-completion \
        all install install-bash-completion


#====== release image =======

FROM ${BASE_IMAGE}

LABEL maintainer="maintainer@helloword"

# TARGETOS is an automatic platform ARG enabled by Docker BuildKit.
ARG TARGETOS
# TARGETARCH is an automatic platform ARG enabled by Docker BuildKit.
ARG TARGETARCH

ARG GIT_COMMIT_VERSION
ENV GIT_COMMIT_VERSION=${GIT_COMMIT_VERSION}
ARG GIT_COMMIT_TIME
ENV GIT_COMMIT_TIME=${GIT_COMMIT_TIME}
ARG VERSION
ENV VERSION=${VERSION}

RUN  echo "GIT_COMMIT_VERSION=$GIT_COMMIT_VERSION" \
        && echo "GIT_COMMIT_TIME=$GIT_COMMIT_TIME" \
        && echo "VERSION=$VERSION"

RUN groupadd -f test \
    && echo ". /etc/profile.d/bash_completion.sh" >> /etc/bash.bashrc


COPY --from=builder /tmp/install/${TARGETOS}/${TARGETARCH}/bin/*   /usr/bin/
COPY --from=builder /tmp/install/${TARGETOS}/${TARGETARCH}/bash-completion/*  /etc/bash_completion.d

CMD ["/usr/bin/sleep","10d"]











