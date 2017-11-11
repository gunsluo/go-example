package main

import (
	"fmt"
	"os"
	"time"
)

func main() {
	filename := "test.txt"

	fd, err := os.OpenFile(filename, os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0644)
	if err != nil {
		panic(err)
	}

	for {
		_, err := fd.WriteString(fmt.Sprintf("%s\n", "luoji"))
		if err != nil {
			fmt.Println("error: %v", err)
			break
		}
		fd.Sync()
		time.Sleep(100 * time.Millisecond)
	}

	fd.Close()
}
