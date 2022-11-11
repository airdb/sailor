test:
	#go test -count=1 -v -test.short dbutil/dbutil_test.go  -run TestInitDB
	go test -count=1 -v -covermode=count -coverprofile=coverage.out ./...

testdb:
	#         docker compose up -d --build --force-recreate
	docker compose up -d mysql --remove-orphans
	go test -count=1 -v -test.short dbutil/dbutil_test.go  -run TestInitDB
	docker compose stop mysql

testall:
	#go get github.com/smartystreets/goconvey
	#goconvey

lint:
	go fmt ./...
	golangci-lint run

bash:
	docker compose exec testdb bash

redis:
	docker-compose up -d
	go test -count=1 -v -test.short redisutil/redis_test.go -run TestNewRedisClient
