GO=go
GOCOVER=$(GO) tool cover
GOTEST=$(GO) test

.PHONY: test
test:
	$(GO) test ./...
	

.PHONY: test-coverage
test-coverage:
	$(GOTEST) -v ./... -coverprofile=coverage.out -covermode=count -coverpkg=./...
	$(GOCOVER) -html=coverage.out


.PHONY: run
run:
	docker-compose down && docker-compose up --force-recreate

.PHONY: dev
dev:
	docker-compose down && docker-compose up --force-recreate postgres pgadmin

.PHONY: run-dev
run-dev:
	$(GO) run ./cmd/app