package agents

import (
	"io/fs"
	"os"
	"path/filepath"
	"strings"
)

func IndexDirectory(root string, s *store) error {
	return filepath.Walk(root, func(path string, info fs.FileInfo, err error) error {
		if err != nil || info.IsDir() {
			return err
		}
		if !strings.HasSuffix(path, ".go") {
			return nil
		}

		content, err := os.ReadFile(path)
		if err != nil {
			return err
		}
		chunks := splitIntoChunks(string(content))

		for _, text := range chunks {
			vector, err := Embed(text)
			if err != nil {
				return err
			}
			s.Add(Chunk{Text: text, Vector: vector, Source: path})
		}
		return nil
	})
}

func splitIntoChunks(content string) []string {
	lines := strings.Split(content, "\n")
	var chunks []string
	const size = 40
	for i := 0; i < len(lines); i += size {
		end := i + size
		if end > len(lines) {
			end = len(lines)
		}
		chunks = append(chunks, strings.Join(lines[i:end], "\n"))
	}
	return chunks
}
