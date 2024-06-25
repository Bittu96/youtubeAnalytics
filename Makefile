BUILD_DIR := bin
TOOLS_DIR := tools

default: all

all: clean lint test build 

.PHONY: $(BUILD_DIR)/publisher
bin/publisher: cmd/publisher/*.go
	CGO_ENABLED=0 go build -mod vendor -ldflags="-s -w" -o ./bin/publisher ./cmd/publisher/

.PHONY: $(BUILD_DIR)/consumer
bin/consumer: cmd/consumer/*.go
	CGO_ENABLED=0 go build -mod vendor -ldflags="-s -w" -o ./bin/consumer ./cmd/consumer/

.PHONY: build
build: bin/publisher
build: bin/consumer

.PHONY: clean
clean:
	rm -rf $(BUILD_DIR)
	rm -rf $(TOOLS_DIR)
	@go mod vendor
	@go mod tidy

.PHONY: runPublisher
runPublisher: build
	bin/publisher

.PHONY: runConsumer
runConsumer: build
	bin/consumer

tools/golangci-lint/golangci-lint:
	mkdir -p tools/
	curl -sfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh| sh -s -- -b tools/golangci-lint latest

.PHONY: lint
lint: $(TOOLS_DIR)/golangci-lint/golangci-lint
	./$(TOOLS_DIR)/golangci-lint/golangci-lint run ./...

.PHONY: test
test:
	go test -mod vendor -race -cover -coverprofile=coverage.txt -covermode=atomic ./...
