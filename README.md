# playground

This is Wanjia's coding playground

# build and push docker image
export TAG=ghcr.io/opheadacheh/bazel_build_base:0.0
docker build -f BuildBaseDockerImages/v0.0/Dockerfile . -t $TAG && docker push $TAG

# enter build base
docker run -v $(pwd):/playground -it --rm ghcr.io/opheadacheh/bazel_build_base:0.0 bash

# run build base and connect to VScode
docker run -v $(pwd):/playground -it -d ghcr.io/opheadacheh/bazel_build_base:0.0 bash
