.PHONY: all
all: build
FORCE: ;

SHELL  := env LIBRARY_ENV=$(LIBRARY_ENV) $(SHELL)
LIBRARY_ENV ?= dev

BIN_DIR = $(PWD)/bin

.PHONY: build

clean:
	rm -rf bin/*

dependencies:
	go mod download

build: dependencies build-api build-cmd

build-api:
	go build -tags $(LIBRARY_ENV) -o ./bin/iitd_apiserver.exe cmd/main.go

build-cmd:
	go build -tags $(LIBRARY_ENV) -o ./bin/iitd_cli.exe cmd/iitd_server/main.go

linux-binaries:
	CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -tags "$(LIBRARY_ENV) netgo" -installsuffix netgo -o $(BIN_DIR)/iitd_apiserver api/main.go
	CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -tags "$(LIBRARY_ENV) netgo" -installsuffix netgo -o $(BIN_DIR)/iitd_cli cmd/iitd_server/main.go

ci: dependencies test

build-mocks:
	go get github.com/golang/mock/gomock
	go install github.com/golang/mock/mockgen
	mockgen -source=usecase/student/interface.go -destination=usecase/student/mock/student.go -package=mock

test:
	go test -tags testing ./...

fmt: ## gofmt and goimports all go files
	find . -name '*.go' -not -wholename './vendor/*' | while read -r file; do gofmt -w -s "$$file"; goimports -w "$$file"; done