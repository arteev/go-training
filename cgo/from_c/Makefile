CC=gcc
GO=go build

.PHONY: all clean

all: run

build:
	$(GO) -o main . 

run: build 
	@./main

clean:
	@rm -f main