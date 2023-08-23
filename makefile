.PHONY: test

test:
	@echo "---- RUNNING UNIT TEST UTILS FOLDER ----\n"
	go test -v ./internal/utils/
	@echo "\n---- RUNNING UNIT TEST TESTS/UNIT FOLDER ----\n"
	go test -v ./tests/unit/
	@echo "\n---- RUNNING UNIT TEST TESTS/INTEGRATION FOLDER ----\n"
	go test -v ./tests/integration/
	go clean -testcache

linter:
	@echo "---- START LINTER GOLANG ----\n"
	golangci-lint version
	golangci-lint run -c .golangci.yml ./...

go-generate:
	go generate ./...