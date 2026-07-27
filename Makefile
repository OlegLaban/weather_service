.PHONY: compile

compile:
	go build -o ./build/rest ./cmd/rest/main.go
