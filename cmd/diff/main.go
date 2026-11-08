package main

import (
	"os"
	"strings"

	"github.com/ericoericochen/godiff"
)

func getLines(s string) []string {
	text := strings.ReplaceAll(s, "\r\n", "\n")
	lines := strings.Split(text, "\n")
	return lines
}

func main() {
	srcFilePath, destFilePath := os.Args[1], os.Args[2]
	srcFile, err := os.ReadFile(srcFilePath)
	if err != nil {
		panic(err)
	}
	destFile, err := os.ReadFile(destFilePath)
	if err != nil {
		panic(err)
	}
	srcTxt := string(srcFile)
	destTxt := string(destFile)

	srcLines := getLines(srcTxt)
	destLines := getLines(destTxt)

	diffs := godiff.GetDiff(srcLines, destLines)
	// fmt.Println(diffs)

	godiff.Display(srcLines, destLines, diffs)
}
