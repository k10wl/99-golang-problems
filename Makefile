test:
	go test ./...

.PHONY: problem
problem:
	go run main.go -problem $(name)
