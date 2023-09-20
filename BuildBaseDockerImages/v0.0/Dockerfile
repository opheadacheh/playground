FROM ubuntu:22.04

# switch ubuntu sources list to aliyun
RUN printf "deb http://mirrors.aliyun.com/ubuntu/ jammy main restricted universe multiverse\n\
deb-src http://mirrors.aliyun.com/ubuntu/ jammy main restricted universe multiverse\n\
deb http://mirrors.aliyun.com/ubuntu/ jammy-security main restricted universe multiverse\n\
deb-src http://mirrors.aliyun.com/ubuntu/ jammy-security main restricted universe multiverse\n\
deb http://mirrors.aliyun.com/ubuntu/ jammy-updates main restricted universe multiverse\n\
deb-src http://mirrors.aliyun.com/ubuntu/ jammy-updates main restricted universe multiverse\n\
deb http://mirrors.aliyun.com/ubuntu/ jammy-proposed main restricted universe multiverse\n\
deb-src http://mirrors.aliyun.com/ubuntu/ jammy-proposed main restricted universe multiverse\n\
deb http://mirrors.aliyun.com/ubuntu/ jammy-backports main restricted universe multiverse\n\
deb-src http://mirrors.aliyun.com/ubuntu/ jammy-backports main restricted universe multiverse\n" > /etc/apt/sources.list

# install essential tools
RUN apt-get update
RUN apt-get upgrade -y
RUN apt-get install sudo vim openjdk-11-jdk wget git unzip build-essential -y

# setup basic environment variables
ENV PATH="$PATH:$HOME/bin"

# initial work directory
WORKDIR /playground

# install golang
RUN wget https://golang.org/dl/go1.20.6.linux-amd64.tar.gz
RUN sudo tar -xvf go1.20.6.linux-amd64.tar.gz -C /usr/local
ENV GOROOT=/usr/local/go
ENV GOPATH=/root/go
ENV PATH=$GOPATH/bin:$GOROOT/bin:$PATH
RUN rm -f go1.20.6.linux-amd64.tar.gz

# install bazelisk
RUN go install github.com/bazelbuild/bazelisk@latest

# install bazel
RUN bazelisk version