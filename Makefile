.PHONY: vet go-test lint ui-test test

vet:
	go vet ./cmd/... ./internal/... ./internal/pdf/...

go-test:
	go test ./cmd/... ./internal/... ./internal/pdf/...

lint:
	npm run lint --prefix internal/ui

ui-test:
	npm test --prefix internal/ui

test: go-test ui-test
