# suppress output, run `make XXX V=` to be verbose
V := @

.PHONY: test
test: GO_TEST_FLAGS += -race
test:
	$(V)go test -mod=vendor $(GO_TEST_FLAGS) --tags=$(GO_TEST_TAGS) ./...

.PHONY: bench
bench:
	$(V)go test ./... -bench=. -run=NONE

.PHONY: lint
lint:
	$(V)golangci-lint run

.PHONY: vendor
vendor:
	$(V)go mod tidy -compat=1.17
	$(V)go mod vendor
	$(V)git add vendor go.mod go.sum

