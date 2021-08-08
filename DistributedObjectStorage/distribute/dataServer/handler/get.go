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
@File    : get.go
@Software: GoLand
*/

func get(w http.ResponseWriter, r *http.Request) {
	storageRoot := os.Getenv("STORAGE_ROOT")
	objectName := strings.Split(r.URL.EscapedPath(), "/")[2]
	path := filepath.Join(storageRoot, "objects", objectName)

	f, err := os.Open(path)
	defer f.Close()
	if err != nil {
		log.Panicln("open file error: ", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	io.Copy(w, f)
}
