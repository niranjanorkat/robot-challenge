package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/niranjanorkat/robot-challenge/repl"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println(repl.MsgWelcome)

	for {
		fmt.Print(repl.MsgPrompt)
		input, _ := reader.ReadString('\n')
		parts := strings.Fields(strings.TrimSpace(input))
		if len(parts) == 0 {
			continue
		}
		if repl.HandleCommand(parts) {
			break
		}
	}
}
