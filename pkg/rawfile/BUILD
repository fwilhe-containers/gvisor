load("//tools:defs.bzl", "go_library")

package(
    default_applicable_licenses = ["//:license"],
    licenses = ["notice"],
)

go_library(
    name = "rawfile",
    srcs = [
        "blockingpoll_amd64.s",
        "blockingpoll_arm64.s",
        "blockingpoll_noyield_unsafe.go",
        "blockingpoll_yield_unsafe.go",
        "rawfile_unsafe.go",
    ],
    visibility = ["//visibility:public"],
    deps = ["@org_golang_x_sys//unix:go_default_library"],
)
