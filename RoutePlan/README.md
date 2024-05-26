# Prerequisites

1. A linux host machine
2. VS code
3. Docker

# Instructions

## Recommended setup:

1. VS code as the IDE.
2. Use docker container as your build environment, on your host machine, under directory playground, run command below
```
docker run -v $(pwd):/playground -it -d registry.jihulab.com/opheadacheh/playground/bazel_build_base:0.1 bash
```
3. Install "Remote Development" extension inside VS code, use "Remote Explorer" to enter the container you just created in step 2.
4. Enter the work directory, all the command running inside the container should be executed under RoutePlan/
```
cd RoutePlan
```

## Build a new server version and push to registry

1. Modify the image tag under "tarball" target in server/BUILD.bazel file. For example: 
```
oci_tarball(
    name = "tarball",
    # Use the image built for the exec platform rather than the target platform
    image = ":image",
    repo_tags = ["registry.jihulab.com/opheadacheh/playground/route_plan_server:20231030.1"],
)
```

2. Run command below to generate the image file.
```
bazelisk build //server:tarball
```

3. Run command below to cp the generated image file to host machine, so that host machine can load with Docker.
```
cp bazel-bin/server/tarball/tarball.tar tarball.tar
```

4. Run command below on your host machine to load the image
```
docker load --input RoutePlan/tarball.tar
```

5. Push the image to registry
```
docker push registry.jihulab.com/opheadacheh/playground/route_plan_server:20231030.1
```

## Run the newly created server on the cloud machine
1. Connect to the cloud machine
2. Stop the original docker container
```
docker ps
docker stop {container id}
```
3. Start a new container with the new image, -p for port forwarding, -v for certificates.
```
docker run -d -p 443:8080 -v /root:/root registry.jihulab.com/opheadacheh/playground/route_plan_server:20231030.1
```