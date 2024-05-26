# Stage 4

In this stage we step it up and showcase how to have a unit test and run to see if the test passes. 

```
cc_library(
    name = "hello-time",
    srcs = ["hello-time.cc"],
    hdrs = ["hello-time.h"],
    visibility = ["//main:__pkg__"],
)
```

To use our ```hello-time``` library, an extra dependency is added in the form of //path/to/package:target_name, in this case, it's ```//lib:hello-time```

```
cc_binary(
    name = "hello-world",
    srcs = ["hello-world.cc"],
    deps = [
        ":hello-greet",
        "//lib:hello-time",
    ],
)
```

To run the unit test, use
```
bazelisk test //lib:hello-time-test
```

To build this example, use
```
bazelisk build //main:hello-world
```

To run code coverage, use
```
bazelisk coverage --combined_report=lcov //lib:hello-time-test

genhtml {path_of_dat_file}
```