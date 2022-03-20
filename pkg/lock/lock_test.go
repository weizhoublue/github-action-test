//  a package ending in *_test is allowed to live in the same directory as the package being tested
package lock_test

import (
	"fmt"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"github.com/weizhoublue/github-action-test/pkg/lock"
	"os"
	"path/filepath"
	"time"
)

// ================== 基本使用 ============

// Describe, Context, and When 三者是完全相同的 alias
// 对于 Describe, Context, It, When and Measure ，如果前面加上可前缀 P or an X，那么则意味这些 case是 Pending 的，那么，在运行整个测试时，这些 case 是不会被运行的
// Describe主要可用于归类 It ，我们可在Describe中定义多个Describe
var _ = Describe("Lock", func() {

	// 定义全局变量，以供后续 BeforeEach 和 It 函数间共享
	// 此处，变量只声明 ，不初始化
	var keyA, keyB string

	// 在每个It.测试用例运行前，都会运行一下本函数
	// Describe 中能够使用任意数量的 BeforeEach, AfterEach, JustBeforeEach, It, and Measurement blocks
	BeforeEach(func() {
		keyA = "init"
		Expect(true).To(BeTrue())
		fmt.Printf("BeforeEach : I am just run before each It \n")

	})
	// JustBeforeEach(...)
	JustAfterEach(func() {
		By("每个It 运行之后 ，先运行JustAfterEaches  ，后运行 AfterEach")
		By("This can be useful if you need to collect diagnostic information")

		// https://onsi.github.io/ginkgo/#attaching-data-to-reports
		if CurrentSpecReport().Failed() {
			// AddReportEntry(name string, args ...interface{})
			AddReportEntry("data-dump", []int{1, 2, 3})
		}

	})
	AfterEach(func() {
		By("clean up code here")
	})

	// 依据测试场景，可用 Context来归类 It
	Context("With more than 300 pages", func() {

		// 针对 本 Context 中的 it 的配置初始化，其他的 Context 可以有自己的 BeforeEach
		// 如果嵌套定义了多层的JustAfterEaches 和 AfterEach, from the outer-most to the inner-most
		BeforeEach(func() {
			keyB = "context_1_wrong_value_test"

			err := os.Setenv("WEIGHT_UNITS", "smoots")
			Expect(err).NotTo(HaveOccurred())

			// 我们可以使用 DeferCleanup  来替代 AfterEach ， 这样，在恢复初始值 时，代码 更加简洁
			//  You can also pass a function that returns a single value. DeferCleanup interprets this value as an error and fails the spec if the error is non-nil
			// 如果 DeferCleanup 和 AfterEach 都存在，那么 AfterEach 先于  DeferCleanup 运行
			originalWeightUnits := os.Getenv("WEIGHT_UNITS")
			DeferCleanup(func() {
				err := os.Setenv("WEIGHT_UNITS", originalWeightUnits)
				Expect(err).NotTo(HaveOccurred())
			})
		})

		// ---------- 日志输出
		It("test output", func() {
			//  ginkgo -v  或者  ginkgo  --always-emit-ginkgo-writer  就会显示如下打印
			// https://onsi.github.io/ginkgo/#logging-output
			GinkgoWriter.Println("run test label filter , ")
			GinkgoWriter.Printf("byby , keyA=%v , keyB=%v \n", keyA, keyB)

			// https://onsi.github.io/ginkgo/#attaching-data-to-reports
			// AddReportEntry generates and adds a new ReportEntry to the current spec's SpecReport
			// 能追加到 测试报告的 本specs中 ， ginkgo --json-report ./a.json
			// 我们也可以控制打印格式 If the value implements fmt.Stringer or types.ColorableStringer then value.String() or value.ColorableString() (which takes precedence) is used to generate the representation, otherwise Ginkgo uses fmt.Sprintf("%#v", value)
			AddReportEntry("Report_testwelan", []string{"anyDataStruct"})

			// 当 测试失败  或者 使用 (ginkgo -v 或者  ginkgo  --always-emit-ginkgo-writer )  时，这些信息 才会被打印出来
			// The string passed to By is emitted via the GinkgoWriter. If a test succeeds you won't see any output beyond Ginkgo's green dot. If a test fails, however, you will see each step printed out up to the step immediately preceding the failure. Running with ginkgo -v always emits all steps.
			// 这种输出 也会 写入到 测试报告中 By also adds a ReportEntry to the running spec
			By("Running sleep command which is longer than timeout for context")
			By("report the runtime of callback function ", func() {
				time.Sleep(1 * time.Second)
			})

		})

		// https://onsi.github.io/ginkgo/#mental-model-how-ginkgo-handles-failure
		It("test fail", func() {
			if false {
				// 宣告测试失败，提前终止
				// fail 的实现机制是 调用了 panics
				Fail("Gomega generates a failure message and passes it to Ginkgo to signal that the spec has failed")
			}
			fmt.Println("you could not get  here ")

			t := &lock.Mutex{}
			t.Lock()
			t.Unlock()
		})
	})

})

// ==================基于 table 的写法=======================

// https://onsi.github.io/ginkgo/#table-specs
var _ = Describe("Lock_table", func() {

	var glo string
	BeforeEach(func() {
		glo = "i am initialized"
	})

	// https://pkg.go.dev/github.com/onsi/ginkgo/extensions/table
	// `DescribeTable` simply generates a new Ginkgo `Describe`. Each `Entry` is turned into an `It` within the `Describe`
	// func DescribeTable(description string, itBody interface{}, entries ...TableEntry) bool
	DescribeTable("the > inequality",
		// 以下是每个用例都会调用的 函数
		func(x int, y int, expected bool) {
			// .... do something

			Expect(x > y).To(Equal(expected))
		},
		// 以下是每个 用例
		// func Entry(description interface{}, parameters ...interface{}) TableEntry
		// 每个entry 都会 运行一个 It ，每个entry 都会调用 func(x int, y int, expected bool)
		// 每个 entry 的第一个参数是一个描述，后续的所有参数 都是传递给func(x int, y int, expected bool)
		// Individual Entries can be focused (with FEntry) or marked pending (with PEntry or XEntry)
		Entry("x > y", 1, 0, true),
		Entry("x == y", 0, 0, false),
		Entry("x < y", 0, 1, false),
	)

	// 在 entry 中，不要引用Describe中的全局变量，因为 在ginko解析翻译DescribeTable阶段，全局变量还未初始化，读区不到值
	DescribeTable("failed to get global variable value",
		func(a string) {
			By(fmt.Sprintf("global var value:%s.", a))
			Expect(a).NotTo(Equal("i am initialized"))
		},
		// 错误的用法
		Entry("bad", glo),
	)

	// 指定 entry 的 description
	DescribeTable("set entry description",
		func(x int, y int, expected bool) {
			Expect(x > y).To(Equal(expected))
		},
		EntryDescription("customized description: %v %v %v"),
		Entry("x > y", 1, 0, true),
	)

})

// ================= 有序运行测试用例 ========================

// https://onsi.github.io/ginkgo/#ordered-containers
// 使用 Ordered 修饰符，能够使得 所有的测试用例 都是  按照顺序来执行的
var _ = Describe("Lock_Ordered", Ordered, func() {
	var a string

	// 不同于 BeforeEach（会为每个用例分别运行一次） ， BeforeAll 只会在 运行所有用例前 只运行一次
	// BeforeAll 必须在 an Ordered container 中使用
	BeforeAll(func() {
		a = "1"
		fmt.Printf("BeforeAll : a=%v \n", a)

		// 我们可以使用 DeferCleanup  来替代 AfterAll ， 这样，在恢复初始值 时，代码 更加简洁
		// 如果 DeferCleanup 和 AfterAll 都存在，那么 AfterAll 先于  DeferCleanup 运行
		DeferCleanup(func() {
			fmt.Printf("AfterAll : a1=%v \n", a)
		})
	})

	It("step 1", func() {
		fmt.Printf("step1: a=%v \n", a)
		a = "2"

		// Fail("在order的 用例中，调用 fail 会 停止运行 后续的用例")
	})

	It("step 2", func() {
		fmt.Printf("step2: a=%v \n", a)
		a = "3"
	})

	AfterAll(func() {
		fmt.Printf("AfterAll : a2=%v \n", a)
	})
})

var _ = Describe("Lock_Ordered2", func() {
	var a, b string

	// 对于 Context 为 Ordered  的外部，如果定义了一个BeforeEach ， 那么会出问题的
	// 默认，BeforeEach 会为 所有Ordered 测试完成一次初始化，那么  Ordered 测试之间的  数据依赖会被 BeforeEach 破坏
	// 所有，为了避免该问题，需要为 BeforeEach 添加OncePerOrdered 修饰符，这样，BeforeEach 会为 Ordered block 只跑一次
	BeforeEach(OncePerOrdered, func() {
		b = "1"
	})

	Context("With more than 300 pages", Ordered, func() {

		BeforeAll(func() {
			a = "1"
			fmt.Printf("BeforeAll : a=%v \n", a)
			fmt.Printf("BeforeAll : b=%v \n", b)
		})

		It("step 1", func() {
			fmt.Printf("step1: a=%v \n", a)
			a = "2"
			fmt.Printf("step1: b=%v \n", b)
			b = "2"
		})

		It("step 2", func() {
			fmt.Printf("step2: a=%v \n", a)
			a = "2"
			fmt.Printf("step2: b=%v \n", b)
			b = "2"
		})

	})
})

// ================= 测试 过滤 ========================

var _ = Describe("Lock_filter", func() {

	// https://pkg.go.dev/github.com/onsi/ginkgo/v2#Label
	// 可以给 各种 block 添加 Label(...) ， 其中可以添加多个 字符串
	// Labels can container arbitrary strings but cannot contain any of the characters in the set: "&|!,()/"
	// 可添加多个， 运行  ginkgo run -label-filter='integrater_test && unit_test' ， 就实现 只运行这些 用例
	It("test label filter", Label("integrater_test", "unit_test"), func() {
		GinkgoWriter.Printf("byby")
	})
	// 这种写法 效果 一样
	It("test label filter", Label("label1"), Label("label2"), func() {
		GinkgoWriter.Printf("byby")
	})

	// https://onsi.github.io/ginkgo/#focused-specs
	// 若某个 测试被 打上 focus ，那么 ginkgo 运行时，该 测试 suit下，只会运行 focus 的测试，这样，主要用于 debug 时 使用，不跑其他 不相关的测试
	// 运行ginkgo unfocus 会忽略 是否有Focus之分
	// It("test label focus", Focus, func() {
	//	GinkgoWriter.Printf("focus test")
	// })

	// https://onsi.github.io/ginkgo/#pending-specs
	// 若某个 测试被 打上 Pending ，那么 ginkgo 运行时 是 不会 跑这些 测试
	It("test pending", Pending, func() {
		GinkgoWriter.Printf("pending test")
	})

})

// ================= 测试 超时 ========================

var _ = Describe("test timeout", func() {

	// ginkgo --timeout=3s  指定 一个用例的 最长运行时间，防止其 卡住或出问题了，如果超时，则失败。 如果不指定，默认是 1小时
	// ginkgo --slow-spec-threshold=20s  若一个用例运行时间超过该值，则会在测试报告中 体现 该用例 跑得久 “[SLOW TEST]...”。默认是 5s
	It("test time", func() {
		GinkgoWriter.Printf("go to sleep ")
		time.Sleep(10 * time.Second)
	})

})

// ============= 协程使用 ==============
var _ = Describe("test goroutine", Label("routine"), func() {

	It("test goroutine panic", func() {
		done := make(chan interface{})
		go func() {
			// 使用本函数，能让 ginkgo 捕获任意在 协程中的 断言失败、 panic 和 fail()
			defer GinkgoRecover()
			Expect(1).To(BeEquivalentTo(1))
			if false {
				Fail("boom")
			}
			close(done)
		}()
		Eventually(done).Should(BeClosed())
	})

	// 协程的运行时间，可能会影响 断言失败
	It("test goroutine timeout", func() {
		done := make(chan interface{})
		go func() {
			defer GinkgoRecover()
			time.Sleep(2 * time.Second)
			close(done)
		}()
		// https://pkg.go.dev/github.com/onsi/gomega@v1.18.1#Eventually
		// Eventually 默认值会 等待 done 为 1s ， 所以，如果协程中 运行过久，会使得 断言失败
		// Eventually(done).Should(BeClosed())
		// 可以使用如下方式来 指定 等待接口 的 超时时间
		Eventually(done, "10s").Should(BeClosed())
	})

})

// ================= 测试 并发的 竞争问题 ========================

var _ = Describe("test race", Label("race"), func() {

	// 规避 文件读写的 并发
	// 为每个 并发的 进程生成一个 独立的 临时目录, 供它们 各自 用来 读写、生成 文件
	// 而同一个进程中的 用例是顺序执行的，且每次用例执行完成后，DeferCleanup 也会清除临时目录里的数据，这样也保证了 同一个进程中的用例之间不会相互干扰
	// 其它 类似的竞争问题，都可以 借鉴 GinkgoParallelProcess() 和 DeferCleanup 来解决
	Describe("resolve file race", func() {
		var pathTo func(path string) string

		BeforeEach(func() {
			// shard based on our current process index.
			// this starts at 1 and goes up to N, the number of parallel processes.
			dir := fmt.Sprintf("./tmp-%d", GinkgoParallelProcess())
			os.MkdirAll(dir, os.ModePerm)

			// 配合 ginkgo --fail-fast ，若测试用例失败时，能够保留 临时目录里的数据，供debug，否则删除
			DeferCleanup(func() {
				suiteConfig, _ := GinkgoConfiguration()
				if CurrentSpecReport().Failed() && suiteConfig.FailFast {
					GinkgoWriter.Printf("Preserving artifacts in %s\n", dir)
					return
				}
				Expect(os.RemoveAll(dir)).To(Succeed())
			})

			// 传递给每个用例
			pathTo = func(path string) string { return filepath.Join(dir, path) }
		})

		Context("Publishing books", func() {
			It("can publish a complete epub", func() {
				filepath := pathTo("ourData.filename")
				fmt.Printf(" we can use file path %v to save data\n", filepath)
			})

			It("can publish a preview that contains just the first chapter", func() {
				filepath := pathTo("ourData.filename")
				fmt.Printf(" we can use file path %v to save data\n", filepath)
			})
		})
	})

})
