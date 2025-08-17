GO ?= /opt/homebrew/opt/go@1.22/bin/go

.DEFAULT_GOAL := help

help:
	@echo "Targets: env, tidy, vet, test, build, run FILE=path/to/main.go, run-example NAME=examples/.../main.go"

env:
	$(GO) version && $(GO) env GOMOD

tidy:
	$(GO) mod tidy

vet:
	$(GO) vet ./algorithm ./examples/...

test:
	$(GO) test ./algorithm -run .

build:
	$(GO) build ./examples/...

run:
	@test -n "$(FILE)" || (echo "Usage: make run FILE=path/to/main.go" && exit 1)
	$(GO) run $(FILE)

run-example:
	@test -n "$(NAME)" || (echo "Usage: make run-example NAME=examples/path/to/main.go" && exit 1)
	$(GO) run $(NAME)


