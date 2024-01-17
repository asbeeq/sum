GOROUTINES_NUM = 5
FILE_NAME = objects.json
PROCESS_TYPE = concurrent

run:
	go mod tidy
	go run ./cmd/script
	go run ./cmd/cli -g=${GOROUTINES_NUM} -f=${FILE_NAME} -p=${PROCESS_TYPE}

test:
	go test -v ./...