package main

import (
	"fmt"
	"bufio"
	"os"
	"strings"
)

var _ = fmt.Print

func handleCommand(arrCmd []string, command string) {
	// exit cmd
	if arrCmd[0] == "echo" {

		if len(arrCmd) == 1 {
			fmt.Println("")
		}

		arrLen := len(arrCmd)

		for i := 1; i < arrLen; i++ {
			if i == arrLen-1 {
				fmt.Println(arrCmd[i])
			} else {
				fmt.Print(arrCmd[i] + " ")
			}
		}

	} else {
		fmt.Println(command + ": command not found")
	}
}

func main() {
	readLine := bufio.NewReader(os.Stdin)
	for {
		fmt.Print("$ ")
		command, err := readLine.ReadString('\n')
		if err != nil {
			fmt.Fprintln(os.Stderr, "Error reading input:", err)
			os.Exit(1)
		} 

		command = strings.TrimSpace(command)
		arrCmd := strings.Fields(command)

		if len(arrCmd) > 0 {
			if arrCmd[0] == "exit" {
				break
			}
			handleCommand(arrCmd, command)
		} else {
		fmt.Print("")
		}
	}
}

