package main

import (
	"encoding/json"
	"fmt"
	"os"
	"os/exec"
	"strconv"
	"strings"
)

type DuDir struct {
	Path string `json:"path"`
	Size int64  `json:"size"`
}

func main() {
	// Note: you can use a constant array instead of command line arguments here:
	// err := doMain([]string{"/tmp", "/etc"})
	err := doMain(os.Args[1:])
	if err != nil {
		_, _ = fmt.Fprintf(os.Stderr, err.Error())
		os.Exit(1)
	}
}

func doMain(dirNames []string) error {
	if len(dirNames) == 0 {
		return fmt.Errorf("usage: %s <path> [<path>...]", os.Args[0])
	}

	result := make([]DuDir, 0)
	for _, d := range dirNames {
		st, err := os.Stat(d)
		if os.IsNotExist(err) {
			return fmt.Errorf("%s does not exist", d)
		}
		if err != nil {
			return fmt.Errorf("cannot stat %s: %w", d, err)
		}
		if !st.IsDir() {
			return fmt.Errorf("cannot stat %s: not a directory", d)
		}
		cmd := exec.Command("du", "-b", "-d0", d)
		out, err := cmd.CombinedOutput()
		parts := strings.Split(string(out), "\t")
		if len(parts) < 2 {
			return fmt.Errorf("du -b -d0 %v: %v %w", d, string(out), err)
		}
		size, err := strconv.ParseInt(string(parts[0]), 10, 64)
		if err != nil {
			return fmt.Errorf("du -b -d0 %v: cannot parse %v: %w", d, parts[0], err)
		}
		result = append(result, DuDir{d, size})
	}
	output, err := json.Marshal(result)
	if err != nil {
		return err
	}
	fmt.Println(string(output))
	return nil
}
