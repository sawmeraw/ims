BINARY_NAME=myapp
SRC_DIR=./cmd/server

.PHONY: all
all: build

.PHONY: build
build:
	go build -o $(BINARY_NAME) $(SRC_DIR)

.PHONY: run
run:build
	./$(BINARY_NAME)

.PHONY: test
test:
	go test ./...

.PHONY: clean
clean:
	rm -f $(BINARY_NAME)
