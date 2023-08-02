RELEASE_TYPE ?= patch
BRANCH ?= null
#CURRENT_VERSION := $(shell git ls-remote --tags | awk '{ print $$2}'| sort -nr | head -n1|sed 's/refs\/tags\///g')
#CURRENT_VERSION := $(shell git describe --tags `git rev-list --tags --max-count=1`)
CURRENT_VERSION := $(shell git ls-remote --sort=-version:refname -q --tags | awk '{ print $$2}' | head -n1|sed 's/refs\/tags\///g')
ifndef CURRENT_VERSION
  CURRENT_VERSION := 0.0.0
endif

NEXT_VERSION := $(shell docker run --rm alpine/semver semver -c -i $(RELEASE_TYPE) $(CURRENT_VERSION))

current-version:
	@echo $(CURRENT_VERSION)

next-version:
	@echo $(NEXT_VERSION)

release:
	git checkout $(BRANCH);
	git config credential.helper store
	git pull
	git tag $(NEXT_VERSION)-$(BRANCH)
	git push --tags
	git checkout $(NEXT_VERSION)-$(BRANCH)