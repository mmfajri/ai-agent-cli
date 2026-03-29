package agents

func BuildPrompt(input string) string {

	return `
	You are a CLI agent.

	Available tools:
	- list_files(path): list files in a directory

	Rules:
	- If you need a tool, respond ONLY in JSON:
	{"tool": "tool_name", "args":{"key":"value"}}

	- Otherwise respond normally.


	User: ` + input
}
