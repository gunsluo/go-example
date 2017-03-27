package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

const (
	underline = '_'
)

func main() {
	readFileToCamel("./test")
}

func readFileToCamel(filename string) {
	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	var strbuf []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		text := strings.TrimSpace(scanner.Text())
		if strings.Index(text, "//") == 0 {
			strbuf = append(strbuf, text)
		} else {
			text = underlineToCamel(text)
			strbuf = append(strbuf, text)
		}
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}

	for _, v := range strbuf {
		fmt.Println(v)
	}
}

func underlineToCamel(param string) string {
	if len(param) == 0 {
		return ""
	}

	buf := []byte(strings.ToLower(param))
	for i := 0; i < len(buf); i++ {
		if buf[i] == underline {
			buf[i+1] -= 32
		}
	}

	return strings.Replace(string(buf), string(underline), "", -1)
}
