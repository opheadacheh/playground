FROM ghcr.io/opheadacheh/bazel_build_base:0.3

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
RUN rm -rf /root/*
RUN sudo apt update
RUN apt-get install git -y
RUN sudo apt upgrade -y