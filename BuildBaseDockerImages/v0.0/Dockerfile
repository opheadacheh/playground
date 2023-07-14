FROM ubuntu:22.04

RUN apt-get update
RUN apt-get install sudo vim openjdk-11-jdk wget -y
RUN wget https://github.com/bazelbuild/bazel/releases/download/6.2.1/bazel-6.2.1-installer-linux-x86_64.sh
RUN apt-get install unzip -y
RUN chmod +x bazel-6.2.1-installer-linux-x86_64.sh
RUN ./bazel-6.2.1-installer-linux-x86_64.sh
RUN export PATH="$PATH:$HOME/bin"