build:
	@go build -o out/grape

install:
	@go install

test:
	@go test ./...

clean:
	@rm -rf out/