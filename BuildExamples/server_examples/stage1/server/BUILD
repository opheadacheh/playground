load("@io_bazel_rules_go//go:def.bzl", "go_binary", "go_library")

go_binary(
    name = "server",
    embed = [":server_lib"],
    visibility = ["//visibility:public"],
)

go_library(
    name = "server_lib",
    srcs = ["main.go"],
    importpath = "project/server",
    visibility = ["//visibility:private"],
    deps = [
        "//proto:proto",
        "@org_golang_google_grpc//:go_default_library"
    ],
)
