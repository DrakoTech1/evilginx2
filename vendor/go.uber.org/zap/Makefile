<<<<<<< HEAD
# Directory containing the Makefile.
PROJECT_ROOT = $(dir $(abspath $(lastword $(MAKEFILE_LIST))))

export GOBIN ?= $(PROJECT_ROOT)/bin
export PATH := $(GOBIN):$(PATH)

GOVULNCHECK = $(GOBIN)/govulncheck
BENCH_FLAGS ?= -cpuprofile=cpu.pprof -memprofile=mem.pprof -benchmem

# Directories containing independent Go modules.
MODULE_DIRS = . ./exp ./benchmarks ./zapgrpc/internal/test

# Directories that we want to track coverage for.
COVER_DIRS = . ./exp
=======
export GOBIN ?= $(shell pwd)/bin

GOLINT = $(GOBIN)/golint
STATICCHECK = $(GOBIN)/staticcheck
BENCH_FLAGS ?= -cpuprofile=cpu.pprof -memprofile=mem.pprof -benchmem

# Directories containing independent Go modules.
#
# We track coverage only for the main module.
MODULE_DIRS = . ./benchmarks ./zapgrpc/internal/test

# Many Go tools take file globs or directories as arguments instead of packages.
GO_FILES := $(shell \
	find . '(' -path '*/.*' -o -path './vendor' ')' -prune \
	-o -name '*.go' -print | cut -b3-)
>>>>>>> deathstrox/main

.PHONY: all
all: lint test

.PHONY: lint
<<<<<<< HEAD
lint: golangci-lint tidy-lint license-lint

.PHONY: golangci-lint
golangci-lint:
	@$(foreach mod,$(MODULE_DIRS), \
		(cd $(mod) && \
		echo "[lint] golangci-lint: $(mod)" && \
		golangci-lint run --path-prefix $(mod)) &&) true

.PHONY: tidy
tidy:
	@$(foreach dir,$(MODULE_DIRS), \
		(cd $(dir) && go mod tidy) &&) true

.PHONY: tidy-lint
tidy-lint:
	@$(foreach mod,$(MODULE_DIRS), \
		(cd $(mod) && \
		echo "[lint] tidy: $(mod)" && \
		go mod tidy && \
		git diff --exit-code -- go.mod go.sum) &&) true


.PHONY: license-lint
license-lint:
	./checklicense.sh

$(GOVULNCHECK):
	cd tools && go install golang.org/x/vuln/cmd/govulncheck
=======
lint: $(GOLINT) $(STATICCHECK)
	@rm -rf lint.log
	@echo "Checking formatting..."
	@gofmt -d -s $(GO_FILES) 2>&1 | tee lint.log
	@echo "Checking vet..."
	@$(foreach dir,$(MODULE_DIRS),(cd $(dir) && go vet ./... 2>&1) &&) true | tee -a lint.log
	@echo "Checking lint..."
	@$(foreach dir,$(MODULE_DIRS),(cd $(dir) && $(GOLINT) ./... 2>&1) &&) true | tee -a lint.log
	@echo "Checking staticcheck..."
	@$(foreach dir,$(MODULE_DIRS),(cd $(dir) && $(STATICCHECK) ./... 2>&1) &&) true | tee -a lint.log
	@echo "Checking for unresolved FIXMEs..."
	@git grep -i fixme | grep -v -e Makefile | tee -a lint.log
	@echo "Checking for license headers..."
	@./checklicense.sh | tee -a lint.log
	@[ ! -s lint.log ]
	@echo "Checking 'go mod tidy'..."
	@make tidy
	@if ! git diff --quiet; then \
		echo "'go mod tidy' resulted in changes or working tree is dirty:"; \
		git --no-pager diff; \
	fi

$(GOLINT):
	cd tools && go install golang.org/x/lint/golint

$(STATICCHECK):
	cd tools && go install honnef.co/go/tools/cmd/staticcheck
>>>>>>> deathstrox/main

.PHONY: test
test:
	@$(foreach dir,$(MODULE_DIRS),(cd $(dir) && go test -race ./...) &&) true

.PHONY: cover
cover:
<<<<<<< HEAD
	@$(foreach dir,$(COVER_DIRS), ( \
		cd $(dir) && \
		go test -race -coverprofile=cover.out -coverpkg=./... ./... \
		&& go tool cover -html=cover.out -o cover.html) &&) true
=======
	go test -race -coverprofile=cover.out -coverpkg=./... ./...
	go tool cover -html=cover.out -o cover.html
>>>>>>> deathstrox/main

.PHONY: bench
BENCH ?= .
bench:
	@$(foreach dir,$(MODULE_DIRS), ( \
		cd $(dir) && \
		go list ./... | xargs -n1 go test -bench=$(BENCH) -run="^$$" $(BENCH_FLAGS) \
	) &&) true

.PHONY: updatereadme
updatereadme:
	rm -f README.md
	cat .readme.tmpl | go run internal/readme/readme.go > README.md

<<<<<<< HEAD
.PHONY: vulncheck
vulncheck: $(GOVULNCHECK)
	$(GOVULNCHECK) ./...
=======
.PHONY: tidy
tidy:
	@$(foreach dir,$(MODULE_DIRS),(cd $(dir) && go mod tidy) &&) true
>>>>>>> deathstrox/main
