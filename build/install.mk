OSX_PLuGIN_DIR = ~/Library/Application\ Support/kubebuilder/plugins/kubebuilder-initializer-plugin/v1-alpha
LINUX_PLuGIN_DIR = ~/.config/kubebuilder/plugins/kubebuilder-initializer-plugin/v1-alpha
.PHONY: install
install:
	@echo "=== $(PROJECT_NAME) === [ install ]: install localy... $(UNAME)"
ifeq ($(UNAME), Linux)
	mkdir -p $(LINUX_PLuGIN_DIR)
	cp dist/kubebuilder-initializer-plugin $(LINUX_PLuGIN_DIR)/kubebuilder-initializer-plugin
endif
ifeq ($(UNAME), Darwin)
	mkdir -p $(OSX_PLuGIN_DIR)
	cp dist/kubebuilder-initializer-plugin $(OSX_PLuGIN_DIR)/kubebuilder-initializer-plugin
endif