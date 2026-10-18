package tools

import (
	"io"
	"os"
)

func CopyFile(args map[string]string) string {
	src := args["source"]
	dst := args["destination"]

	safeSrc, err := resolveSafePath(src)
	if err != nil {
		return err.Error()
	}
	safeDst, err := resolveSafePath(dst)
	if err != nil {
		return err.Error()
	}

	in, err := os.Open(safeSrc)
	if err != nil {
		return err.Error()
	}

	out, err := os.Create(safeDst)
	if err != nil {
		return err.Error()
	}
	defer out.Close()

	if _, err := io.Copy(out, in); err != nil {
		return err.Error()
	}

	return "Copied " + src + " -> " + dst
}
