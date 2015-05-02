TEST?=./...

default: test

setup:
	go get github.com/onsi/ginkgo/ginkgo
	go get github.com/alecthomas/gometalinter
	gometalinter --install

test:
	ginkgo -r -cover

testrace:
	go test -race $(TEST)

lint:
	gometalinter ./...

watch:
	ginkgo watch -notify -v -cover -succinct -r

.PHONY: setup test testrace lint watch
