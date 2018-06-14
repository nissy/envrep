VERSION := 0.0.1
GOOS := $(shell go env GOOS)
GOARCH := $(shell go env GOARCH)

ifeq ($(GOOS),windows)
dist/txtrep-$(VERSION)_$(GOOS)_$(GOARCH).zip:
	go build -ldflags="-X main.version=$(VERSION)" -o dist/txtrep-$(VERSION)/txtrep.exe cmd/txtrep/main.go
	zip -j dist/txtrep-$(VERSION)_$(GOOS)_$(GOARCH).zip dist/txtrep-$(VERSION)/txtrep.exe
else
dist/txtrep-$(VERSION)_$(GOOS)_$(GOARCH).tar.gz:
	go build -ldflags="-X main.version=$(VERSION)" -o dist/txtrep-$(VERSION)/txtrep cmd/txtrep/main.go
	tar cfz dist/txtrep-$(VERSION)_$(GOOS)_$(GOARCH).tar.gz -C dist/txtrep-$(VERSION) txtrep
endif

.PHONY: clean
clean:
	rm -rf dist