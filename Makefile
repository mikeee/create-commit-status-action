RELEASEDIR := ./
RELEASENAME := main

# populate release version
RELEASEVERSION := release

BUILD_MATRIX := darwin-amd64 darwin-arm64 linux-amd64 linux-arm64 windows-amd64 windows-arm64

$(BUILD_MATRIX):
	@echo "Building for $(@)"
	CGO_ENABLED=0 GOOS=$(shell echo $@ | cut -d'-' -f1) GOARCH=$(shell echo $@ | cut -d'-' -f2) go build -ldflags="-s -w -X 'main.Version=$(RELEASEVERSION)'" -installsuffix static -o "$(RELEASEDIR)$(RELEASENAME)-$(shell echo $@ | cut -d'-' -f1)-$(shell echo $@ | cut -d'-' -f2)" ./cmd/createcommitstatus/

.PHONY: release
release: $(BUILD_MATRIX)
	@echo Release builds complete

# windows-arm64 is currently skipped as unsupported by upx
.PHONY: compress-releases
RELEASE_FILES := $(wildcard $(RELEASEDIR)$(RELEASENAME)-*)
compress-releases:
	@for f in $(RELEASE_FILES); do \
		if [[ ! "$$f" =~ "windows-arm64" ]] && [[ ! "$$f" =~ "darwin-arm64" ]]; then \
			upx -q -9 --force-macos $$f; \
		fi; \
	done
	@echo Release compression completed


.PHONY: cleanup
cleanup:
	rm -rf $(RELEASEDIR)
