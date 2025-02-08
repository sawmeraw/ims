BINARY_NAME=bin/myapp
SRC_DIR=./cmd/server
CSS_INPUT=./assets/css/input.css
CSS_OUTPUT=./assets/css/styles.css

.PHONY: all
all: build

.PHONY: build
build: generate-templates generate-css
	go build -o $(BINARY_NAME) $(SRC_DIR)

.PHONY: run
run: build
	./$(BINARY_NAME)

.PHONY: test
test:
	go test ./...

.PHONY: clean
clean:
	rm -f $(BINARY_NAME)
	rm -f $(CSS_OUTPUT)

.PHONY: generate-templates
generate-templates:
	templ generate

.PHONY: generate-css
generate-css:
	tailwindcss -i $(CSS_INPUT) -o $(CSS_OUTPUT)

.PHONY: dev
dev: generate-templates
	make watch-css & air

.PHONY: watch-css
watch-css:
	tailwindcss -i $(CSS_INPUT) -o $(CSS_OUTPUT) --watch

.PHONY: watch-templates
watch-templates:
	templ generate --watch ./internal/views/...