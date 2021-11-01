package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strconv"
	"time"
)

var c = make(chan int)
var downloading bool

// var addr = "https://www.code-aster.org/FICHIERS/salome_meca-2020.0.1-1-universal.tgz"
var addr = "https://www.code-aster.org/FICHIERS/singularity/salome_meca-lgpl-2021.0.0-0-20210601-scibian-9.sif"
var total int64

func main() {
	f, err := os.OpenFile("salome_meca-lgpl-2021.0.0-0-20210601-scibian-9.sif", os.O_RDWR|os.O_APPEND|os.O_CREATE, os.ModePerm)
	if err != nil {
		panic(err)
	}
	resp1, _ := http.Get(addr)
	total = resp1.ContentLength
	resp1.Body.Close()
	go func() {
		var lastSize int64
		for {
			if !downloading {
				time.Sleep(time.Second)
				continue
			}
			<-time.After(time.Second)
			s, _ := f.Stat()
			if total == s.Size() {
				c <- 1
			}
			fmt.Printf("%d kb/s\t%.2f\n", (s.Size()-lastSize)/int64(1024), float64(s.Size())/float64(total)*100)
			lastSize = s.Size()
		}
	}()
	defer f.Close()
EXIT:
	for {

		select {
		case <-c:
			break EXIT
		default:
			stat, _ := f.Stat()
			size := stat.Size()
			req := Download(addr, size)
			fmt.Println(req.Header)
			client := &http.Client{}
			client.Transport = &http.Transport{}
			resp, err := client.Do(req)

			if err != nil {
				downloading = false
				fmt.Println(err)
				time.Sleep(1 * time.Second)
				continue
			}
			fmt.Println(resp.Header)
			downloading = true
			_, err = io.Copy(f, resp.Body)
			if err != nil {
				fmt.Println(err)
				downloading = false
				time.Sleep(1 * time.Second)
				continue
			}
		}

	}
}

func Download(address string, size int64) *http.Request {
	req, _ := http.NewRequest("GET", address, nil)
	req.Close = true
	if size != 0 {
		size += 1
	}
	req.Header.Set("Range", "bytes="+strconv.FormatInt(size, 10)+"-")
	return req
}
