GO := CGO_ENABLED=0 go
GO_TAGS ?=
TARGET=wtm
INSTALL = $(QUIET)install
BINDIR ?= /usr/local/bin

TEST_TIMEOUT ?= 5s

GOLANGCILINT_WANT_VERSION = 1.40.1
GOLANGCILINT_VERSION = $(shell golangci-lint version 2>/dev/null)

$(TARGET):
	$(GO) build \
		-o $(TARGET) \
		./cmd/wtm

install: $(TARGET)
	$(INSTALL) -m 0755 -d $(DESTDIR)$(BINDIR)
	$(INSTALL) -m 0755 $(TARGET) $(DESTDIR)$(BINDIR)

clean:
	rm -f $(TARGET)
	rm -rf ./release

test:
	go test -timeout=$(TEST_TIMEOUT) -race -cover $$(go list ./...)

bench:
	go test -timeout=30s -bench=. $$(go list ./...)

ifneq (,$(findstring $(GOLANGCILINT_WANT_VERSION),$(GOLANGCILINT_VERSION)))
check:
	golangci-lint run
endif