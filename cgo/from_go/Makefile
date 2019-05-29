CC=gcc
GO=go build

.PHONY: all clean

all: run

build:
	$(CC) -O2 -c simple.c
	ar q libsimple.a simple.o
	CGO_ENABLED=1 $(GO) main.go

run: build 
	@./main

clean:
	@rm -f libsimple.a
	@rm -f simple.o
	@rm -f main