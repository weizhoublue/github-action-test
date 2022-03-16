#!/usr/bin/make -f

include ./Makefile.defs

all: build-bin install-bin

.PHONY: all build install

SUBDIRS := cmd/controller


build-bin:
	for i in $(SUBDIRS); do $(MAKE) $(SUBMAKEOPTS) -C $$i all; done

install-bin:
	$(QUIET)$(INSTALL) -m 0755 -d $(DESTDIR_BIN)
	for i in $(SUBDIRS); do $(MAKE) $(SUBMAKEOPTS) -C $$i install; done

install-bash-completion:
	$(QUIET)$(INSTALL) -m 0755 -d $(DESTDIR_BIN)
	for i in $(SUBDIRS); do $(MAKE) $(SUBMAKEOPTS) -C $$i install-bash-completion; done

clean:
	-$(QUIET) for i in $(SUBDIRS); do $(MAKE) $(SUBMAKEOPTS) -C $$i clean; done
	-$(QUIET) rm -rf $(DESTDIR_BIN)
	-$(QUIET) rm -rf $(DESTDIR_BASH_COMPLETION)

.PHONY: lint
lint: ## Run golangci-lint and check if the helper headers in bpf/mock are up-to-date.
	@$(ECHO_CHECK) golangci-lint
	$(QUIET) golangci-lint run

.PHONY: dev-doctor
dev-doctor:
	$(QUIET) echo "validate local development environment"

.PHONY: update-go-version
update-go-version: ## Update Go version for all the components (images, CI, dev-doctor etc.).
	# ===== Update Go version for GitHub workflow
	$(QUIET) for fl in $(shell find .github/workflows -name "*.yaml" -print) ; do sed -i 's/go-version: .*/go-version: $(GO_IMAGE_VERSION)/g' $$fl ; done
	@echo "Updated go version in GitHub Actions to $(GO_IMAGE_VERSION)"
	# ======= Update Go version in main.go.
	$(QUIET) for fl in $(shell find .  -name main.go -not -path "./vendor/*" -print); do \
		sed -i \
			-e 's|^//go:build go.*|//go:build go$(GO_MAJOR_AND_MINOR_VERSION)|g' \
			-e 's|^// +build go.*|// +build go$(GO_MAJOR_AND_MINOR_VERSION)|g' \
			$$fl ; \
	done
ifeq (${shell [ -f .travis.yml ] && echo done},done)
	# ====== Update Go version in Travis CI config.
	$(QUIET) sed -i 's/go: ".*/go: "$(GO_VERSION)"/g' .travis.yml
	@echo "Updated go version in .travis.yml to $(GO_VERSION)"
endif
	# ======= Update Go version in test scripts.
	$(QUIET) sed -i 's/GO_VERSION=.*/GO_VERSION="$(GO_VERSION)"/g' test/kubernetes-test.sh
	$(QUIET) sed -i 's/GOLANG_VERSION=.*/GOLANG_VERSION="$(GO_VERSION)"/g' test/packet/scripts/install.sh
	@echo "Updated go version in test scripts to $(GO_VERSION)"
	# ===== Update Go version in Dockerfiles.
	$(QUIET) sed -i 's/^go_version=.*/go_version=$(GO_IMAGE_VERSION)/g' images/scripts/update-golang-image.sh
	$(QUIET) $(MAKE) -C images update-golang-image
	@echo "Updated go version in image Dockerfiles to $(GO_IMAGE_VERSION)"
