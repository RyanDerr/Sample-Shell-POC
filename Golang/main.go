package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"os/exec"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	for {
		showWorkingDir()
		//Read Input From Keyboard
		input, err := reader.ReadString('\n')
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
		}

		if err = executeCommand(input); err != nil {
			fmt.Fprintln(os.Stdout, err)
		}
	}
}

func showWorkingDir() {
	//Get the users current directory
	workingDir, err := os.Getwd()
	//Print error to console if and error is returned
	if err != nil {
		log.Println(err)
	}
	//Print a '>' to the console before entering commands
	fmt.Print(workingDir + "> ")
}

func executeCommand(commandInput string) error {
	commandInput = strings.TrimSuffix(commandInput, "\n")

	args := strings.Split(commandInput, " ")

	cmd := exec.Command(args[0], args[1:]...)

	cmd.Stderr = os.Stderr
	cmd.Stdout = os.Stdout

	return cmd.Run()
}
