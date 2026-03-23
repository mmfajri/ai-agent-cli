package agents

import "agent-cli/types"

func FakeLLM(input string) types.LLMResponse {
	if input == "list files" {
		return types.LLMResponse{
			ToolCall: &types.ToolCall{
				Name: "list_files",
				Args: map[string]string{
					"path": ".",
				},
			},
		}
	}
	return types.LLMResponse{
		Text: "I don't know what to do yet.",
	}
}
