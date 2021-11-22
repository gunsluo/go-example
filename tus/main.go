package main

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"os"
	"time"
)

func main() {
	var totalSize int64

	sm := time.Now()
	data, err := os.ReadFile("/Users/luoji/gopath/src/tespkg.in/meerastorage/docker/demo/128m.data")
	if err != nil {
		log.Fatal(err)
	}

	//data := []byte("1234567890123")
	totalSize = int64(len(data))
	reader := bytes.NewReader(data)
	src := io.LimitReader(reader, totalSize)
	//src := io.NewSectionReader(src1, 0, totalSize)
	em := time.Now()
	fmt.Println("------------->read file, d:", totalSize, em.Sub(sm))

	/*
		sm := time.Now()
		f, err := os.Open("/Users/luoji/gopath/src/tespkg.in/meerastorage/docker/demo/128m.data")
		if err != nil {
			log.Fatal(err)
		}
		fs, _ := f.Stat()
		totalSize = fs.Size()
		src := io.LimitReader(f, totalSize)
		em := time.Now()
		fmt.Println("------------->read file, d:", em.Sub(sm))
	*/

	//var src io.Reader

	var optimalPartSize int64 = 1024 * 1024 * 64
	//var optimalPartSize int64 = 4

	fileChan := make(chan *os.File, 100)
	doneChan := make(chan struct{})
	defer close(doneChan)

	partProducer := s3PartProducer{
		temporaryDirectory: "./data",
		done:               doneChan,
		files:              fileChan,
		r:                  src,
	}
	go partProducer.produce(totalSize, optimalPartSize)

	for file := range fileChan {
		stat, err := file.Stat()
		if err != nil {
			log.Fatal(err)
		}
		n := stat.Size()

		fmt.Println("->", n, file.Name())
	}
}

func cleanUpTempFile(file *os.File) {
	file.Close()
	os.Remove(file.Name())
}
