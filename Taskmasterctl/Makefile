GOCMD=go
GOINSTALL=$(GOCMD) install
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOGET=$(GOCMD) get
FMT=$(GOCMD) fmt
GOU= /root/work/bin/Taskmasterctl

all:  build

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
	rm *.log
	rm *.pid
