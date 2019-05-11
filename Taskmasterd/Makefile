GOCMD=go
GOINSTALL=$(GOCMD) install
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOGET=$(GOCMD) get
BINARY_SERVER=taskmasterd
BINARY_UNIX=$(BINARY_NAME)_unix

all:  build
build:
	$(GOGET)
	$(GOINSTALL)
	$(GOBUILD)

clean: 
	$(GOCLEAN)
	rm -f $(BINARY_SERVER)
