# Concurrent sum 

## Description
- Project reads json file with structure like [{"a": 0, "b": 0},], and sums all elements using concurrency. Number of goroutines can be set up via args.
- ./cmd/script/main.go is for generating objects.json file. 1_000_000 objects can be generated
- ./cmd/cli/main.go is main program that calculates sum of all objects in generated objects.json file
- look Makefile

## Requirements:
- Golang version 1.18

## Set up
- Clone:
    ```
        git clone https://github.com/asbeeq/sum 
    ```
- Dependencies:
    ```
        go mod tidy
    ```

- Generate objects.json:
    ```
        go run ./cmd/script
    ```

- Test
    ```
        make test
        OR: go test -v ./...
    ```

-  Run:
    ```
        make run
        OR: go run ./cmd/cli -g=1 -f=objects.json -p=concurrent
    ```
