GITCOMMIT := $(shell git rev-parse HEAD)
GITDATE := $(shell git show -s --format='%ct')

LDFLAGSSTRING +=-X main.GitCommit=$(GITCOMMIT)
LDFLAGSSTRING +=-X main.GitDate=$(GITDATE)
LDFLAGS := -ldflags "$(LDFLAGSSTRING)"

hsm-sign-service:
	go mod tidy
	env GO111MODULE=on go build -v $(LDFLAGS) ./cmd/hsm-sign-service

clean:
	rm hsm-sign-service

test:
	go test -v ./...

lint:
	golangci-lint run ./...


.PHONY: \
	hsm-sign-service \
	binding \
	clean \
	test \
	lint

