LINTER_HOME ?= /tmp/go/lint/${ORG}/${NAME}
LINTER_VERSION  ?= v1.39.0


.PHONY: lint
lint:
	@mkdir -p ${LINTER_HOME}
	@docker run --rm \
	  -v ${PWD}:/app \
	  -v ${LINTER_HOME}:/root \
	  -e GOLANGCI_LINT_CACHE=/root/lint/cache \
	  -w /app \
	  golangci/golangci-lint:${LINTER_VERSION} golangci-lint run -v

.PHONY: swagger
swagger:
	@swagger generate spec -o ./api/docs/swagger.json --scan-models