.PHONY: vet go-test lint deps ui-test test

vet:
	go vet ./cmd/... ./internal/... ./internal/pdf/...

go-test:
	go test ./cmd/... ./internal/... ./internal/pdf/...

lint:
	npm run lint --prefix internal/ui

deps:
	npm install --prefix internal/ui

ui-test:
	test -d internal/ui/node_modules || npm install --silent --prefix internal/ui
	npm test --prefix internal/ui

test: go-test ui-test
