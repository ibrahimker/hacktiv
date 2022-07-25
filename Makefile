GO_FILES = $(shell go list ./...)

.PHONY: mockgen
mockgen:
	bin/generate-mock.sh $(name)

.PHONY: cleantestcache
cleantestcache:
	go clean -testcache

.PHONY: test.unit
test.unit: cleantestcache
	for s in $(GO_FILES); do if ! go test -failfast -v -race $$s; then break; fi; done
