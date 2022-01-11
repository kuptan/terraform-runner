IMG ?= terraform-runner:latest

docker-build: unit_test
	docker build -t ${IMG} .


.PHONY: fmt
fmt: ## Run go fmt against code.
	go fmt ./...

.PHONY: vet
vet: ## Run go vet against code.
	go vet ./...

build: fmt vet
	go build -o bin/manager main.go

run: build
	go run ./main.go

unit_test:
	go test ./... -coverprofile cover.out

ginkgo:
	${HOME}/go/bin/ginkgo test ./... -coverprofile cover.out