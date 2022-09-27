OUT_DIR=bin
BIN_DIR=$(shell pwd)/$(OUT_DIR)

.PHONY: deps
# recommend to fixed version.
deps:
ifeq (,$(wildcard $(BIN_DIR)/mockgen))
	@GOBIN=$(BIN_DIR) go install github.com/golang/mock/mockgen@v1.6.0
endif
ifeq (,$(wildcard $(BIN_DIR)/golangci-lint))
	@GOBIN=$(BIN_DIR) go install github.com/golangci/golangci-lint/cmd/golangci-lint@v1.49.0
endif

.PHONY: test
test: generate
	@go test -v -cover -timeout=30m ./...

.PHONY: cover
cover:
	go test -race -covermode=count -coverprofile=coverage.out -run="^Test" ./...

.PHONY: cover.html
cover.html: cover
	@go tool cover -html=coverage.out

.PHONY: lint
lint: deps
	@$(BIN_DIR)/golangci-lint run ./...

.PHONY: generate
generate: deps
	@PATH=$(BIN_DIR):$${PATH} go generate ./...

.PHONY: modtidy
modtidy:
	@go mod tidy

.PHONY: clean
clean:
	@rm -rf $(shell find $(OUT_DIR)/ ! -name '.gitignore')
	@rm -rf $(shell find . -name *.out)
