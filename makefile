build:
	go mod tidy && go install && go build

test:
	go test -v ./...

run:
	make build && make test && ./interview
