GOCMD=go
GOINSTALL=$(GOCMD) install
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOGET=$(GOCMD) get
FMT=$(GOCMD) fmt
GOINIT=$(GOCMD) mod init test
GOU= ./test

all:  build

init:
	$(GOFMT)
	$(GOINIT)

build:
	$(FMT)
	$(GOGET)
	$(GOINSTALL)
	$(GOBUILD)

fmt:
	$(FMT)

clean: 
	$(GOCLEAN)
	rm -rf $(GOU)

fclean:
	$(GOCLEAN)
	rm -rf $(GOU) go.sum go.mod
