package main

import (
	"fmt"
	"bufio"
	"os"
	"strings"
	"os/exec"
)

var _ = fmt.Print

func handleExecBin(arrCmd []string, command string) {
	// ::TODO:: will use this line below after finish for absolute path
	// path, err := exec.LookPath(arrCmd[0])
	_, err := exec.LookPath(arrCmd[0])
	if err != nil {
		fmt.Println(command + ": command not found")
	} else {
		// ::TODO:: will use this line after finish project
		// cmd := exec.Command(path,arrCmd[1:]...)
		cmd := exec.Command(arrCmd[0],arrCmd[1:]...)

		cmd.Stdin = os.Stdin
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		err = cmd.Run()
		if err != nil {
			fmt.Println(err)
		}
	}
}


func handleCommand(arrCmd []string, command string) {
	switch arrCmd[0] {
	//echo cmd
	case "echo":
		handleEcho(arrCmd)

	//type cmd
	case "type":
		handleType(arrCmd)
	
	//pwd cmd
	case "pwd":
		handlePwd(arrCmd)

	default:
		handleExecBin(arrCmd, command)
	}
}

func handlePwd(arrCmd []string) {
	if len(arrCmd) > 1 {
		fmt.Println("pwd: too many arguments")
	}
	dir, err := os.Getwd()
	if err != nil {
		fmt.Println(err);
	} else {
		fmt.Println(dir);
	}
}

func handleType(arrCmd []string) {
	if len(arrCmd) == 1 {
		return
	}

	var builtinCmd = []string{"echo","exit","type"}
	for _, s1 := range arrCmd[1:] {
		indicator := 0;
		for _, s2 := range builtinCmd {
			if s1 == s2 {
				fmt.Println(s1 + " is a shell builtin")
				indicator++;
			}

		}
		if (indicator == 0) {
			path, err := exec.LookPath(s1)
			if err != nil {
				fmt.Println(s1 + ": not found")
			} else {
				fmt.Println(s1 + " is " + path)
			}
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

