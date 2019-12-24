.PHONY: build clean contrib_check coverage help install isntall lint run size test uninstall

# detect GOPATH if not set
ifndef $(GOPATH)
	$(info GOPATH is not set, autodetecting..)
	TESTPATH := $(dir $(abspath ../../..))
	DIRS := bin pkg src

	# create a ; separated line of tests and pass it to shell
	MISSING_DIRS := $(shell $(foreach entry,$(DIRS),test -d "$(TESTPATH)$(entry)" || echo "$(entry)";))
	ifeq ($(MISSING_DIRS),)
		$(info Found GOPATH: $(TESTPATH))
		export GOPATH := $(TESTPATH)
	else
		$(info ..missing dirs "$(MISSING_DIRS)" in "$(TESTDIR)")
		$(info GOPATH autodetection failed)
	endif
endif

# Set go modules to on and use GoCenter for immutable modules
export GO111MODULE = on
export GOPROXY = https://proxy.golang.org,direct

# Determines the path to this Makefile
THIS_FILE := $(lastword $(MAKEFILE_LIST))

APP=taskhelper

# -------------------- Actions -------------------- # 

## build: builds a local version
build:
	go build -o bin/${APP} -mod=vendor
	@echo "Done building"

run: build
	bin/${APP} --config ./taskhelper2.yaml
