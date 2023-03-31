.PHONY: docker-%
docker-%: ## Run commands inside a docker container
	docker run --rm --workdir /app -v $(CURDIR):/app golang:$(GO_VERSION) \
	make $*