
REPO = airdb/sailor
all: fmt lint

fmt:
	gofmt -s -w .

exec:
	docker run -it -v $(shell pwd):/go/src/${REPO}/ airdb/golint /bin/bash
lint:
	docker run -it -v $(shell pwd):/go/src/github.com/${REPO} airdb/golint ${REPO}
