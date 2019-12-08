package main

var (
	comopt = map[string]func(){
		"start":    nil,
		"stop":     nil,
		"reread":   nil,
		"quit":     nil,
		"fg":       nil,
		"clear":    nil,
		"maintail": nil,
		"reload":   nil,
		"restart":  nil,
		"shutdown": nil,
		"signal":   nil,
		"status":   nil,
		"tail":     nil,
		"update":   nil,
		"version":  nil,
	}
	comserv   = map[string]func(){}
	comcleint = map[string]func(){}
)
