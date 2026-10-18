package agents

func BuildPrompt(input string, context []Chunk) string {
	var ctxText string
	for _, c := range context {
		ctxText += "// from " + c.Source + "\n" + c.Text + "\n\n"
	}

	return `
    You are a CLI agent.

    Relevant code context:
    ` + ctxText + `

    Available tools:
    - list_files(path): list files in a directory
    - read_file(path): read the full contents of a file
    - write_file(path, content): write text inside the file based on suggestion
    - search_files(path): search the files within directory project, exclude the main system directory
    - copy_file(source, destination): copy a file within the project
    - move_file(source, destination): move or rename a file within the project
    - delete_file(path): permanently delete a file within the project

    Rules:
    - If you need a tool, respond ONLY in JSON:
    {"tool": "tool_name", "args":{"key":"value"}}

    Examples:
    {"tool":"read_file","args":{"path":"main.go"}}
    {"tool":"write_file","args":{"path":"notes.txt","content":"hello"}}
    {"tool":"copy_file","args":{"source":"a.txt","destination":"b.txt"}}

    - Otherwise respond normally.

    - Otherwise respond normally.


    User: ` + input
}
