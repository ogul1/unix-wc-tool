package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"log"
	"os"
	"syscall"
	"unicode"
)

func ProcessFile(fileName string) (lineCount, wordCount, byteCount int64) {
	file, err := os.Open(fileName)
	if err != nil {
		log.Fatalf("Error opening the file: %v\n", err)
		return
	}
	defer file.Close()

	stat, err := file.Stat()
	if err != nil {
		log.Fatalf("Error obtaining the stats: %v\n", err)
		return
	}

	size := stat.Size()

	data, err := syscall.Mmap(int(file.Fd()), 0, int(size), syscall.PROT_READ, syscall.MAP_SHARED)
	if err != nil {
		log.Fatalf("Error memory mapping: %v\n", err)
		return
	}
	defer syscall.Munmap(data)

	lineCount = int64(bytes.Count(data, []byte("\n")))
	wordCount = int64(len(bytes.Fields(data)))
	byteCount = size

	return
}

func ProcessStream() (lines, words, bytes int64) {
	reader := bufio.NewReader(os.Stdin)
	wordStart := false
	for {
		c, err := reader.ReadByte()
		if err == io.EOF {
			break
		} else if err != nil {
			log.Fatalf("Error reading from stdin: %v\n", err)
			return
		}
		bytes++
		if c == '\n' {
			lines++
		}
		if unicode.IsSpace(rune(c)) && wordStart {
			words++
			wordStart = false
		}
		if !wordStart && !unicode.IsSpace(rune(c)) {
			wordStart = true
		}
	}
	if wordStart {
		words++
	}
	return
}

func FormatPrint(countLines, countWords, countBytes bool, lines, words, bytes int64, fileName string) {
	if countLines {
		fmt.Printf("%8d", lines)
	}
	if countWords {
		fmt.Printf("%8d", words)
	}
	if countBytes {
		fmt.Printf("%8d", bytes)
	}
	if fileName != "" {
		fmt.Printf(" %s", fileName)
	}
	fmt.Printf("\n")
}
