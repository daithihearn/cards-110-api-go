help:
	@egrep -h '\s#@\s' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?#@ "}; {printf "\033[36m  %-30s\033[0m %s\n", $$1, $$2}'

docs: #@ Generate docs
	swag init -g cmd/api/main.go
.PHONY:docs
test: fmt vet #@ Run tests
	go test -coverprofile=coverage-full.out ./...
	grep -v "_mocks.go" coverage-full.out | grep -v "handlers.go" | grep -v "collection.go" | grep -v "alexa-validation-service.go" > coverage.out
	go tool cover -html=coverage.out -o coverage.html
.PHONY:test
fmt: #@ Format the code
	go fmt ./...
vet: fmt #@ VET the code
	go vet ./...
lint: fmt #@ Run the linter
	golint ./...
run: test docs vet #@ Start locally
	go run cmd/api/main.go
update: #@ Update dependencies
	go mod tidy
build: test docs vet clear-build copy-translations #@ Build the api and sync binaries
	go build -o build/api/main cmd/api/main.go
.PHONY:build
image: #@ Build docker image
	docker build -t electricity-prices . --load
