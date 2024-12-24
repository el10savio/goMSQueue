
fmt:
	@echo "go fmt goMSQueue"
	go fmt ./...

vet:
	@echo "go vet goMSQueue"
	go vet ./...

lint:
	@echo "go lint goMSQueue"
	golint ./...

test:
	@echo "Testing goMSQueue"
	go test -v -race --cover ./...

bench:
	@echo "Running benchmarks on goMSQueue"
	go test -bench=. -count=3 -timeout=25m ./...

codespell:
	@echo "checking app spellings"
	codespell
