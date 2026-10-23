package tools

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

const projectRoot = "."

func resolveSafePath(userPath string) (string, error) {
	joined := filepath.Join(projectRoot, userPath)
	absRoot, err := filepath.Abs(projectRoot)
	if err != nil {
		return "", err
	}

	absJoined, err := filepath.Abs(joined)
	if err != nil {
		return "", err
	}

	if !strings.HasPrefix(absJoined, absRoot+string(os.PathSeparator)) && absJoined != absRoot {
		return "", fmt.Errorf("access denied: path escapes project root")
	}
	return absJoined, nil
}

func SearchFiles(args map[string]string) string {
	pattern := args["pattern"]
	root := args["path"]

	if root == "" {
		root = "."
	}

	safeRoot, err := resolveSafePath(root)
	if err != nil {
		return err.Error()
	}

	var matches []string
	filepath.Walk(safeRoot, func(path string, info os.FileInfo, err error) error {
		if err != nil || info.IsDir() {
			return nil
		}
		content, err := os.ReadFile(path)
		if err == nil && strings.Contains(string(content), pattern) {
			matches = append(matches, path)
		}
		return nil
	})
	return strings.Join(matches, "\n")
}
