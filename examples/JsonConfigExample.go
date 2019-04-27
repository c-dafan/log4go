package main

import (
	"github.com/c-dafan/log4go"
	"time"
)

func main() {

	log4go.LoadJsonConfiguration("example.json")
	log4go.Info("hello world")
	log4go.Close()
	time.Sleep(time.Second)
}
