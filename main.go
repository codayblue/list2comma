package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var commit string = ""
var version string = "dev"

func printVersion() {
	fmt.Printf("Version: %s\n", version)
	fmt.Printf("Git SHA: %s\n", commit)
}

func list2comma(stdin *os.File) (string, error) {
	scanner := bufio.NewScanner(stdin)
	lines := []string{}

	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		return "", err
	}

	return strings.Join(lines, ","), nil
}

func main() {

	if os.Args[1] == "version" {
		printVersion()
		os.Exit(0)
	}

	result, err := list2comma(os.Stdin)
	if err != nil {
		panic(err)
	}

	fmt.Println(result)
}
