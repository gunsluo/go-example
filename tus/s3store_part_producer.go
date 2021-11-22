package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"time"
)

// s3PartProducer converts a stream of bytes from the reader into a stream of files on disk
type s3PartProducer struct {
	temporaryDirectory string
	files              chan<- *os.File
	done               chan struct{}
	err                error
	r                  io.Reader
}

func (spp *s3PartProducer) produce(totalSize, partSize int64) {
	for {
		if totalSize <= 0 {
			close(spp.files)
			return
		}

		file, err := spp.nextPart(partSize)
		if err != nil {
			spp.err = err
			close(spp.files)
			return
		}
		/*
			if file == nil {
				close(spp.files)
				return
			}
		*/
		select {
		case spp.files <- file:
			totalSize = totalSize - partSize
			fmt.Println("----aaaa", totalSize)
		case <-spp.done:
			close(spp.files)
			return
		}
	}
}

func (spp *s3PartProducer) nextPart(size int64) (*os.File, error) {
	sm := time.Now()
	defer func() {
		em := time.Now()
		fmt.Println("------------->nextPart, d:", size, em.Sub(sm))
	}()
	// Create a temporary file to store the part
	file, err := ioutil.TempFile(spp.temporaryDirectory, "tusd-s3-tmp-")
	if err != nil {
		return nil, err
	}

	limitedReader := io.LimitReader(spp.r, size)
	_, err = io.Copy(file, limitedReader)
	if err != nil {
		return nil, err
	}

	// If the entire request body is read and no more data is available,
	// io.Copy returns 0 since it is unable to read any bytes. In that
	// case, we can close the s3PartProducer.
	//if n == 0 {
	//cleanUpTempFile(file)
	//	return nil, nil
	//}

	// Seek to the beginning of the file
	file.Seek(0, 0)

	return file, nil
}
