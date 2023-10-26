include .bingo/Variables.mk

.PHONY: drone, test, generate-client

GRAFANA_TARGET_VERSION ?= v10.2.0

drone:
	drone jsonnet --stream --format --source .drone/drone.jsonnet --target .drone/drone.yml
	drone lint .drone/drone.yml
	drone sign --save grafana/grafana-openapi-client-go .drone/drone.yml

test:
	go version
	golangci-lint --version
	golangci-lint run ./...
	go test -cover -race -vet all -mod readonly ./...

##@ Generate Client

ifneq ($(API_TAG),)
CLIENT_GENERATION_ARGS += --tags=$(API_TAG)
$(info Generate client for API with tag '$(API_TAG)'')
endif

ifneq ($(MODEL),)
CLIENT_GENERATION_ARGS += --model=$(MODEL)
$(info Generate model '$(MODEL)')
endif

generate-client: ${SWAGGER}
	rm -rf ./models ./client

	$(SWAGGER) generate client \
	-f https://raw.githubusercontent.com/grafana/grafana/${GRAFANA_TARGET_VERSION}/public/api-merged.json \
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
