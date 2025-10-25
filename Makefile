default: build

clean:
	rm -rf ./bin

build: clean
	go build -o bin/pocketmoney ./cmd

run: build
	./bin/pocketmoney

test_unit:
	go test ./internal/...

test_api:
	go test ./api_tests/...

test: test_unit test_api
