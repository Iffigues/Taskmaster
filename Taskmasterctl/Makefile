GOCMD=go
GOINSTALL=$(GOCMD) install
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOGET=$(GOCMD) get
FMT=$(GOCMD) fmt
GOINIT=$(GOCMD) mod init taskmasterctl
GOU= /root/work/bin/Taskmasterctl

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
