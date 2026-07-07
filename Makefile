ATLAS_DATA_REPO ?= mitre-atlas/atlas-data
ATLAS_RELEASE_TAG ?=
UPDATE_ATLAS_DATA ?= 1
HUGO_BASEURL ?=
GITHUB_PAGES_REMOTE ?= origin
GITHUB_PAGES_BRANCH ?= gh-pages

export ATLAS_DATA_REPO ATLAS_RELEASE_TAG UPDATE_ATLAS_DATA HUGO_BASEURL GITHUB_PAGES_REMOTE GITHUB_PAGES_BRANCH

.PHONY: update-data generate build serve deploy clean

update-data:
	scripts/update-atlas-data.sh

generate:
	go run ./cmd/atlasgen

build:
	scripts/build-site.sh

serve: generate
	hugo server -D

deploy:
	scripts/deploy-github-pages.sh

clean:
	rm -rf public resources/_gen
