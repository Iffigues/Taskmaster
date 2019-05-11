all: build

build:
	@make -C Taskmasterd
	@make -C Taskmasterctl

clean:
	@make clean -C Taskmasterd
	@make clean -C Taskmasterctl

fmt:
	@make fmt -C Taskmasterd
	@make fmt -C Taskmasterctl

push:
	@git add .
	git commit -m "push"
	git push
