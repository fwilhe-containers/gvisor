load("//tools:defs.bzl", "go_library")

package(
    default_applicable_licenses = ["//:license"],
    licenses = ["notice"],
)

go_library(
    name = "k8sctx",
    testonly = True,
    srcs = [
        "k8sctx.go",
    ],
    nogo = False,
    visibility = [
        "//visibility:public",
    ],
    deps = [
        "//test/kubernetes/testcluster",
    ],
)
