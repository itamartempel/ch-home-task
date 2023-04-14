.PHONY: fmt
fmt: ## Run go fmt against code.
	go fmt ./...

.PHONY: vet
vet: ## Run go vet against code.
	go vet ./...

.PHONY: test
test: fmt vet ## Run tests.
	go test ./...

.PHONY: gen-testdata
gen-testdata: ## Run tests.
	go run ./testdata/gen-testdata.go

.PHONY: build
build: fmt vet test ## Build binary.
	go build -o bin/ch-home-task cli/main.go