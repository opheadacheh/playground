# playground

This is Wanjia's coding playground

# build and push docker image
export TAG=ghcr.io/opheadacheh/bazel_build_base:0.3
docker build -f BuildBaseDockerImages/v0.3/Dockerfile . -t $TAG && docker push $TAG

# enter build base
docker run -v $(pwd):/playground -it --rm ghcr.io/opheadacheh/bazel_build_base:0.3