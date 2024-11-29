GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
GOMOD=$(GOCMD) mod
SRCFOLDER = src
BINFOLDER = bin
BINARY_NAME=main

all: build

build: binFolder
	cd $(SRCFOLDER) && $(GOBUILD) -o ../$(BINFOLDER)/$(BINARY_NAME) -v

clean: 
	$(GOCLEAN)
	rm -rf $(BINFOLDER)
	rm -rf test
	rm -rf main

binFolder:
	mkdir -p $(BINFOLDER)

run:
	./$(BINFOLDER)/$(BINARY_NAME)

tidy:
	cd $(SRCFOLDER) && $(GOMOD) tidy

mod:
	$(GOMOD) tidy

.PHONY: all build test clean run mod
