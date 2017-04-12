.PHONY: all
all: build test

.PHONY: build
build:
	bazel build //...

.PHONY: test
test:
	bazel test //...

.PHONY: gazelle
gazelle:
	gazelle -build_file_name=BUILD.bazel
