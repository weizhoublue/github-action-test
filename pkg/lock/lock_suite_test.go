// Copyright 2022 Authors of welan
// SPDX-License-Identifier: Apache-2.0

//  a package ending in *_test is allowed to live in the same directory as the package being tested
package lock_test

import (
	"flag"
	"testing"

	"fmt"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	
)

func TestLock(t *testing.T) {
	// 在使用 gomega断言前，需要使用如下函数注册  断言失败时的回调函数
	RegisterFailHandler(Fail)
	RunSpecs(t, "Lock Suite")
}

// -------------------

var myFlag string

func init() {
	// 我们可以使用 这种方式 获取 ginkgo 运行命令传入的参数    ginkgo -- myFlag=something
	flag.StringVar(&myFlag, "myFlag", "defaultvalue", "myFlag is used to control my behavior")
}

// -------------------

// 在 ginkgo -p  并发运行 测试时，每个ginkgo并发的每一个进程  都会 运行 ，

var _ = BeforeSuite(func() {
	fmt.Println("BeforeSuite：we can initial environment,  for all test before hand here")

	// 当在BeforeSuite 中 setup 依赖服务时，每个并发进程都运行 一份服务，
	// 优点是能规避 ginkgo 并发测试 时 的 测试用例之间  数据竞争问题 ，避免相互失败
	// 缺点是 消耗资源，且要注意规避 网络监听端口的 竞争问题
	// 注意的是，每个 测试用例 可在 BeforeSuite 中 对 外部服务进行 初始化数据恢复，避免 用例之间 相互 干扰
	libraryAddr := fmt.Sprintf("127.0.0.1:%d", 50000+GinkgoParallelProcess())
	fmt.Printf(" we could setup data server here with different port %v for different Process\n", libraryAddr)
	
})

var _ = AfterSuite(func() {
	fmt.Println("AfterSuite：we can close environment,  for all test after hand here")
})

// -------------------SynchronizedBeforeSuite 适合用于setup 某个服务 , 使用一个 独立进程来 运行
// ----------------- 缺点：  ginkgo并发多个进程  来跑 测试用例时，不同进程上的 测试  可能会 竞争 同一服务数据，导致相互 失败

// SynchronizedBeforeSuite 和 BeforeSuite 只能 二选一 来使用
/*
 https://onsi.github.io/ginkgo/#parallel-suite-setup-and-cleanup-synchronizedbeforesuite-and-synchronizedaftersuite
 https://pkg.go.dev/github.com/onsi/ginkgo/v2#SynchronizedBeforeSuite
 https://onsi.github.io/ginkgo/#managing-external-processes-in-parallel-suites
	在 ginkgo -p 并发运行所有用例时，有多个进程来完成。
	（1）一些 环境搭建的 工作，如果放在 BeforeSuite 中，那么所有的 进程都会 运行一遍，这样，可能 会浪费资源
	since BeforeSuite runs on every parallel process this would result in N independent databases spinning up.
	Sometimes that's exactly what you want - as it provides maximal isolation for the running specs and is a natural way to shard data access. Sometimes, however, spinning up multiple external processes is too resource intensive or slow and it is more efficient to share access to a single resource
	（2）可以使用 SynchronizedBeforeSuite 和 SynchronizedAfterSuite ，它 会在一个进程中完成 初始化工作，然后把 信息传递给 其他所有 运行用例的进程
	 allows us to set up state in one process, and pass information to all the other processes ,
	This is useful for performing expensive or singleton setup once, then passing information from that setup to all parallel processes
*/

// 例如如下，在 进程1 运行 数据共享服务，把 其服务地址 传递给 其他 运行测试用例的进程
// var _ = SynchronizedBeforeSuite(
//	//process1Body，建立服务
//	func() []byte {
//		//runs *only* on process #1
//		fmt.Printf(" setup database server here")
//		// 把服务地址传递给 所有进程
//		return []byte("127.0.0.1:6443")
//	},
//	//allProcessBody ， 接收服务的一些信息
//	func(address []byte) {
//		// 所有集成中，可以创建一个client 接入server
//		fmt.Printf(" create a client connected to %v , the client can be used by all test case", string(address))
//		DeferCleanup(func() {
//			fmt.Printf(" close the client ")
//		})
//	},
// )

// var _ = SynchronizedAfterSuite(
//	func() {
//		//runs on *all* processes
//		Expect(dbClient.Cleanup()).To(Succeed())
//	},
//	func() {
//		//runs *only* on process #1
//		Expect(dbRunner.Stop()).To(Succeed())
//	},
// )
