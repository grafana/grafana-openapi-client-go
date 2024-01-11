include .bingo/Variables.mk

.PHONY: generate-client, pull-schema, golangci-lint


##@ Generate Client

ifneq ($(API_TAG),)
CLIENT_GENERATION_ARGS += --tags=$(API_TAG)
$(info Generate client for API with tag '$(API_TAG)'')
endif

ifneq ($(MODEL),)
CLIENT_GENERATION_ARGS += --model=$(MODEL)
$(info Generate model '$(MODEL)')
endif

pull-schema:
	bash ./scripts/pull-schema.sh

generate-client: ${SWAGGER} pull-schema
	rm -rf ./models ./client

	$(SWAGGER) generate client \
	-f scripts/schema.json \
	--skip-validation \
	--with-flatten=remove-unused \
	--additional-initialism=DTO \
	--additional-initialism=API \
	--additional-initialism=OK \
	--additional-initialism=LDAP \
	--additional-initialism=ACL \
	--additional-initialism=SNS \
	--additional-initialism=CSV \
	--additional-initialism=PDF \
	--template-dir=templates \
	$(CLIENT_GENERATION_ARGS)
	go mod tidy

golangci-lint:
	docker run \
		--rm \
		--volume "$(shell pwd):/src" \
		--workdir "/src" \
		golangci/golangci-lint:v1.55 golangci-lint run ./... -v
