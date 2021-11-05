package main

import (
	"fmt"
	"sync"
	"time"
)

// 实现goroutine池思路 要有并发原语和任务缓存channel

type Jober interface {
	Do(int)
}

type Scheduling struct {
	wg       sync.WaitGroup
	workPool chan Jober
}

func NewScheduling(cors int) *Scheduling {
	s := Scheduling{
		workPool: make(chan Jober, 100),
	}
	s.wg.Add(cors)
	for i := 0; i < cors; i++ {
		// 每个携程处理任务，从channel中获取
		go func(i int) {
			for j := range s.workPool {
				fmt.Println("current goroutine is ", i)
				j.Do(i)
			}
			// 如果channel关闭，则不阻塞
			s.wg.Done()
		}(i)

	}
	return &s
}

func (s *Scheduling) Add(job Jober) {
	s.workPool <- job
}

func (s *Scheduling) Close() {
	close(s.workPool)
	s.wg.Wait()
}

type J struct {
}

func (j *J) Do(i int) {
	fmt.Println("current ", i)
}

func main() {
	a := NewScheduling(3)
	for {
		j := J{}
		a.Add(&j)
		time.Sleep(time.Second * 4)
	}
}
