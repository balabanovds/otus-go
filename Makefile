GOROOT=/home/dbalaban/apps/dev/go/current
go=$(GOROOT)/bin/go

build: clean
	GOPATH=$(PWD) $(go) build -o build/app app

clean:
	rm -rf build/

test:
	GOPATH=$(PWD) $(go) test app