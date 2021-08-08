package objects

import (
	"distribute/interfaceServer/heartbeat"
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
@File    : put.go
@Software: GoLand
*/

func put(w http.ResponseWriter, r *http.Request) {
	object := strings.Split(r.URL.EscapedPath(), "/")[2]
	c, e := storeObject(r.Body, object)
	if e != nil {
		log.Println(e)
	}
	w.WriteHeader(c)
}

func storeObject(r io.Reader, object string) (int, error) {
	stream, e := putStream(object)
	if e != nil {
		return http.StatusServiceUnavailable, e
	}

	io.Copy(stream, r)
	e = stream.Close()
	if e != nil {
		return http.StatusInternalServerError, e
	}
	return http.StatusOK, nil
}

func putStream(object string) (*PutStream, error) {
	server := heartbeat.RandomSelectDataServersPoint()
	if server == "" {
		return nil, fmt.Errorf("cannot find any dataServer")
	}

	return NewPutStream(server, object), nil
}
