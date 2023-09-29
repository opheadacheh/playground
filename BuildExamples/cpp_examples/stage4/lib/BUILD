load("@rules_cc//cc:defs.bzl", "cc_library")

cc_library(
    name = "hello-time",
    srcs = ["hello-time.cc"],
    hdrs = ["hello-time.h"],
    visibility = [
        "//main:__pkg__",
        "//lib:__pkg__"
    ],
)

cc_test(
    name = "hello-time-test",
    srcs = ["hello-time_test.cc"],
    copts = ["-Iexternal/gtest/include"],
    deps = [
        "@com_google_googletest//:gtest_main",
        "//lib:hello-time",
    ],
)