.PHONY: install
install: fmt
	go install ./cmd/sqlfmt

.PHONY: fmt
fmt:
	go fmt ./...
