GOCMD=go
GOINSTALL=$(GOCMD) install
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOGET=$(GOCMD) get

all:  build
build:
	$(GOGET)
	$(GOINSTALL)
	$(GOBUILD)

clean: 
	$(GOCLEAN)
