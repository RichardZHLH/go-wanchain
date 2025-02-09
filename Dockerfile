# ubuntu

FROM ubuntu:22.04
MAINTAINER molin
WORKDIR /root

RUN apt-get update
RUN apt-get install -y wget

COPY ./loadScript/monitor.sh /bin/
COPY ./build/bin/* /bin/
EXPOSE 17717/tcp 17717/udp