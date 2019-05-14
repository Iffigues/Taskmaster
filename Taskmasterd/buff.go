package main

import (
	reader "github.com/chzyer/readline"
)

var completer = reader.NewPrefixCompleter(
	reader.PcItem("login"),
)

func filterInput(r rune) (rune, bool) {
	switch r {
	case reader.CharCtrlZ:
		return r, false
	}
	return r, true
}

func set_read() (ar *reader.Config) {
	ar = &reader.Config{
		Prompt:              "\033[31m»\033[0m ",
		HistoryFile:         "/tmp/readline.tmp",
		AutoComplete:        completer,
		InterruptPrompt:     "",
		EOFPrompt:           "",
		HistorySearchFold:   true,
		FuncFilterInputRune: filterInput,
	}
	return
}
