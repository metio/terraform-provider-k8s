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

out/docs-sentinel: out/tf-format-sentinel $(shell find ./internal -type f -name '*.go') $(shell find ./examples -type f -name '*.tf' -or -name '*.sh') $(shell find ./templates -type f -name '*.tmpl')
	mkdir --parents $(@D)
	go generate ./...
	touch $@

# see https://www.terraform.io/cli/config/config-file#implied-local-mirror-directories
out/install-sentinel: out/${PROVIDER}
	mkdir --parents $(@D)
	mkdir --parents ${XDG_DATA_HOME}/terraform/plugins/localhost/${NAMESPACE}/${NAME}/${VERSION}/${OS_ARCH}
	cp out/${PROVIDER} ${XDG_DATA_HOME}/terraform/plugins/localhost/${NAMESPACE}/${NAME}/${VERSION}/${OS_ARCH}/${PROVIDER}
	touch $@

out/terratests-run-sentinel: out/install-sentinel $(shell find ./terratest -type f -name '*.go') $(shell find ./examples/resources -type f -name '*.tf')
	mkdir --parents $(@D)
	find ./terratest -type f -name 'k8s_*_test.go' | xargs --max-args=1 --max-procs=4 go test
	touch $@

out/tests-sentinel: $(shell find ./internal -type f -name '*.go')
	mkdir --parents $(@D)
	gotestsum --format=testname -- -v -cover -timeout=120s -parallel=4 ./internal/provider
	touch $@

out/coverage.out: $(shell find ./internal -type f -name '*.go')
	mkdir --parents $(@D)
	gotestsum --format=testname -- -v -cover -coverprofile=out/coverage.out -timeout=120s -parallel=4 ./internal/provider

out/coverage.html: out/coverage.out
	go tool cover -html=out/coverage.out -o out/coverage.html

out/go-format-sentinel: $(shell find . -type f -name '*.go')
	mkdir --parents $(@D)
	gofmt -s -w -e .
	touch $@

out/go-lint-sentinel: $(shell find . -type f -name '*.go')
	mkdir --parents $(@D)
	golangci-lint run
	touch $@

out/tf-format-sentinel: $(shell find ./examples -type f -name '*.tf')
	mkdir --parents $(@D)
	terraform fmt -recursive ./examples
	touch $@

.PHONY: install
install: out/install-sentinel ## install the provider locally

.PHONY: docs
docs: out/docs-sentinel ## generate the documentation

.PHONY: terratests
terratests: out/terratests-run-sentinel ## run all terratest tests

.PHONY: terratest
terratest: out/install-sentinel ## run specific terratest tests
	go test -v -timeout=120s -parallel=4 -count=1 -tags testing -run $(filter-out $@,$(MAKECMDGOALS)) ./terratest

.PHONY: tests
tests: out/tests-sentinel ## run the unit tests

.PHONY: test
test: ## run specific unit tests
	go test -v -timeout=120s -tags testing -run $(filter-out $@,$(MAKECMDGOALS)) ./internal/provider

.PHONY: coverage
coverage: out/coverage.html ## generate coverage report

.PHONY: format
format: out/go-format-sentinel out/tf-format-sentinel ## format Go code and Terraform config

.PHONY: lint
lint: out/go-lint-sentinel ## lint all Go code

.PHONY: update
update: ## update all dependencies
	go get -u
	go mod tidy

.PHONY: clean
clean: ## removes all output files
	rm -rf ./out
