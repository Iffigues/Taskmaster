all: build

build:
	@make -C Taskmasterd
	@make -C Taskmasterctl

clean:
	@make clean -C Taskmasterd
	@make clean -C Taskmasterctl
