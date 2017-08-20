.PHONY: all test

all: test bench

test:
	cd ./routes && go test

bench:
	cd ./routes && go test -bench=.
