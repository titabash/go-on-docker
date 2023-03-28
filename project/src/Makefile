OPENAPI_DIR := ./openapi
MODELS_OUT := ./domain/entity
APIS_OUT := ./web

.PHONY: models
models:
	/usr/local/bin/openapi-generator-cli generate -i $(OPENAPI_DIR)/openapi.yaml -g go \
	--global-property models,supportingFiles=utils.go -o $(MODELS_OUT) --generate-alias-as-model --package-name entity

.PHONY: apis
apis:
	/usr/local/bin/openapi-generator-cli generate -i $(OPENAPI_DIR)/openapi.yaml -g go-server \
	--global-property apis,supportingFiles --additional-properties=onlyInterfaces=true \
	-o $(APIS_OUT) --package-name openapi

.PHONY: build-openapi
build-openapi:
	make models
	make apis
