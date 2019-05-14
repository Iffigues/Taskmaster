package main

import (
	reader "github.com/chzyer/readline"
)

var completer = reader.NewPrefixCompleter(
	reader.PcItem("login"),
)

func filterInput(r rune) (rune, bool) {
	switch r {
	// block CtrlZ feature
	}
	return r, true
}

func set_read() (ar *reader.Instance, err error) {
	ar, err = reader.NewEx(&reader.Config{
		Prompt:              "\033[31m»\033[0m ",
		HistoryFile:         "/tmp/readline.tmp",
		AutoComplete:        completer,
		InterruptPrompt:     "^C",
		EOFPrompt:           "exit",
		HistorySearchFold:   true,
		FuncFilterInputRune: filterInput,
	})
	return
}
