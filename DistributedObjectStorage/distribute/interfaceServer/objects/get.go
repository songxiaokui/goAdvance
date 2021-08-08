package objects

import (
	"distribute/interfaceServer/locate"
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"
)

/*
@Time    : 2021/8/8 20:04
@Author  : austsxk
@Email   : austsxk@163.com
@File    : get.go
@Software: GoLand
*/

func get(w http.ResponseWriter, r *http.Request) {
	object := strings.Split(r.URL.EscapedPath(), "/")[2]
	stream, e := getStream(object)
	if e != nil {
		log.Println(e)
		w.WriteHeader(http.StatusNotFound)
		return
	}
	io.Copy(w, stream)
}

func getStream(object string) (io.Reader, error) {
	server := locate.Locate(object)
	if server == "" {
		return nil, fmt.Errorf("object %s locate fail", object)
	}
	return NewGetStream(server, object)
}
