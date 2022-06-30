
fmt:
	@echo "go fmt goIntervalSearchTree"	
	go fmt ./...

vet:
	@echo "go vet goIntervalSearchTree"	
	go vet ./...

lint:
	@echo "go lint goIntervalSearchTree"	
	golint ./...

golanglintci:
	@echo "golanglintci goIntervalSearchTree"	
	docker run --rm -v $(shell pwd):/app -w /app golangci/golangci-lint:v1.42.1 golangci-lint run --out-format tab --enable-all

semgrep:
	@echo "semgrep goIntervalSearchTree"	
	docker run --rm -v "$(shell pwd):/src" returntocorp/semgrep --config=auto

test:
	@echo "Testing goIntervalSearchTree"	
	go test -v --cover ./...

codespell:
	@echo "checking goIntervalSearchTree spellings"
	codespell
