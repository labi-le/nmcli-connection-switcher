.PHONY: build

PACKAGE_NAME = nmcli-connection-switcher
INSTALL_DIR = /usr/local/bin

build-release:
	go build -ldflags "-s" -a -v -o build/package/$(PACKAGE_NAME) main.go

build:
	go build -v -o build/package/$(PACKAGE_NAME)-debug main.go

install: build-release
	sudo cp build/package/$(PACKAGE_NAME) $(INSTALL_DIR)/$(PACKAGE_NAME)

uninstall:
	sudo rm $(INSTALL_DIR)/$(PACKAGE_NAME)

