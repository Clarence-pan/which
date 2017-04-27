package main

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
)

func main() {
	args := os.Args

	if len(args) <= 1 {
		printHelp()
		return
	}

	switch args[1] {
	case "-h":
		fallthrough
	case "-?":
		fallthrough
	case "--help":
		printHelp()
	case "-a":
		fallthrough
	case "--all":
		if len(args) < 3 {
			printHelp()
			return
		}
		printWhichAll(args[2])
	default:
		printWhich1(args[1])
	}
}

func printWhich1(exe string) {
	fullPath, err := exec.LookPath(exe)
	if err != nil {
		panic(err)
	}

	fmt.Println(fullPath)
}

func printHelp() {
	fmt.Println("Usage: which <target>")
	fmt.Println("       - find the one is using now.")
	fmt.Println("Usage: which -a <target>")
	fmt.Println("       - find all in PATH.")
	os.Exit(1)
}

func printWhichAll(exe string) {

	pathList := strings.Split(os.Getenv("PATH"), string(os.PathListSeparator))
	pathList = append([]string{"."}, pathList...)

	pathExtList := strings.Split(os.Getenv("PATHEXT"), string(os.PathListSeparator))
	if os.PathSeparator == '/' {
		pathExtList = []string{""}
	}

	for _, p := range pathList {
		for _, ext := range pathExtList {
			fullPath := strings.TrimRight(p, `/\`) + string(os.PathSeparator) + exe + strings.ToLower(ext)
			if existsFile(fullPath) {
				fmt.Println(fullPath)
			}
		}
	}
}

func existsFile(path string) bool {
	_, err := os.Stat(path)
	if err == nil {
		return true
	}

	if os.IsNotExist(err) {
		return false
	}

	panic(err)
}
