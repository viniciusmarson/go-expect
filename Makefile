PACKAGE=`pwd`/expect

all: test

test:
	@GOPATH=`pwd` go test ./expect -v