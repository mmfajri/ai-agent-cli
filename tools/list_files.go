package tools

import (
	"os"
)

func ListFiles(args map[string]string) string {

	path := args["path"]

	files, err := os.ReadDir(path)
	if err != nil {
		return err.Error()
	}

	var result string
	for _, f := range files {
		result += f.Name() + "\n"
	}

	return result
}
