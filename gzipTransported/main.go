package main

/*
@Time    : 2021/12/2 10:15
@Author  : austsxk
@Email   : austsxk@163.com
@File    : main.go
@Software: GoLand
*/

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"net/http"
	"regexp"
	"strings"
)

func CodeStyle() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Header("Content-Encoding", "gzip")
		c.Next()
	}
}

func IsExistFilePath(fileName string) bool {
	sourceUri := "http://127.0.0.1:8888/file/"
	resp, err := http.Get(sourceUri)
	if err != nil {
		fmt.Println(err)
		return false
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	fmt.Println(string(body))
	reString := `<a[\s\S]+?href="([\s\S]+?)">([\s\S]+?)</a>`
	re := regexp.MustCompile(reString)
	data := re.FindAllStringSubmatch(string(body), -1)

	dataMap := make(map[string]bool, 0)
	for _, value := range data {
		if len(value[1]) == 0 {
			continue
		}
		dataMap[strings.ReplaceAll(value[1], " ", "")] = true
	}
	fmt.Println("dataMap；", dataMap)

	if _, ok := dataMap[fileName]; !ok {
		return false
	}
	return true
}

type reqQuery struct {
	FileName string `json:"file_name" form:"file_name"`
}

func GetList(c *gin.Context) {
	var f reqQuery
	err := c.ShouldBindQuery(&f)
	fmt.Println(err)
	if IsExistFilePath(f.FileName) {
		c.JSON(http.StatusOK, gin.H{"message": "ok"})
		return
	}
	c.JSON(http.StatusNotFound, gin.H{"message": "failed"})
}
func main() {
	s := gin.Default()
	s.StaticFS("/file", http.Dir("./static"))
	s.GET("/query/", GetList)
	if err := s.Run(":8888"); err != nil {
		panic("server failed!")
	}

}
