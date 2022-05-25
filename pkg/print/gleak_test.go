package print_test

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	. "github.com/onsi/gomega/gleak"
	"time"
)

// Please note: gleak is an experimental new Gomega package.
// https://onsi.github.io/gomega/#codegleakcode-finding-leaked-goroutines
// https://github.com/onsi/gomega/tree/master/gleak

/*
 ----------- 下面这种简单的快照写法， 有几种可能，会受到 其它 it 协程的干扰
（1） 在运行 ginkgo  -p 并发时，有并发的其它 it 协程在 干扰。
		解决：可以给 Describe 带上 Serial， 使得 内部所有的 IT 都是 非并发 运行

（2） 即使 ginkgo 不带 -p 并发，有可能出问题：上一个 it1 有 泄露的协程1， 在运行 协程泄露检测的IT2 过程中， 泄露的协程1 完毕退出了
    导致 IT2 的 初始和结束 快照 不一致
		无解：？

*/

var _ = Describe("gleak", Serial, Label("leak"), func() {

	/*
		BeforeEach(func() {
			goods := Goroutines()
			DeferCleanup(func() {
				Eventually(Goroutines).ShouldNot(HaveLeaked(goods))
			})
		})
	*/

	It("test good", func() {
		goods := Goroutines()

		m := Goroutines()
		for t, v := range m {
			GinkgoWriter.Printf("before %v : %+v \n", t, v)
		}

		c := make(chan int)
		go func() {
			time.Sleep(5 * time.Second)
			close(c)
		}()
		<-c

		m = Goroutines()
		for t, v := range m {
			GinkgoWriter.Printf("after %v : %+v \n", t, v)
		}
		// 断言除了 goods外，不会有 泄露的协程
		Eventually(Goroutines).ShouldNot(HaveLeaked(goods))
	})

	It("test bad", func() {
		// Goroutines returns information about all goroutines of a program at this moment.
		// 对当前所有的协程先做个快照，后续 检测协程泄露时，排除他们
		goods := Goroutines()

		go func() {
			time.Sleep(5 * time.Second)
		}()

		// 断言除了 goods外，还发现有 泄露的协程
		Eventually(Goroutines).Should(HaveLeaked(goods))
	})

})

var _ = Describe("gleak2", Serial, Label("leak2"), func() {

	It("test good", func() {
		start := Goroutines()

		m := Goroutines()
		for t, v := range m {
			GinkgoWriter.Printf("before %v : %+v \n", t, v)
		}

		c := make(chan int)
		go func() {
			time.Sleep(5 * time.Second)
			close(c)
		}()
		<-c

		m = Goroutines()
		for t, v := range m {
			GinkgoWriter.Printf("after %v : %+v \n", t, v)
		}

		end := Goroutines()
		// 断言除了 goods外，不会有 泄露的协程
		Expect(start).To(ContainElements(end))
	})

})
