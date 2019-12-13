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
	mkdir ./log/stdout
	mkdir ./log/stderr
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
fclean:
	$(GOCLEAN)
	rm -rf $(GOU) go.sum go.mod
	rm -rf ./log/stdout
	rm -rf ./log/stderr
	rm ./log/registre.log
