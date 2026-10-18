package tools

import "os"

func MoveFile(args map[string]string) string {
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

	if err := os.Rename(safeSrc, safeDst); err != nil {
		return err.Error()
	}

	return "Moved " + src + " -> " + dst
}
