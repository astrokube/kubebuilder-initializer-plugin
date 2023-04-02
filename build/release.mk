.PHONY: release
release:
	@echo "=== $(PROJECT_NAME) === [ release ]: running go release..."
	goreleaser --clean