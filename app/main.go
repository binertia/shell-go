package main

import (
	"fmt"
	"bufio"
	"os"
	"strings"
)

var _ = fmt.Print


func handleCommand(arrCmd []string, command string) {
	switch arrCmd[0] {
	//echo cmd
	case "echo":
		handleEcho(arrCmd)

	//type cmd
	case "type":
		handleType(arrCmd)

	default:
		fmt.Println(command + ": command not found")
	}
}

func handleType(arrCmd []string) {
	if len(arrCmd) == 1 {
		return
	}

	var totalCmd = []string{"echo","exit","type"}
	for _, s1 := range arrCmd[1:] {
		indicator := 0;
		for _, s2 := range totalCmd {
			if s1 == s2 {
				fmt.Println(s1 + " is a shell builtin")
				indicator++;
			}

		}
		if (indicator == 0) {
			fmt.Println(s1 + ": not found")
		}
	}
}

func handleEcho(arrCmd []string) {
	if len(arrCmd) == 1 {
		fmt.Println("")
	} else {
		arrLen := len(arrCmd)

		for i := 1; i < arrLen; i++ {
			if i == arrLen-1 {
				fmt.Println(arrCmd[i])
			} else {
				fmt.Print(arrCmd[i] + " ")
			}
		}
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
			// exit cmd
			if arrCmd[0] == "exit" {
				break
			}
			handleCommand(arrCmd, command)
		} else {
		fmt.Print("")
		}
	}
}

