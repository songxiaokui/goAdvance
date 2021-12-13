package main

import (
	"errors"
	"fmt"
	"os"
	"strings"
)

/*
@Time    : 2021/11/15 15:42
@Author  : austsxk
@Email   : austsxk@163.com
@File    : parser.go
@Software: GoLand
*/

// 日志解析
/*
salome类型的操作日志:
	最后一行以: EXIT_CODE=0 为成功

aster类型的日志:
	最后一行: EXECUTION_CODE_ASTER_EXIT_%s=0 为成功

saturne类型日志:
	===========================================================




						  END OF CALCULATION
						  ==================


	===============================================================
	以 END OF CALCULATION 为求解计算标准
*/

// 工厂模式
// LogParser 日志解析接口
type LogParser interface {
	Parser(path string) (bool, error)
}

// PathExists 判断路径是否存在
func PathExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}

// ReadLogContainsEndMark 读取文件结尾多少字节判断是否包含结束标示
// path 文件路径
// mark 结束标示关键词
// threshold 读取字节阈值
func ReadLogContainsEndMark2(path, mark string, threshold int64) bool {
	file, err := os.Open(path)
	defer file.Close()
	if err != nil {
		return false
	}
	buf := make([]byte, threshold)
	stat, err := os.Stat(path)
	if err != nil {
		return false
	}
	start := stat.Size() - threshold
	if start < 0 {
		start = 0
	}
	_, err = file.ReadAt(buf, start)
	lines := strings.Split(string(buf), "\n")
	for _, line := range lines {

		if len(line) == 0 {
			continue
		}
		//s := strings.TrimSpace(line)
		//
		//fmt.Println("长度:", len(s) )
		//fmt.Println("值:", s )
		//fmt.Printf("%T", s )
		fmt.Println(mark)
		if strings.Contains(line, mark) {
			return true
		}
	}
	return false
}

// salome 日志解析
type SalomeLogParser struct {
}

func (this *SalomeLogParser) Parser(path string) (bool, error) {
	ok, err := PathExists(path)
	if !ok || err != nil {
		return ok, err
	}
	data := ReadLogContainsEndMark2(path, "EXIT_CODE=0", 400)
	if !data {
		return false, errors.New("calculate is failed.\n")
	}
	return data, nil
}

// aster日志解析
type AsterLogParser struct {
}

func (this *AsterLogParser) Parser(path string) (bool, error) {
	ok, err := PathExists(path)
	if !ok || err != nil {
		return ok, err
	}
	data := ReadLogContainsEndMark2(path, "=0", 80)
	if !data {
		return false, errors.New("calculate is failed.\n")
	}
	return data, nil
}

// saturne日志解析
type SaturneLogParser struct {
}

func (this *SaturneLogParser) Parser(path string) (bool, error) {
	ok, err := PathExists(path)
	if !ok || err != nil {
		return ok, err
	}
	data := ReadLogContainsEndMark2(path, "END", 2000)

	if !data {
		return false, errors.New("calculate is failed.\n")
	}
	return data, nil
}

// NewLogParser 根据计算资源类型创建一个日志解析对象
func NewLogParser(resourceType int) LogParser {
	switch resourceType {
	case 1:
		return &SalomeLogParser{}
	case 2:
		return &AsterLogParser{}
	case 3:
		return &SalomeLogParser{}
	default:
		return nil
	}
}

func main() {
	a := NewLogParser(3)
	fmt.Println(a.Parser("run_solver.log"))
}
