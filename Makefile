all: build

build:
	@make -C Taskmasterd
	@make -C Taskmasterctl

clean:
	@make clean -C Taskmasterd
	@make clean -C Taskmasterctl
	@rm *.log
	@rm *.pid

fmt:
	@make fmt -C Taskmasterd
	@make fmt -C Taskmasterctl

push:
	$(clean)
	git add .
	git commit -m "push"
	git push
