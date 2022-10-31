package main

import (
	"log"
	"time"

	"github.com/AlexanderGuan/work"
)

var names = []string{
	"guanliming",
	"huangyuan",
	"cuichengqian",
	"wuhonggang",
	"kangzihao",
}

// namePrinter 使用特定方式打印名字
type namePrinter struct {
	name string
}

// Task 实现worker接口
func (m *namePrinter) Task() {
	log.Println(m.name)
	time.Sleep(time.Second)
}

func main() {
	p := work.New(2)
}
