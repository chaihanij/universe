format:
	go fmt ./...

test:
	export GIN_MODE=release && go test ./... -v -coverprofile .coverage.txt
	go tool cover -func .coverage.txt