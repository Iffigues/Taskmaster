GOCMD=go
GOINSTALL=$(GOCMD) install
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOGET=$(GOCMD) get
FMT=$(GOCMD) fmt

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
