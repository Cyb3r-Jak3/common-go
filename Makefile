.PHONY: lint test scan

full-test: lint test


lint:
	golangci-lint run --config .golangci-lint.yml ./...

test:
	@gotestsum --format testname --junitfile junit.xml -- -coverprofile=cover.out ./...
	go tool cover -func="cover.out"

scan:
	gosec -no-fail -fmt sarif -out security.sarif ./...

bench:
	go test -bench=. ./...