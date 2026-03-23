package main

import (
	"agent-cli/agents"
	"agent-cli/tools"
	"agent-cli/types"
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	fmt.Println("Welcome To Agent LLM with Deepseek Coder")

	agents := agents.NewAgent()

	agents.RegisterTool(types.Tool{
		Name:        "list_files",
		Description: "List Files in directory",
		Execute:     tools.ListFiles,
	})

	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Print(">> ")
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)

		if input == "exit" || input == "quit" {
			fmt.Println("Quit Program ...")
			break
		}

		agents.Run(input)
	}

}
