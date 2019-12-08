all: build

init:
	make init -C Taskmasterd
	make init -C Taskmasterctl
	mkdir ./log/stderr
	mkdir ./log/stdout

build:
	@make -C Taskmasterd
	@make -C Taskmasterctl

clean:
	make clean -C Taskmasterd
	make clean -C Taskmasterctl

fclean:
	make fclean -C Taskmasterd
	make fclean -C Taskmasterctl

fmt:
	@make fmt -C Taskmasterd
	@make fmt -C Taskmasterctl

push:
	make clean -C Taskmasterd
	make clean -C Taskmasterctl
	git add .
	git commit -m "push"
	git push

fpush:
	make fclean -C Taskmasterd
	make fclean -C Taskmasterctl
	git add .
	git commit -m "push"
	git push
