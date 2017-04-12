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

.PHONY: docker-image
docker-image:
	docker build -t jenkinsci/jnlp-slave-bazel:latest jnlp-slave-bazel
