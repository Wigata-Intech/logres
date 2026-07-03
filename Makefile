# Logres — build & quality automation.
# Single binary; subcommands select the surface (api / web / webAdmin / migrate).

BINARY      := logres
PKG         := github.com/wigata-intech/logres
MAIN        := .
BIN_DIR     := bin
DIST_DIR    := dist
COVER_FILE  := coverage.out
COVER_MIN   ?= 80

# Version metadata stamped into the binary.
VERSION     ?= $(shell git describe --tags --always --dirty 2>/dev/null || echo dev)
COMMIT      ?= $(shell git rev-parse --short HEAD 2>/dev/null || echo none)
DATE        ?= $(shell date -u +%Y-%m-%dT%H:%M:%SZ)
LDFLAGS     := -s -w -X main.version=$(VERSION) -X main.commit=$(COMMIT) -X main.buildDate=$(DATE)

# Pinned tool versions (kept out of the module graph to keep the binary lean).
GOLANGCI_VERSION   := v2.12.2
GOSEC_VERSION      := latest
GOVULN_VERSION     := latest
MOCKERY_VERSION    := v3.7.1

GO          := go
GOBIN       := $(CURDIR)/$(BIN_DIR)
export GOBIN

# pipefail is needed by the coverage gate.
SHELL       := /bin/bash
.SHELLFLAGS := -o pipefail -c

.DEFAULT_GOAL := help

## help: list targets
.PHONY: help
help:
	@grep -E '^## ' $(MAKEFILE_LIST) | sed 's/## //' | awk -F': ' '{printf "  \033[36m%-16s\033[0m %s\n", $$1, $$2}'

## tidy: sync go.mod / go.sum
.PHONY: tidy
tidy:
	$(GO) mod tidy

## fmt: format code (gofmt + goimports via golangci)
.PHONY: fmt
fmt: $(BIN_DIR)/golangci-lint
	$(BIN_DIR)/golangci-lint fmt

## generate: run go generate + templ
.PHONY: generate
generate:
	$(GO) generate ./...
	$(GO) tool templ generate

## mock: regenerate mocks (mockery v3)
.PHONY: mock
mock:
	$(GO) run github.com/vektra/mockery/v3@$(MOCKERY_VERSION)

## vet: go vet
.PHONY: vet
vet:
	$(GO) vet ./...

## lint: golangci-lint
.PHONY: lint
lint: $(BIN_DIR)/golangci-lint
	$(BIN_DIR)/golangci-lint run ./...

## vuln: govulncheck (known CVEs)
.PHONY: vuln
vuln: $(BIN_DIR)/govulncheck
	$(BIN_DIR)/govulncheck ./...

## sec: gosec static security scan
.PHONY: sec
sec: $(BIN_DIR)/gosec
	$(BIN_DIR)/gosec -quiet -exclude-generated ./...

## test: race-enabled tests with coverage profile
.PHONY: test
test:
	$(GO) test -race -covermode=atomic -coverprofile=$(COVER_FILE) ./...

## cover: gate — every package that HAS tests must be >= COVER_MIN (default 80)
# The tree is still mostly generated stubs, so a global threshold is meaningless;
# instead each tested package must clear the bar. New code ships with tests, and
# tested packages cannot regress. Packages with no tests are reported, not failed.
.PHONY: cover
cover:
	@$(GO) test -race -covermode=atomic -coverprofile=$(COVER_FILE) ./... | \
	awk -v min=$(COVER_MIN) '\
	  { print } \
	  $$1 == "ok" && /coverage: [0-9]/ { \
	    for (i = 1; i <= NF; i++) if ($$i == "coverage:") { p = $$(i+1); sub(/%/, "", p); \
	      if (p + 0 < min + 0) { printf "  ↳ FAIL %s is %s%% (min %s%%)\n", $$2, p, min; rc = 1 } } \
	  } \
	  END { if (rc) { print "coverage gate failed"; exit 1 } print "coverage gate ok (min " min "%)" }'

## cover-html: open HTML coverage report
.PHONY: cover-html
cover-html: test
	$(GO) tool cover -html=$(COVER_FILE)

## ci: full local gate (mirrors GitHub Actions)
.PHONY: ci
ci: tidy vet lint vuln sec cover build
	@echo "CI OK"

## build: build host binary into bin/
.PHONY: build
build:
	CGO_ENABLED=0 $(GO) build -trimpath -ldflags '$(LDFLAGS)' -o $(BIN_DIR)/$(BINARY) $(MAIN)

## run: build then run (pass ARGS="migrate status")
.PHONY: run
run: build
	$(BIN_DIR)/$(BINARY) $(ARGS)

## release: cross-compile for darwin/linux/windows (amd64+arm64)
.PHONY: release
release:
	@mkdir -p $(DIST_DIR)
	$(call xbuild,darwin,amd64)
	$(call xbuild,darwin,arm64)
	$(call xbuild,linux,amd64)
	$(call xbuild,linux,arm64)
	$(call xbuild,windows,amd64)

# xbuild,<os>,<arch> — modernc.org/sqlite is pure Go, so CGO stays off.
define xbuild
	@echo "build $(1)/$(2)"
	@GOOS=$(1) GOARCH=$(2) CGO_ENABLED=0 $(GO) build -trimpath -ldflags '$(LDFLAGS)' \
		-o $(DIST_DIR)/$(BINARY)-$(1)-$(2)$(if $(filter windows,$(1)),.exe,) $(MAIN)
endef

## clean: remove build + coverage artifacts
.PHONY: clean
clean:
	rm -rf $(BIN_DIR) $(DIST_DIR) $(COVER_FILE)

# --- tool installation (pinned, into ./bin) -------------------------------
$(BIN_DIR)/golangci-lint:
	GOBIN=$(GOBIN) $(GO) install github.com/golangci/golangci-lint/v2/cmd/golangci-lint@$(GOLANGCI_VERSION)

$(BIN_DIR)/govulncheck:
	GOBIN=$(GOBIN) $(GO) install golang.org/x/vuln/cmd/govulncheck@$(GOVULN_VERSION)

$(BIN_DIR)/gosec:
	GOBIN=$(GOBIN) $(GO) install github.com/securego/gosec/v2/cmd/gosec@$(GOSEC_VERSION)
