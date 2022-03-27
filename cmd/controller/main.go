// Ensure build fails on versions of Go that are not supported
// This build tag should be kept in sync with the version specified in go.mod.
//go:build go1.18
// +build go1.18

// Copyright 2022 Authors of welan
// SPDX-License-Identifier: Apache-2.0

package main

import (
	"fmt"
	"github.com/weizhoublue/github-action-test/pkg/lock"
	"github.com/weizhoublue/github-action-test/pkg/print"
	"google.golang.org/grpc"
	"os"
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
	fmt.Println("Testlock 1 ")
	a := &lock.RWMutex{}
	a.RLock()
	a.Lock()
	time.Sleep(10 * time.Second)
	a.Unlock()
}

func main() {
	fmt.Println("hello world")

	val := os.Getenv("GIT_COMMIT_VERSION")
	if len(val) != 0 {
		fmt.Printf("GIT_COMMIT_VERSION=%v\n", val)
	}
	val = os.Getenv("GIT_COMMIT_TIME")
	if len(val) != 0 {
		fmt.Printf("GIT_COMMIT_TIME=%v\n", val)
	}

	_, e := grpc.Dial("localhost:50051")
	if e != nil {
		fmt.Println("failed to  Dial")

	}

	TestRace()
	Testlock()
	Testlock()
	print.MyPrint()
}
