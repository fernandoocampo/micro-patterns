lint-docker:
	docker run --rm -v $(PWD):/app -w /app golangci/golangci-lint:v1.48-alpine golangci-lint run