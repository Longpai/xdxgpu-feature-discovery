MODULE := github.com/NVIDIA/gpu-feature-discovery

VERSION ?= v0.8.2

vVERSION := v$(VERSION:v%=%)

CUDA_VERSION := 12.2.2
GOLANG_VERSION ?= 1.20.5

GIT_COMMIT ?= $(shell git describe --match="" --dirty --long --always --abbrev=40 2> /dev/null || echo "")
