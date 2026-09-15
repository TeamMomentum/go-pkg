# copy from https://github.com/dictav/go-genproto-googleads
# and changed by TeamMomentum

PKG := github.com/TeamMomentum/go-pkg/go-genproto-googleads
GOOGLE_PROTO := google.golang.org/genproto/googleapis/ads/googleads

TARGETS := v25
BAZEL_BIN := $(shell cd googleapis; bazel info bazel-bin)
SRC := $(BAZEL_BIN)/google/ads/googleads/$(VERSION)/gapi-ads-googleads-$(VERSION)-go.tar.gz

all: $(TARGETS)

$(TARGETS): googleapis
	-rm -rf build
	$(MAKE) gen VERSION=$@

gen: build
	-rm -rf $(VERSION) pb/$(VERSION)
	cp -a build/$(PKG)/$(VERSION) $(VERSION)
	cp -a build/$(GOOGLE_PROTO)/$(VERSION) pb/$(VERSION)
	$(MAKE) verify VERSION=$(VERSION)

# Fails the build if generated code imports a pb/$(VERSION) package that
# doesn't actually exist (e.g. a new proto package added upstream that
# googleads-go.patch.sh doesn't yet wire up a go_proto_library for).
.PHONY: verify
verify:
	@echo "verifying pb/$(VERSION) against its own imports ..."
	@status=0; \
	for pkg in $$(grep -oh '$(PKG)/pb/$(VERSION)/[A-Za-z0-9_]*' -r $(VERSION) pb/$(VERSION) | sed -E 's#.*/pb/$(VERSION)/##' | sort -u); do \
		if [ ! -d "pb/$(VERSION)/$$pkg" ]; then \
			echo "ERROR: pb/$(VERSION)/$$pkg is imported but was not generated (update googleads-go.patch.sh)"; \
			status=1; \
		fi; \
	done; \
	exit $$status

build: $(SRC)
	@echo "Using BAZEL_BIN=$(BAZEL_BIN)"
	mkdir -p tmp
	tar zxf $(SRC) -C tmp
	grep -l -R '$(GOOGLE_PROTO)' tmp/ | \
		xargs sed -i -e 's,$(GOOGLE_PROTO),$(PKG)/pb,'
	mv tmp/$(GOOGLE_PROTO)/$(VERSION)/resources/experiment_arm.pb.go \
		tmp/$(GOOGLE_PROTO)/$(VERSION)/resources/experimentarm.pb.go
	mv tmp build

$(SRC): googleapis
	cd googleapis; git reset --hard
	./googleads-go.patch.sh $(PKG) $(VERSION)
	cd googleapis; bazel build --cxxopt=-std=c++17 --host_cxxopt=-std=c++17 //google/ads/googleads/$(VERSION):gapi-ads-googleads-$(VERSION)-go

.PHONY: googleapis
googleapis:
	if [ -d googleapis/.git ]; then \
		cd googleapis && git fetch --depth=1 origin && git reset --hard origin/HEAD; \
	else \
		git clone --depth=1 https://github.com/googleapis/googleapis; \
	fi

clean:
	-rm -rf tmp
	-rm -rf build
	-rm -rf googleapis
