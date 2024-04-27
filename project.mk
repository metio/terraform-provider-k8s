# SPDX-FileCopyrightText: The terraform-provider-k8s Authors
# SPDX-License-Identifier: 0BSD

NAMESPACE     = metio
NAME          = k8s
PROVIDER      = terraform-provider-${NAME}
VERSION       = 9999.99.99
OS_ARCH       ?= linux_amd64
XDG_DATA_HOME ?= ~/.local/share

out/${PROVIDER}: $(shell find ./internal -type f -name '*.go' -and -not -name '*test.go')
	mkdir --parents $(@D)
	go build -o out/${PROVIDER}

out/fetcher-sentinel: $(shell find ./tools/fetcher -type f -name '*.go') $(shell find ./tools/internal/fetcher -type f -name '*.go')
	mkdir --parents $(@D)
	go generate ./tools/fetch.go
	touch $@

out/generate-sentinel: $(shell find ./tools/generator -type f -name '*.go') $(shell find ./tools/internal/generator -type f -name '*.go') $(shell find ./tools/internal/generator/templates -type f -name '*.tmpl') $(shell find ./schemas/crd_v1 -type f -name '*.yaml') $(shell find ./schemas/openapi_v2 -type f -name '*.json')
	mkdir --parents $(@D)
	go generate ./tools/codegen.go
	touch $@

out/docs-sentinel: out/generate-sentinel out/tf-format-sentinel $(shell find ./internal -type f -name '*.go') $(shell find ./examples -type f -name '*.tf' -or -name '*.sh') $(shell find ./templates -type f -name '*.tmpl')
	mkdir --parents $(@D)
	go generate ./tools/docs.go
	touch $@

# see https://www.terraform.io/cli/config/config-file#implied-local-mirror-directories
out/install-sentinel: out/${PROVIDER}
	mkdir --parents $(@D)
	mkdir --parents ${XDG_DATA_HOME}/terraform/plugins/localhost/${NAMESPACE}/${NAME}/${VERSION}/${OS_ARCH}
	cp out/${PROVIDER} ${XDG_DATA_HOME}/terraform/plugins/localhost/${NAMESPACE}/${NAME}/${VERSION}/${OS_ARCH}/${PROVIDER}
	touch $@

out/tools-tests-sentinel: $(shell find ./tools -type f -name '*.go')
	mkdir --parents $(@D).
	gotestsum --format=testname -- -timeout=120s ./tools/internal/...
	touch $@

out/coverage.out: $(shell find ./internal -type f -name '*.go')
	mkdir --parents $(@D)
	gotestsum --format=testname -- -v -cover -coverprofile=out/coverage.out -timeout=120s -parallel=4 ./internal/...

out/coverage.html: out/coverage.out
	go tool cover -html=out/coverage.out -o out/coverage.html

out/go-format-sentinel: $(shell find . -type f -name '*.go')
	mkdir --parents $(@D)
	gofmt -s -w -e .
	touch $@

out/go-vet-sentinel: $(shell find . -type f -name '*.go')
	mkdir --parents $(@D)
	go vet
	touch $@

out/tf-format-sentinel: $(shell find ./examples -type f -name '*.tf')
	mkdir --parents $(@D)
	terraform fmt -recursive ./examples
	touch $@

.PHONY: install
install: out/install-sentinel ## install the provider locally

.PHONY: fetch
fetch: out/fetcher-sentinel ## fetch upstream specs

.PHONY: generate
generate: out/generate-sentinel ## generate the code

.PHONY: docs
docs: out/docs-sentinel ## generate the documentation

.PHONY: coverage
coverage: out/coverage.html ## generate coverage report

.PHONY: format
format: out/go-format-sentinel out/tf-format-sentinel ## format Go code and Terraform config

.PHONY: vet
vet: out/go-vet-sentinel ## vet all Go code

.PHONY: update
update: ## update all dependencies
	go get -u
	go mod tidy

.PHONY: clean
clean: ## removes all output files
	rm -rf ./out

-include terratests.mk
-include tests.mk
