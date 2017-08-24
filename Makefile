.PHONY: all test

all: test bench

test:
	go test ./mux
	go test ./params

bench:
	go test ./mux -bench=.
