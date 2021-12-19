package main

import (
	"os"
	"testing"
)

func TestList2Comma(t *testing.T) {
	tempFile, err := os.CreateTemp("", "sample")
	if err != nil {
		t.Fatal("Could not create test input file")
	}
	defer os.Remove(tempFile.Name())

	if _, err := tempFile.Write([]byte("item1\nitem2\nitem3\nitem4")); err != nil {
		t.Fatal("Could not write to test input file")
	}

	if _, err := tempFile.Seek(0, 0); err != nil {
		t.Fatal("Failed to seek back to start of the file")
	}

	result, err := list2comma(tempFile)
	if err != nil {
		t.Fatal(err)
	}

	if result != "item1,item2,item3,item4" {
		t.Fatalf("Input is not as expected: %s", result)
	}

	if err := tempFile.Close(); err != nil {
		t.Fatal("Failed to close test input file")
	}
}
