all: build
start: goplayground
clean: rm goplayground

build:
	go build
