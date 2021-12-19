package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

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
	result, err := list2comma(os.Stdin)
	if err != nil {
		panic(err)
	}

	fmt.Println(result)
}
