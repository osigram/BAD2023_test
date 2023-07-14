package main

import (
	"fmt"
	"io"
	"os"
	"regexp"
	"strings"
)

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		panic("File opening error")
	}
	textBytes, err := io.ReadAll(file)
	if err != nil {
		panic("File reading error")
	}
	text := string(textBytes)

	re := regexp.MustCompile("\\w+")
	words := re.FindAllString(text, -1)
	found := make([]string, 0, len(words))

	for _, word := range words {
		wordSlice := strings.Split(word, "")

		if symbol := getSymbol(wordSlice); symbol != "" {
			found = append(found, getSymbol(wordSlice))
		}
		continue
	}

	fmt.Println(getSymbol(found))
}

func getSymbol(wordSlice []string) string {
	var result string

symbolOut:
	for num, symbol := range wordSlice {
		for _, symbol2 := range wordSlice[num+1:] {
			if symbol == symbol2 {
				continue symbolOut
			}
		}

		result = symbol
		break
	}

	return result
}
