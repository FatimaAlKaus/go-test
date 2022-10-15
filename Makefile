PACKAGES := $(shell go list ./... | grep -v /vendor/)


.PHONY: test
test: ## run unit tests
	@mkdir -p coverage
	@echo "mode: count" > coverage/coverage-all.out
	@$(foreach pkg,$(PACKAGES), \
		pk=$(shell echo ${pkg} | awk -F/ '{print $$NF}'); \
		echo $$pk;\
		go test -p=1 -cover -covermode=count -coverprofile=coverage/coverage_$$pk.out ${pkg}; \
		tail -n +2 coverage/coverage_$$pk.out >> coverage/coverage-all.out;)
.PHONY: cover
cover:
	go tool cover -html=coverage/coverage-all.out