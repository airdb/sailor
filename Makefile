
REPO = airdb/sailor
IMAGE = airdb/golint
all: fmt lint

fmt:
	gofmt -s -w .

exec:
	docker run -it -v $(shell pwd):/go/src/${REPO}/ airdb/golint /bin/bash
lint:
	docker pull ${IMAGE}
	docker run -it -v $(shell pwd):/go/src/github.com/${REPO} ${IMAGE} ${REPO}
