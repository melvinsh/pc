package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
	"os/exec"
	"strings"
)

func main() {
	linesFlag := flag.Int("l", 3, "Number of lines to group")
	flag.Parse()

	if len(flag.Args()) < 1 {
		fmt.Fprintln(os.Stderr, "Error: Missing command")
		os.Exit(1)
	}

	cmdName := flag.Args()[0]
	cmdArgs := flag.Args()[1:]

	reader := bufio.NewReader(os.Stdin)

	for {
		group, err := readGroup(reader, *linesFlag)
		if err != nil {
			fmt.Fprintln(os.Stderr, "Error:", err)
			os.Exit(1)
		}

		if len(group) == 0 {
			break
		}

		cmd := exec.Command(cmdName, cmdArgs...)
		stdin, err := cmd.StdinPipe()
		if err != nil {
			fmt.Fprintln(os.Stderr, "Error:", err)
			os.Exit(1)
		}

		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr

		if err := cmd.Start(); err != nil {
			fmt.Fprintln(os.Stderr, "Error:", err)
			os.Exit(1)
		}

		for _, line := range group {
			fmt.Fprintln(stdin, line)
		}
		stdin.Close()

		if err := cmd.Wait(); err != nil {
			fmt.Fprintln(os.Stderr, "Error:", err)
			os.Exit(1)
		}

		if err == io.EOF {
			break
		}
	}
}

func readGroup(reader *bufio.Reader, lines int) ([]string, error) {
	var group []string

	for i := 0; i < lines; i++ {
		line, err := reader.ReadString('\n')
		if err != nil && err != io.EOF {
			return group, err
		}
		line = strings.TrimSuffix(line, "\n")

		if len(line) > 0 {
			group = append(group, line)
		}

		if err == io.EOF {
			break
		}
	}

	return group, nil
}
