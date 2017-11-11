package main

import (
	"fmt"
	"io/ioutil"
	"net"
	"net/http"
	"regexp"
	"time"

	"github.com/go-errors/errors"
)

func main() {
	client := &http.Client{
		Transport: &http.Transport{
			Dial: dialTimeout,
		},
	}

	for {
		downloadAndSaveTorrentByUrl(client, "luoi", "http://www.jerrylou.me/")
	}
}

func dialTimeout(network, addr string) (net.Conn, error) {
	dial := net.Dialer{
		Timeout:   5 * time.Second,
		KeepAlive: 30 * time.Second,
	}

	conn, err := dial.Dial(network, addr)
	if err != nil {
		return conn, err
	}
	fmt.Println("connect done, use", conn.LocalAddr().String())

	return conn, err
}

func downloadAndSaveTorrentByUrl(client *http.Client, infoHash, url string) error {
	var (
		resp *http.Response
		body []byte
		err  error
		isOk bool
	)

	// retry 3 times
	for failed := 0; failed < 3; failed++ {
		resp, err = client.Get(url)
		if err != nil {
			//logger.Errorf("downloadTorrent: %v", err)
			continue
		}

		if resp.StatusCode != http.StatusOK {
			err = errors.Errorf("%s url:%s %d", infoHash, url, resp.StatusCode)
			continue
		}

		// 读取数据
		body, err = ioutil.ReadAll(resp.Body)
		if err != nil {
			resp.Body.Close()
			break
		}
		resp.Body.Close()

		// 判断数据正确性
		isOk, err = regexp.Match("\\<html[\\S\\s]+?\\</html\\>", body)
		if err != nil {
			err = errors.Errorf("%s body match:%v", infoHash, err)
			break
		}

		if isOk {
			err = errors.Errorf("%s url:%s is html", infoHash, url)
		}

		// torrent内容正确
		break
	}

	if err != nil {
		return err
	}

	return nil
}
