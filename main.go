package main

import (
	"flag"
)

func main() {
	var countLines, countWords, countBytes bool
	var totalLines, totalWords, totalBytes int64

	flag.BoolVar(&countLines, "l", false, "Count Lines")
	flag.BoolVar(&countWords, "w", false, "Count Words")
	flag.BoolVar(&countBytes, "c", false, "Count Bytes")
	flag.Parse()

	if !countLines && !countWords && !countBytes {
		countLines = true
		countWords = true
		countBytes = true
	}

	if len(flag.Args()) == 0 {
		lines, words, bytes := ProcessStream()
		FormatPrint(countLines, countWords, countBytes, lines, words, bytes, "")
		return
	}

	for _, fileName := range flag.Args() {
		lines, words, bytes := ProcessFile(fileName)
		FormatPrint(countLines, countWords, countBytes, lines, words, bytes, fileName)
		totalLines += lines
		totalWords += words
		totalBytes += bytes
	}

	if len(flag.Args()) > 1 {
		FormatPrint(countLines, countWords, countBytes, totalLines, totalWords, totalBytes, "total")
	}
}
