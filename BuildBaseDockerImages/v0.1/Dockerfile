FROM ghcr.io/opheadacheh/bazel_build_base:0.0

ENV http_proxy http://127.0.0.1:7890
ENV https_proxy http://127.0.0.1:7890

RUN sudo apt-get install tree -y

# install required perl packges for genhtml
RUN perl -MCPAN -e 'install "DateTime"'
RUN perl -MCPAN -e 'install "Date:Parse"'

# install lcov (genhtml included)
RUN wget https://github.com/linux-test-project/lcov/releases/download/v2.0/lcov-2.0.tar.gz
RUN tar -xvzf lcov-2.0.tar.gz
RUN cd lcov-2.0; make install
RUN rm -rf lcov*

# fetch required resources to run code coverage
COPY BuildExamples/cpp_examples/stage4 /stage4
RUN cd /stage4; bazelisk coverage --combined_report=lcov //lib:hello-time-test
RUN rm -rf /stage4