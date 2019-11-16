GOCMD=go
GOINSTALL=$(GOCMD) install
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOGET=$(GOCMD) get
GOFMT=$(GOCMD) fmt
GOINIT=$(GOCMD) mod init taskmasterd
GOU= ./taskmasterd

all:  build

init:
	$(GOFMT)
	$(GOINIT)

build:
	$(GOFMT)
	$(GOGET)
	$(GOINSTALL)
	$(GOBUILD)

fmt:
	$(GOFMT)

clean: 
	$(GOCLEAN)
	rm -rf $(GOU)
