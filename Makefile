
test:
	@echo '=== Running tests for all workspace modules'
	@for mod in $(shell go list -f '{{.Dir}}' -m); do go test -v $$mod/...; done

.PHONY: test mods
