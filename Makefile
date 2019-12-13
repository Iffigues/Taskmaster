all: build

init:
	make init -C Taskmasterd
	make init -C Taskmasterctl
	make init -C test

build:
	@make -C Taskmasterd
	@make -C Taskmasterctl
	@make -C test

clean:
	make clean -C Taskmasterd
	make clean -C Taskmasterctl
	make test -C testl

fclean:
	make fclean -C Taskmasterd
	make fclean -C Taskmasterctl
	make fclean -C test

fmt:
	@make fmt -C Taskmasterd
	@make fmt -C Taskmasterctl
	@make fmt -C test

push:
	make clean -C Taskmasterd
	make clean -C Taskmasterctl
	make clean -C test
	git add .
	git commit -m "push"
	git push

fpush:
	make fclean -C Taskmasterd
	make fclean -C Taskmasterctl
	make fclean -C test
	git add .
	git commit -m "push"
	git push
