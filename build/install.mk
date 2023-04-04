OSX_PLuGIN_DIR = ~/Library/Application\ Support/kubebuilder/plugins/kubebuilder-layout/v1-alpha
.PHONY: install
install:
	@echo "=== $(PROJECT_NAME) === [ install ]: install localy..."
ifeq ($(UNAME), Linux)
# do something Linux-y
	echo "not supported yet"
endif
ifeq ($(UNAME), Darwin)
	mkdir -p $(OSX_PLuGIN_DIR)
	cp dist/kubebuilder-layout-plugin $(OSX_PLuGIN_DIR)/kubebuilder-layout
endif