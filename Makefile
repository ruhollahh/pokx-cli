run:
	go run . | tee repl.log

test:
	go test ./...