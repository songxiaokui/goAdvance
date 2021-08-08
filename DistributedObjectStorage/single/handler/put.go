package handler

import (
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strings"
)

/*
@Time    : 2021/8/8 12:17
@Author  : austsxk
@Email   : austsxk@163.com
@File    : put.go
@Software: GoLand
*/

func put(w http.ResponseWriter, r *http.Request) {
	storageRoot := os.Getenv("STORAGE_ROOT")
	objectName := strings.Split(r.URL.EscapedPath(), "/")[2]
	path := filepath.Join(storageRoot, "objects", objectName)

	f, err := os.Create(path)
	defer f.Close()
	if err != nil {
		log.Fatalln("crate file error: ", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	_, err = io.Copy(f, r.Body)
	if err != nil {
		log.Fatalln("write file error: ", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}
