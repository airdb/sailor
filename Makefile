test:
	docker compose up -d
	#go test -count=1 -v -test.short dbutil/dbutil_test.go  -run TestInitDB
	go test -count=1 -v -covermode=count -coverprofile=coverage.out ./...
	docker compose stop

testall:
	#go get github.com/smartystreets/goconvey
	#goconvey

bash:
	docker compose exec testdb bash

redis:
	docker-compose up -d
	go test -count=1 -v -test.short redisutil/redis_test.go -run TestNewRedisClient