package agents

import (
	"agent-cli/types"
	"fmt"
)

type Agent struct {
	Tools   map[string]types.Tool
	History []types.Message
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

func (a *Agent) Run(input string) {
	a.History = append(a.History, types.Message{
		Role:    "user",
		Content: input,
	})

	resp := FakeLLM(input)

	if resp.ToolCall != nil {
		tool, exists := a.Tools[resp.ToolCall.Name]

		if !exists {
			fmt.Println("Error: tool not found: ", resp.ToolCall.Name)
			return
		}

		if tool.Execute == nil {
			fmt.Println("Error: tool has no execution function: ", resp.ToolCall.Name)
			return
		}

		result := tool.Execute(resp.ToolCall.Args)

		fmt.Println("Tool Result:")
		fmt.Println(result)

		a.History = append(a.History, types.Message{
			Role:    "tool",
			Content: result,
		})
		return
	}

	fmt.Println("Agent: ", resp.Text)
}
