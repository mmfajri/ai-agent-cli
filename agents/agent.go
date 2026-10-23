package agents

import (
	"agent-cli/llms"
	"agent-cli/types"
	"fmt"
)

type Agent struct {
	Tools   map[string]types.Tool
	History []types.Message
	Store   *store
}

func NewAgent() *Agent {
	return &Agent{
		Tools:   make(map[string]types.Tool),
		History: []types.Message{},
	}
}

func (a *Agent) RegisterTool(tool types.Tool) {
	a.Tools[tool.Name] = tool
}

func (a *Agent) SaveToMemoryTool() types.Tool {
	return types.Tool{
		Name:        "save_to_memory",
		Description: "Save a piece of learned information for future recall",
		Execute: func(args map[string]string) string {
			text := args["content"]
			vector, err := Embed(text)
			if err != nil {
				return err.Error()
			}
			a.Store.Add(Chunk{Text: text, Vector: vector, Source: "memory"})
			if err := a.Store.SaveToFile("store.json"); err != nil {
				return err.Error()
			}
			return "Saved to memory"
		},
	}
}

func (a *Agent) Run(input string) {
	a.History = append(a.History, types.Message{Role: "user", Content: input})

	vector, err := Embed(input)
	if err != nil {
		fmt.Println("Embed error: ", err)
		return
	}
	context := a.Store.Search(vector, 3)
	prompt := BuildPrompt(input, context, a.History)

	for i := 0; i < 5; i++ {
		llmText, err := llms.CallDeepSeek(prompt)
		if err != nil {
			fmt.Println("LLM error: ", err)
			return
		}

		toolCall := ParseToolCall(llmText)

		if toolCall == nil {
			fmt.Println("Agent: ", llmText)
			a.History = append(a.History, types.Message{Role: "assistant", Content: llmText})
			return
		}

		tool, exists := a.Tools[toolCall.Name]
		if !exists {
			fmt.Println("Tool not found: ", toolCall.Name)
			prompt = prompt + "\nError: tool '" + toolCall.Name + "' does not exist. Use one of the available tools listed above.\n"
			continue
		}

		if toolCall.Name == "delete_file" {
			fmt.Printf("Confirm delete of %s? (y/n): ", toolCall.Args["path"])
			var resp string
			fmt.Scanln(&resp)
			if resp != "y" {
				fmt.Println("Delete  cancelled")
				return
			}
		}

		result := tool.Execute(toolCall.Args)

		fmt.Println("Tool Result: ", result)

		prompt = prompt + "\nTool result:\n" + result
	}

}
