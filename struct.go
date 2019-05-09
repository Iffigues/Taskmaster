package main

type task struct{
	live    int
	com     string
	restart bool
	reboot  int
	code    int
	time    int
	count   int
	signal  int
	stop    int
	stdout  string
	stderr  string
	env     []string
	work    string
	umask   int
}
