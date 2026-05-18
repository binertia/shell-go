package main

import (
	"fmt"
	"bufio"
	"os"
)

// Ensures gofmt doesn't remove the "fmt" import in stage 1 (feel free to remove this!)
var _ = fmt.Print

func main() {
	command, err := bufio.NewReader(os.Stdin).ReadString("\n")
	fmt.Println(command[:len(command)-1] + ": command not found")
}
