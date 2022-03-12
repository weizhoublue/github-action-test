package main

import (
	"fmt"
	"github.com/weizhoublue/github-action-test/pkg/lock"
	"github.com/weizhoublue/github-action-test/pkg/print"
	"time"
)

func TestRace() {
	a := 10

	fmt.Println("TestRace ")
	go func() {
		a++
	}()
	go func() {
		a++
	}()
	time.Sleep(2 * time.Second)

}

func Testlock() {
	fmt.Println("Testlock ")
	a := &lock.RWMutex{}
	a.RLock()
	a.Lock()
	time.Sleep(10 * time.Second)
	a.Unlock()
	return
}

func main() {
	fmt.Println("hello world")
	TestRace()
	Testlock()
	print.MyPrint()
}
