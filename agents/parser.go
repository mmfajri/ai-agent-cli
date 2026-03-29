package agents

import (
	"agent-cli/types"
	"encoding/json"
)

func ParseToolCall(text string) *types.ToolCall {
	var result map[string]any

	err := json.Unmarshal([]byte(text), &result)
	if err != nil {
		return nil
	}

	toolName, ok := result["tool"].(string)
	if !ok {
		return nil
	}

	argsRaw, ok := result["args"].(map[string]any)
	if !ok {
		return nil
	}

	args := make(map[string]string)
	for k, v := range argsRaw {
		args[k] = v.(string)
	}

	return &types.ToolCall{
		Name: toolName,
		Args: args,
	}
}
