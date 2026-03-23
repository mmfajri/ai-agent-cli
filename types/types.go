package types

type Message struct {
	Role    string
	Content string
}

type ToolCall struct {
	Name string
	Args map[string]string
}

type LLMResponse struct {
	Text     string
	ToolCall *ToolCall
}

type Tool struct {
	Name        string
	Description string
	Execute     func(args map[string]string) string
}
