CC=gcc
GO=go build

.PHONY: all clean

all: run

build:
	$(CC) -O2 -c clibrary.c
	ar q libclibrary.a clibrary.o
	CGO_ENABLED=1 $(GO) main.go c_func.go

run: build 
	@./main

clean:
	@rm -f libclibrary.a
	@rm -f clibrary.o
	@rm -f main