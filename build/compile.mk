BUILD_DIR       ?= $(CURDIR)/dist
APP_CMD         ?= cmd/main.go
COMMIT = $(shell git log --pretty=format:'%H' -n 1)
BUILD_DATE = $(shell date -u +%Y-%m-%dT%H:%M:%SZ)
VERSION  ?= $(shell git describe --tags --always --dirty | sed -e '/^v/s/^v\(.*\)$$/\1/g')
VERSION_TAG  := $(shell git describe --tags --always --abbrev=0 | sed -e '/^v/s/^v\(.*\)$$/\1/g')

LDFLAGS   =  -s -w

.PHONY: build
build: ## build executable for the current os
	@echo "=== $(PROJECT_NAME) === [ build          ]: building commands:"
	@mkdir -p $(BUILD_DIR)/$(GOOS)
	@$(GO_CMD) build -ldflags="$(LDFLAGS)" -o $(BUILD_DIR)/$(GOOS)/$(PROJECT_NAME) $(APP_CMD)

.PHONY: build-all
build-all: build-linux build-darwin build-windows

.PHONY: build-linux
build-linux: ## build executable for linux
	@echo "=== $(PROJECT_NAME) === [ build-linux    ]: building executable for linux..."
	@mkdir -p $(BUILD_DIR)/linux
	@GOOS=linux $(GO_CMD) build -ldflags="$(LDFLAGS)" -o $(BUILD_DIR)/linux/$(PROJECT_NAME) $(APP_CMD)

.PHONY: build-darwin
build-darwin: ## build executable for darwin
	@echo "=== $(PROJECT_NAME) === [ build-darwin   ]: building executable for darwin..."
	@mkdir -p $(BUILD_DIR)/darwin
	@GOOS=darwin $(GO_CMD) build -ldflags="$(LDFLAGS)" -o $(BUILD_DIR)/darwin/$(PROJECT_NAME) $(APP_CMD)

.PHONY: build-windows
build-windows: ## build executable for windows
	@echo "=== $(PROJECT_NAME) === [ build-windows  ]: building executable for windows..."
	@mkdir -p $(BUILD_DIR)/windows
	@GOOS=windows $(GO_CMD) build -ldflags="$(LDFLAGS)" -o $(BUILD_DIR)/windows/$(PROJECT_NAME).exe $(APP_CMD)

.PHONY: build-windows32
build-windows32: ## build executable for windows32
	@echo "=== $(PROJECT_NAME) === [ build-windows  ]: building executable for windows32..."
	@mkdir -p $(BUILD_DIR)/windows
	GOARCH=386 CGO_ENABLED=1 GOOS=windows $(GO_CMD) build -ldflags="$(LDFLAGS)" -o $(BUILD_DIR)/windows/$(PROJECT_NAME) $(APP_CMD)

.PHONY: build-clean
build-clean: ## remove compiled files
	@echo "=== $(PROJECT_NAME) === [ build-clean  ]: removing compiled files..."
	@rm -rfv $(BUILD_DIR)/*