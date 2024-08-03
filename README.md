# playground

This is Wanjia's coding playground

# build and push docker image
export TAG=ghcr.io/opheadacheh/bazel_build_base:0.2
docker build --network host -f BuildBaseDockerImages/v0.2/Dockerfile . -t $TAG && docker push $TAG

# enter build base
docker run --net host -v $(pwd):/playground -it --rm ghcr.io/opheadacheh/bazel_build_base:0.2 bash

# run build base and connect to VScode
docker run --net host -v $(pwd):/playground -it -d ghcr.io/opheadacheh/bazel_build_base:0.2 bash