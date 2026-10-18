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

	agent := agents.NewAgent()
	agent.Store = agents.NewStore(nil)

	if err := agents.IndexDirectory(".", agent.Store); err != nil {
		fmt.Println("Index error", err)
	}
	agent.Store.SaveToFile("store.json")

	agent.RegisterTool(types.Tool{
		Name:        "list_files",
		Description: "List Files in directory",
		Execute:     tools.ListFiles,
	})
	agent.RegisterTool(types.Tool{
		Name:        "read_file",
		Description: "Read the full contents of a file",
		Execute:     tools.ReadFile,
	})
	agent.RegisterTool(types.Tool{
		Name:        "write_file",
		Description: "Write content to the file based on suggestion, creating or overwriting it",
		Execute:     tools.WriteFile,
	})
	agent.RegisterTool(types.Tool{
		Name:        "search_files",
		Description: "Search files on certain project without touching the main system directory",
		Execute:     tools.SearchFiles,
	})
	agent.RegisterTool(types.Tool{
		Name:        "copy_file",
		Description: "Copy files to certain directory within project directory without touching the main system directory",
		Execute:     tools.CopyFile,
	})
	agent.RegisterTool(types.Tool{
		Name:        "move_file",
		Description: "Move files to certain directory within project directory without touching the main system directory",
		Execute:     tools.MoveFile,
	})
	agent.RegisterTool(types.Tool{
		Name:        "delete_file",
		Description: "Delete files to certain directory within project directory without touching the main system directory",
		Execute:     tools.DeleteFile,
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

		agent.Run(input)
	}

}
