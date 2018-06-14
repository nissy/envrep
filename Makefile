VERSION := 0.0.1
GOOS := $(shell go env GOOS)
GOARCH := $(shell go env GOARCH)

ifeq ($(GOOS),windows)
dist/envrep-$(VERSION)_$(GOOS)_$(GOARCH).zip:
	go build -ldflags="-X main.version=$(VERSION)" -o dist/envrep-$(VERSION)/envrep.exe cmd/envrep/main.go
	zip -j dist/envrep-$(VERSION)_$(GOOS)_$(GOARCH).zip dist/envrep-$(VERSION)/envrep.exe
else
dist/envrep-$(VERSION)_$(GOOS)_$(GOARCH).tar.gz:
	go build -ldflags="-X main.version=$(VERSION)" -o dist/envrep-$(VERSION)/envrep cmd/envrep/main.go
	tar cfz dist/envrep-$(VERSION)_$(GOOS)_$(GOARCH).tar.gz -C dist/envrep-$(VERSION) envrep
endif

.PHONY: clean
clean:
	rm -rf dist