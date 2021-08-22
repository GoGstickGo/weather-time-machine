GO := CGO_ENABLED=0 go
GO_TAGS ?=
TARGET=wtmWeb
INSTALL = $(QUIET)install

$(TARGET):
	$(GO) build \
		-o $(TARGET) \
		./main.go

install: $(TARGET)

clean:
	rm -f $(TARGET)
