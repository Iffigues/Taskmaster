package main

import (
	"golang.org/x/crypto/ssh/terminal"
	"io"
	"os"
)

var (
	oldState, errTerm = terminal.MakeRaw(0)
	screen            = struct {
		io.Reader
		io.Writer
	}{os.Stdin, os.Stdout}
	term *terminal.Terminal
)

func is_term() (ok bool) {
	if !terminal.IsTerminal(0) || !terminal.IsTerminal(1) || errTerm != nil {
		return false
	}
	return true
}

func make_term() {
	term = terminal.NewTerminal(screen, "taskmasterctl")
	term.SetPrompt(string(term.Escape.Red) + "> " + string(term.Escape.Reset))
}
