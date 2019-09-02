export GO111MODULE=on

.PHONY: test
test: lint gofmt


.PHONY: testdeps
testdeps:
	go get -d -v -t ./...
	GO111MODULE=off \
	go get golang.org/x/lint/golint \
		golang.org/x/tools/cmd/cover \
		github.com/axw/gocov/gocov \
		github.com/mattn/goveralls

LINT_RET = .golint.txt
.PHONY: lint
lint: testdeps
	go vet .
	rm -f $(LINT_RET)
	golint ./... | tee $(LINT_RET)
	test ! -s $(LINT_RET)

GOFMT_RET = .gofmt.txt
.PHONY: gofmt
gofmt: testdeps
	rm -f $(GOFMT_RET)
	gofmt -s -d *.go | tee $(GOFMT_RET)
	test ! -s $(GOFMT_RET)

.PHONY: cover
cover: testdeps
	goveralls



## Install dependencies
.PHONY: deps
deps:
	go get -v -d


## Setup build
.PHONY: pre-build
build-deps:
	go get -u github.com/mitchellh/gox


## Build binaries
.PHONY: build
build: build-deps
	gox -osarch="linux/amd64" -output=./bin/ore-mkr_linux-amd64 -ldflags "-s -w"
	gox -osarch="darwin/amd64" -output=./bin/ore-mkr_macOS-amd64 -ldflags "-s -w"
	gox -osarch="windows/amd64" -output=./bin/ore-mkr_windows-amd64 -ldflags "-s -w"
	zip -r ./bin/cross-build.zip ./*
	rm -f ./bin/ore-mkr_*
	gobump show



## Show help
.PHONY: help
help:
	@make2help $(MAKEFILE_LIST)
