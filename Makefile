test:
	go test ./...

.PHONY: problem
problem:
	@echo creating problem $(name)
	@go run main.go -problem $(name)
	@echo created files: `ls | grep $(name)`
