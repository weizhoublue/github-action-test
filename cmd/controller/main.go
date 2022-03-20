// Ensure build fails on versions of Go that are not supported
// This build tag should be kept in sync with the version specified in go.mod.
//go:build go1.17
// +build go1.17

package main

import (
	"fmt"
	"github.com/weizhoublue/github-action-test/pkg/lock"
	"github.com/weizhoublue/github-action-test/pkg/print"
	"google.golang.org/grpc"
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
}

func main() {
	fmt.Println("hello world")

	grpc.Dial("localhost:50051")

	TestRace()
	Testlock()
	Testlock()
	print.MyPrint()
}
