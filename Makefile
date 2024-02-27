.PHONY: all
all:
	go build -v -mod vendor -o chaindata src/main.go

.PHONY: clean
clean:
	go clean -x -cache