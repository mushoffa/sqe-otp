ROOT_PATH = $(shell pwd)
DB_PASS=9lgYZWB9Ynwj57syKkowdl7gACLk2M7cSxcPNPCVwAZmhqrVqUqG5A0wLR1Dbvyy

REGEX_SEMVER=[^0-9.]*\([0-9.]*\).*/\1/

STDV=standard-version
STDV_PREVIEW=${STDV} --dry-run --release-as
STDV_RELEASE=${STDV} --release-as
GET_VERSION=sed -n '/release/ s/${REGEX_SEMVER}p'
GIT_BRANCH=$$(git branch --show-current)
VERSION_FILE=config/version.go
VERSION_PACKAGE=config

format:
	gofmt -s -w .; \
	goreportcard-cli

run:
	go run cmd/app/main.go

deploy:
	docker compose  up -d ;\
	go run cmd/app/main.go ;\

test:
	go test $(shell go list ./... | grep -v /mocks) -coverprofile=coverage/coverage.out ;\
	go tool cover -func=coverage/coverage.out ;\
	go tool cover -html=coverage/coverage.out ;\

migration:
	docker run --network=postgres_default --rm \
		-v $(ROOT_PATH)/migrations:/liquibase/changelog liquibase/liquibase \
		--searchPath=/liquibase/changelog/ \
		--changeLogFile=changelog.xml \
		--url="jdbc:postgresql://pg-sqe/sqe?user=admin&password=admin&sslmode=disable" \
		update