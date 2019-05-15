GOCMD=go
GOINSTALL=$(GOCMD) install
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOGET=$(GOCMD) get
GOFMT=$(GOCMD) fmt
GOU= /root/work/bin/Taskmasterd

all:  build

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

