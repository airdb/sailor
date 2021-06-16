test:
	go test -v acs/client_test.go
testall:
	#go test -count=1 -v -covermode=count -coverprofile=coverage.out ./...
	#go get github.com/smartystreets/goconvey
	goconvey
