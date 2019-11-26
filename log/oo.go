package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"os/exec"
)

func main() {
	cmd := exec.Command("cat")
	// create a pipe for the output of the script
	cmdReader, err := cmd.StdoutPipe()
	errReader, err := cmd.StderrPipe()
	stdin, err := cmd.StdinPipe()
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error creating StdoutPipe for Cmd", err)
		return
	}
	go func() {
		scanner := bufio.NewScanner(cmdReader)
		for scanner.Scan() {
			fmt.Printf("\t > %s\n", scanner.Text())
		}
	}()
	go func() {
		scanner := bufio.NewScanner(errReader)
		for scanner.Scan() {
			fmt.Printf("\t > %s\n", scanner.Text())
		}
	}()
	err = cmd.Start()
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error starting Cmd", err)
		return
	}
	go func() {
		for {
			io.WriteString(stdin, "4\n")
		}
	}()
	err = cmd.Wait()
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error waiting for Cmd", err)
		return
	}
}
