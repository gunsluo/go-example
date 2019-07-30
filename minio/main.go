package main

import (
	"bufio"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"

	"github.com/minio/minio-go"
	"github.com/pkg/errors"
)

const (
	minioURL = "127.0.0.1:9000"
	username = "minio_access"
	password = "minio_secret"
	bucket   = "uploads"

	ThumbnailsKey = "X-Amz-Meta-Thumbnails"
)

func main() {
	client, err := minio.New(minioURL, username, password, false)
	if err != nil {
		panic(err)
	}

	file, err := os.Open("./ldh.jpg")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	fi, err := file.Stat()
	if err != nil {
		// Could not obtain stat, handle error
	}

	r := bufio.NewReader(file)
	size := fi.Size()

	object := "luoji/ldh.jpg"
	opt := minio.PutObjectOptions{}
	ctx := context.Background()

	fileSize, err := client.PutObjectWithContext(ctx, bucket, object, r, size, opt)
	if err != nil {
		panic(err)
	}

	{
		originalMo, existed := ObjectExistsNoCloser(ctx, client, bucket, object)
		if !existed {
			log.Fatal("file not found")
		}
		defer originalMo.Close()

		originalObjInfo, err := originalMo.Stat()
		if err != nil {
			panic(err)
		}
		var thumbnailsInfo []*ThumbnailInfo
		ths := originalObjInfo.Metadata.Get(ThumbnailsKey)
		fmt.Println("-----------", string(ths))
		if ths != "" {
			if err := json.Unmarshal([]byte(ths), &thumbnailsInfo); err != nil {
				panic(err)
			}
		}
		fmt.Println("-----------", string(ths), thumbnailsInfo)
	}

	thumbnails := []*ThumbnailInfo{
		&ThumbnailInfo{
			Width:  100,
			Height: 100,
		},
	}

	b, err := json.Marshal(thumbnails)
	if err != nil {
		panic(err)
		//return errors.Wrapf(err, "unable to marshal new thumbnails info")
	}

	{

		dst, err := minio.NewDestinationInfo(bucket, object, nil, map[string]string{
			ThumbnailsKey: string(b),
		})
		if err != nil {
			panic(err)
		}
		src := minio.NewSourceInfo(bucket, object, nil)

		err = client.CopyObject(dst, src)
		if err != nil {
			panic(err)
		}

		{
			originalMo, existed := ObjectExistsNoCloser(ctx, client, bucket, object)
			if !existed {
				log.Fatal("file not found")
			}
			defer originalMo.Close()

			originalObjInfo, err := originalMo.Stat()
			if err != nil {
				panic(err)
			}
			var thumbnailsInfo []*ThumbnailInfo
			ths := originalObjInfo.Metadata.Get(ThumbnailsKey)
			fmt.Println("-----------", string(ths))
			if ths != "" {
				if err := json.Unmarshal([]byte(ths), &thumbnailsInfo); err != nil {
					panic(err)
				}
			}
			fmt.Println("-----------", string(ths), thumbnailsInfo)
		}
	}

	/*
		{

			file2, err := os.Open("./ldh.jpg")
			if err != nil {
				panic(err)
			}
			defer file2.Close()
			r2 := bufio.NewReader(file2)

			n, err := client.PutObjectWithContext(context.Background(), bucket, object, r2, size, minio.PutObjectOptions{
				//Progress: newProgressReader(originalObjInfo.Size),
				UserMetadata: map[string]string{
					ThumbnailsKey: string(b),
				},
			})
			if err != nil {
				panic(err)
			}
			fmt.Println("---->", n)
		}
	*/
	/*
		if err := originalImageUpdate(client, originalMo, bucket, object, thumbnails...); err != nil {
			panic(err)
		}
	*/

	fmt.Println("---->", fileSize)
}

func ObjectExistsNoCloser(ctx context.Context, client *minio.Client, bucket, object string) (*minio.Object, bool) {
	o, err := client.GetObjectWithContext(ctx, bucket, object, minio.GetObjectOptions{})
	if err != nil {
		return nil, false
	}
	/*
		fi, err := o.Stat()
		if err != nil || fi.Err != nil {
			return nil, false
		}
	*/
	return o, true
}

type ThumbnailInfo struct {
	Width       int    `json:"width, omitempty"`  // thumbnail size
	Height      int    `json:"height, omitempty"` // thumbnail size
	ContentType string `json:"contentType,omitempty"`

	// post filled fields
	URL      string `json:"url,omitempty"`      // download URL for thumbnail
	Bucket   string `json:"bucket,omitempty"`   // bucket that stored the thumbnail
	Object   string `json:"object,omitempty"`   // the object of thumbnail, equal to original file id
	Filename string `json:"filename,omitempty"` // eg. the original file name is: google.png, width is: 10, height is:20, this value will be 10-20-google.png
	Size     int64  `json:"size"`               // thumbnail file size
}

// update user-defined metadata of the original file with new thumbnail info
func originalImageUpdate(client *minio.Client, originalMo *minio.Object, originalBucket, originalObject string, newThumbnails ...*ThumbnailInfo) error {
	originalObjInfo, err := originalMo.Stat()
	if err != nil {
		return errors.Wrapf(err, "unable to stat the original file")
	}
	var thumbnails []*ThumbnailInfo
	ths := originalObjInfo.Metadata.Get(ThumbnailsKey)
	if ths != "" {
		if err := json.Unmarshal([]byte(ths), &thumbnails); err != nil {
			return errors.Wrapf(err, "unable to marshal thumbnail of original file")
		}
	}

	thumbnails = append(thumbnails, newThumbnails...)
	b, err := json.Marshal(thumbnails)
	if err != nil {
		return errors.Wrapf(err, "unable to marshal new thumbnails info")
	}

	originalMo.Seek(0, io.SeekStart)
	fmt.Println("--------->", 1111111111)
	n, err := client.PutObjectWithContext(context.Background(), originalBucket, originalObject, originalMo, originalObjInfo.Size, minio.PutObjectOptions{
		//Progress: newProgressReader(originalObjInfo.Size),
		UserMetadata: map[string]string{
			ThumbnailsKey: string(b),
		},
	})
	fmt.Println("--------->", 22222)
	if err != nil {
		return errors.Wrapf(err, "unable to update original file")
	}
	if n != originalObjInfo.Size {
		return errors.New("unable to write back the exactly byte size into original file")
	}

	return nil
}

/*
func readFile(name string) ([]byte, error) {
	file, err := os.Open(name)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	b, err := ioutil.ReadAll(file)
	if err != nil {
		return nil, err
	}

	return b, nil
}
*/
