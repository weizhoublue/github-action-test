// Copyright 2022 Authors of welan
// SPDX-License-Identifier: Apache-2.0

package testdoc_test

import (
	"fmt"
	"github.com/weizhoublue/github-action-test/pkg/testdoc"
)

// this will show in godoc as a example of function usage
// Example${functionName}
func ExampleTestBug() {
	fmt.Printf("generate godoc showing how to call function TestBug() \n")
	testdoc.TestBug()

}

// this will show in godoc as a example of datatype
// Example${dataType}
func ExampleTester() {
	fmt.Printf("generate godoc showing how to use dataType \n")
	t := testdoc.Tester{
		Buf: []int{1, 2, 3},
		R:   100,
	}
	t.Method()
}

// this will show in godoc as a example of datatype method
// Example${dataType}_{Methodname}
func ExampleTester_Method() {
	fmt.Printf("generate godoc showing how to call method of sturct \n")
	t := testdoc.Tester{
		Buf: []int{1, 2, 3},
		R:   100,
	}
	t.Method()
}
