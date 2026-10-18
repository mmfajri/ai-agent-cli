package tools

import "os"

func WriteFile(args map[string]string) string {
	path := args["path"]
	content := args["content"]

	err := os.WriteFile(path, []byte(content), 0644)
	if err != nil {
		return err.Error()
	}

	return "File written: " + path

}
