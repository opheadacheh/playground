load("@io_bazel_rules_go//go:def.bzl", "go_binary", "go_library")

go_binary(
    name = "client",
    embed = [":client_lib"],
    visibility = ["//visibility:public"],
)


go_library(
    name = "client_lib",
    srcs = ["main.go"],
    importpath = "project/client",
    visibility = ["//visibility:private"],
    deps = [
        "//proto:proto",
        "@org_golang_google_grpc//:go_default_library",
        "@org_golang_x_net//context:go_default_library"
    ],
)
