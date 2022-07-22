format:
	@goimports -w ./

build:
	@go build -o ./bin/etcddump main.go


