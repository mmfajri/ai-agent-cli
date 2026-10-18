package tools

import "os"

func ReadFile(args map[string]string) string {
	path := args["path"]

	content, err := os.ReadFile(path)
	if err != nil {
		return err.Error()
	}

	return string(content)

}
