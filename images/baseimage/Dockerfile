# on the local directory, ,run: docker buildx build -t testbase:latest --platform=linux/arm64,linux/amd64 --output type=docker .


ARG GOLANG_IMAGE=docker.io/library/golang:1.17.8@sha256:f675106e44f205a7284e15cd75c41b241329f3c03ac30b0ba07b14a6ea7c99d9
ARG UBUNTU_IMAGE=docker.io/library/ubuntu:20.04@sha256:8ae9bafbb64f63a50caab98fd3a5e37b3eb837a3e0780b78e5218e63193961f9

#========= build gops ===============
FROM --platform=${BUILDPLATFORM} ${GOLANG_IMAGE} as gops-cni-builder
COPY /build-gops.sh /build-gops.sh

#use alibaba debian source
#COPY /sources.list /etc/apt/sources.list
#RUN apt-key adv --keyserver keyserver.ubuntu.com --recv-keys 648ACFD622F3D138

#with debian 11 os
RUN apt-get update && apt-get install -y binutils-aarch64-linux-gnu binutils-x86-64-linux-gnu

# build-gops.sh will build both archs at the same time
WORKDIR /
RUN chmod +x /build-gops.sh && /build-gops.sh


#========== root image ==============
FROM ${UBUNTU_IMAGE} as rootfs
COPY /install-others.sh /install-others.sh
COPY /configure-iptables-wrapper.sh /configure-iptables-wrapper.sh
COPY /iptables-wrapper /usr/sbin/iptables-wrapper
COPY /test-image.sh /test-image.sh

# Change the number to force the generation of a new git-tree SHA. Useful when
# we want to re-run 'apt-get upgrade' for stale images.
ENV FORCE_BUILD=1

#fix warning: debconf: delaying package configuration, since apt-utils is not installed
ARG DEBIAN_FRONTEND=noninteractive
RUN apt-get update && apt-get upgrade -y

WORKDIR /
RUN chmod +x /install-others.sh && /install-others.sh && rm -f /install-others.sh
RUN chmod +x /configure-iptables-wrapper.sh && /configure-iptables-wrapper.sh && rm -f /configure-iptables-wrapper.sh

ARG TARGETPLATFORM
COPY --from=gops-cni-builder /out/${TARGETPLATFORM}/bin/gops /bin/gops

#test result
RUN chmod +x /test-image.sh && /test-image.sh && rm -f /test-image.sh


#============
FROM scratch
LABEL maintainer="maintainer@helloword"
COPY --from=rootfs / /




