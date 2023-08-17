include .bingo/Variables.mk

.PHONY: drone, test

GRAFANA_TARGET_VERSION ?= v10.0.1

drone:
	drone jsonnet --stream --format --source .drone/drone.jsonnet --target .drone/drone.yml
	drone lint .drone/drone.yml
	drone sign --save grafana/grafana-openapi-client-go .drone/drone.yml

test:
	go version
	golangci-lint --version
	golangci-lint run ./...
	go test -cover -race -vet all -mod readonly ./...

generate-client: ${SWAGGER}
	mkdir -p ./goclient
	$(SWAGGER) generate client \
	--spec https://raw.githubusercontent.com/grafana/grafana/${GRAFANA_TARGET_VERSION}/public/api-merged.json \
	--target ./goclient \
	--with-flatten=remove-unused \
	--additional-initialism=DTO,API,OK,LDAP,ACL,SNS,CSV \
	--tags=folders
	go mod tidy