# from the root of project , run :
# GIT_COMMIT_VERSION=$( git show -s --format='format:%H')
# GIT_COMMIT_TIME=$( git show -s --format='format:%aI')
# docker buildx build -t release:latest --build-arg GIT_COMMIT_TIME=${GIT_COMMIT_TIME} --build-arg GIT_COMMIT_VERSION=${GIT_COMMIT_VERSION} --platform=linux/amd64 --output type=docker -f ./images/controllerimage/Dockerfile .

ARG BASE_IMAGE=weizhoulan/baseimage:bdecfc7c5e21b5341e12a35b7dcb846c5eaaeae2
ARG GOLANG_IMAGE=docker.io/library/golang:1.18.0@sha256:03baa0921adf5fb547f48df8b16137de82b5ccc34986c22a967055379369f614

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
#go to the controler directory and make
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

RUN  echo "GIT_COMMIT_VERSION=$GIT_COMMIT_VERSION" \
        && echo "GIT_COMMIT_TIME=$GIT_COMMIT_TIME"

RUN groupadd -f test \
    && echo ". /etc/profile.d/bash_completion.sh" >> /etc/bash.bashrc


COPY --from=builder /tmp/install/${TARGETOS}/${TARGETARCH}/bin/*   /usr/bin/
COPY --from=builder /tmp/install/${TARGETOS}/${TARGETARCH}/bash-completion/*  /etc/bash_completion.d

CMD ["/usr/bin/sleep","10d"]











