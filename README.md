# playground

This is Wanjia's coding playground

# build and push docker image
export TAG=registry.jihulab.com/opheadacheh/playground/bazel_build_base:0.1
docker build -f BuildBaseDockerImages/v0.1/Dockerfile . -t $TAG && docker push $TAG

# enter build base
docker run -v $(pwd):/playground -it --rm registry.jihulab.com/opheadacheh/playground/bazel_build_base:0.1 bash

# run build base and connect to VScode
docker run -v $(pwd):/playground -it -d registry.jihulab.com/opheadacheh/playground/bazel_build_base:0.1 bash
