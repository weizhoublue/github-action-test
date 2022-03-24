// Copyright 2016 The Go Authors. All rights reserved.

//go:build go1.7
// +build go1.7

// testdoc implements somthing, this line1 will show up
// this will not be a new line
//
// this line will show up, and will be a new line
package testdoc

import (
	"fmt"
)

const (
	// MaxScanTokenSize is the maximum size used to buffer a token
	// unless the user provides an explicit buffer with Scanner.Buffer.
	MaxScanTokenSize = 64 * 1024

	startBufSize = 4096 // Size of initial allocation for buffer.
)

// Tester implements buffering for an io.Reader object.
type Tester struct {
	// Buf this line will show up
	Buf []int
	R   int // this line will not show up
}

// Show do something ....
//
// this is description for Show
// func must be seen outside , and will be printed by godoc
// the following show how to use
//
// 	func try( )  {
// 		Show()
// 	}
func Show() {
	fmt.Printf("MaxScanTokenSize=%+v startBufSize=%+v \n", MaxScanTokenSize, startBufSize)

	a := Tester{
		Buf: []int{10, 20},
		R:   100,
	}
	fmt.Printf("a=%+v \n", a)
}

// TestBug do something ....
//
// BUG(weizhoublue): this is description for the bug
func TestBug() {
	fmt.Printf("MaxScanTokenSize=%+v startBufSize=%+v \n", MaxScanTokenSize, startBufSize)

}

// TestDeprecated do something ....
//
// Deprecated: this is description
func TestDeprecated() {
	fmt.Printf("MaxScanTokenSize=%+v startBufSize=%+v \n", MaxScanTokenSize, startBufSize)

}
