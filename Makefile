TEST?=./...

setup:
	go get github.com/alecthomas/gometalinter
	gometalinter --install

test:
	go test $(TEST) -parallel=4

testrace:
	go test -race $(TEST)

lint:
	gometalinter ./...

watch:
	ginkgo watch -notify -v -cover -succinct -r

.PHONY: setup test testrace lint watch
