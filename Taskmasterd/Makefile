GOCMD=go
GOINSTALL=$(GOCMD) install
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOGET=$(GOCMD) get
GOFMT=$(GOCMD) fmt
<<<<<<< HEAD
GOU= /root/work/bin/Taskmasterd
all:  build
=======
GOINIT=$(GOCMD) mod init taskmasterd
GOU= ./taskmasterd

all:  build

init:
	$(GOFMT)
	$(GOINIT)

>>>>>>> fd333b86f6e05a39689a5586234a87eeb6377488
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
<<<<<<< HEAD
fclean:
	$(GOCLEAN)
	rm -rf $(GOU) go.sum go.mod
=======
>>>>>>> fd333b86f6e05a39689a5586234a87eeb6377488
