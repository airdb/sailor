test:
	go test -v acs/client_test.go

testall:
	#go test -count=1 -v -covermode=count -coverprofile=coverage.out ./...
	#go get github.com/smartystreets/goconvey
	goconvey

db:
	docker-compose -f dbutil/docker-compose.yml up -d
	#go test -count=1 -v -test.short dbutil/dbutil_test.go  -run TestInitDB

stopdb:
	docker-compose -f dbutil/docker-compose.yml stop

bash:
	sudo docker-compose -f dbutil/docker-compose.yml exec testdb bash
