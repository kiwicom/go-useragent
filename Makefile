GOLANGCI_LINT := $(shell command -v golangci-lint 2> /dev/null)

# Colorful output
color_off = \033[0m
color_cyan = \033[1;36m
color_green = \033[1;32m

define log_info
	@printf "$(color_cyan)$(1)$(color_off)\n"
endef
define log_success
	@printf "$(color_green)$(1)$(color_off)\n"
endef

test:
	$(call log_info,Run tests and check race conditions)
	# https://golang.org/doc/articles/race_detector.html
	go test -race -v ./...
	$(call log_success,All tests succeeded)

go-mod-tidy:
	$(call log_info,Check that go.mod and go.sum don't contain any unnecessary dependency)
	$(eval TMPDIR=$(shell mktemp -d))
	cp -f go.mod $(TMPDIR)
	cp -f go.sum $(TMPDIR)
	go mod tidy -v
	diff -u $(TMPDIR)/go.mod go.mod
	diff -u $(TMPDIR)/go.sum go.sum
	rm -rf $(TMPDIR)
	$(call log_success,Go mod check succeeded!)

ci: test go-mod-tidy

lint: ## Runs golangci-lint. It outputs to the code-climate json file in if CI is defined.
	$(call log_info, Running golangci-lint)
ifndef GOLANGCI_LINT
	@echo "Can\'t find executable of the golangci-lint. Make sure it is installed. See github.com\/golangci\/golangci-lint#install"
	@exit 1
endif
	golangci-lint run $(if $(CI),--out-format code-climate > gl-code-quality-report.json)
	$(call log_success,Linting with golangci-lint succeeded!)