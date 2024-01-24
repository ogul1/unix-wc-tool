package main

import (
	"testing"
)

func TestProcessFile(t *testing.T) {
	lines, words, bytes := ProcessFile("data/test.txt")

	if !(lines == 7144 && words == 58164 && bytes == 334998) {
		t.Errorf("Got %8d %8d %8d, wanted %8d %8d %8d", lines, words, bytes, 7144, 58164, 334998)
	}
}

func TestProcessFile2(t *testing.T) {
	lines, words, bytes := ProcessFile("data/example.txt")

	if !(lines == 108861 && words == 136078 && bytes == 505446) {
		t.Errorf("Got %8d %8d %8d, wanted %8d %8d %8d", lines, words, bytes, 108861, 136078, 505446)
	}
}
