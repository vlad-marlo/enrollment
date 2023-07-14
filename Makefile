.DEFAULT_GOAL := build

.PHONY: build
build:
	go build --o server.out cmd/server/main.go
	go build --o enrollment-cli cmd/client/main.go

.PHONY: run
run: build
	./server.out

.PHONY: runm
runm: build
	./server.out -migrate=true

.PHONY: gen
gen:
	swag fmt
	swag init --d cmd/server/,internal/controller/http/,internal/model
	go generate ./...


.PHONY: test
test:
	go test ./... -v -coverpkg=./internal/...,./pkg/... -coverprofile=coverage.out

.PHONY: testshort
testshort:
	go test ./... -v -test.short=true -coverpkg=./internal/...,./pkg/... -coverprofile=coverage.out

.PHONY: c
c:
	go tool cover -func coverage.out

.PHONY: tc
tc: testshort c

.PHONY: lines
lines:
	git ls-files | xargs wc -l
