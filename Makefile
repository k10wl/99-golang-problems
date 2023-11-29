.PHONY: problem test

export GODEBUG=randautoseed=0

test:
	go test

problem:
	@echo creating problem $(name)
	@go run main.go -problem $(name)
	@echo created files: `ls | grep $(name)`
