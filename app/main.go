package main

import (
	"fmt"
	"bufio"
	"os"
	"strings"
)

var _ = fmt.Print

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
		arrCmd := strings.Split(command, " ")

		// exit cmd
		if  (len(arrCmd) > 0 && arrCmd[0] == "exit") {
			os.Exit(0)
		}
		fmt.Println(command + ": command not found")
	}
}
