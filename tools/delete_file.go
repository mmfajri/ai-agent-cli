package tools

import "os"

func DeleteFile(args map[string]string) string {
	path := args["path"]

	safePath, err := resolveSafePath(path)
	if err != nil {
		return err.Error()
	}

	if err := os.Remove(safePath); err != nil {
		return err.Error()
	}

	return "Deleted " + path
}
