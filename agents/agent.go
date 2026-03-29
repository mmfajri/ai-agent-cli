package agents

import (
	"agent-cli/llms"
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

	prompt := BuildPrompt(input)

	for i := 0; i < 5; i++ {
		llmText, err := llms.CallDeepSeek(prompt)
		if err != nil {
			fmt.Println("LLM error: ", err)
			return
		}

		toolCall := ParseToolCall(llmText)

		if toolCall == nil {
			fmt.Println("Agent: ", llmText)
			return
		}

		tool, exists := a.Tools[toolCall.Name]
		if !exists {
			fmt.Println("Tool not found: ", toolCall.Name)
			return
		}

		result := tool.Execute(toolCall.Args)

		fmt.Println("Tool Result: ", result)

		prompt = prompt + "\nTool result:\n" + result
	}

}
