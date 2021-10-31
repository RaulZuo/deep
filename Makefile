# Build all by default, even if it's not first
.DEFAULT_GOAL := all

.PHONY: all
all: gen add-copyright format lint cover build

# ==============================================================================
# Build options

ROOT_PACKAGE=github.com/RaulZuo/deep
VERSION_PACKAGE=github.com/marmotedu/component-base/pkg/version

# ==============================================================================
# Includes

include scripts/make-rules/common.mk # make sure include common.mk at the first include line
include scripts/make-rules/golang.mk

# build: Build source code for host platform
.PHONY: build
build:
	@$(MAKE) go.build