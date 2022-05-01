// Copyright 2022 Authors of welan
// SPDX-License-Identifier: Apache-2.0

package lock_test

import (
	"context"
	"fmt"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"github.com/onsi/gomega/gbytes"
	"github.com/onsi/gomega/gexec"
	"github.com/onsi/gomega/gmeasure"
	"os"
	"os/exec"
	"time"
)

// ================================

// https://pkg.go.dev/github.com/onsi/gomega@v1.18.1

var _ = Describe("gomega_assert", func() {

	It("t1", func() {

		_, cancel := context.WithTimeout(context.Background(), 3*time.Second)
		defer cancel()

		// https://pkg.go.dev/github.com/onsi/gomega
		// 类型和值 的相等比较
		// with reflect.DeepEqual , Equal() is strict about types and value comparisons
		Expect(1).To(Equal(1))
		// 尝试类型转换后 进行值 相等比较
		// wiith reflect.DeepEqual , BeEquivalentTo 在使用 reflect.DeepEqual 比较前，会尝试 先进行双方的类型转化
		Expect("123").To(BeEquivalentTo("123"))

		// 数字比较  请使用 BeNumerically()
		// https://pkg.go.dev/github.com/onsi/gomega@v1.18.1#BeNumerically
		// BeNumerically 会尝试转化类型后，进行数值比较
		Expect(0.1).To(BeNumerically("<=", 0.2))
		Expect(0.1).To(BeNumerically("==", 0.1))
		// 如下 比较符号 ，判断 Actual 是否 在 expected 的 指定 偏差范围内
		Expect(1).To(BeNumerically("~", 0.9, 0.1))

		// 时间比较
		// actual 是否 在 预期时间的 浮动范围 内
		Expect(time.Duration(10) * time.Second).To(BeNumerically("~", 8*time.Second, 3*time.Second))

		// 字符串的比较
		Expect("i am substring").To(Equal("i am substring"))
		Expect("i am substring").To(ContainSubstring("am"))

		// bool 比较
		Expect(true).To(BeTrue())
		Expect(false).To(BeFalse())

		// error 的比较
		err := os.Setenv("WEIGHT_UNITS", "smoots")
		// BeNil() 比较 ACTUAL 是否是 nil
		Expect(err).To(BeNil())
		// 没有错误发生
		Expect(err).To(Succeed())
		Expect(err).NotTo(HaveOccurred())
		// BeZero() 比较 ACTUAL 是否是 nil 或者 数据类型的0值（ 未初始化）
		Expect(err).To(BeZero())

		// 只返回 error 的函数 的 成功调用
		// Succeed passes if actual is a nil error
		// Succeed is intended to be used with functions that return a single error value
		Expect(os.Setenv("WEIGHT_UNITS", "smoots")).To(Succeed())

		// 空值的 数据类型 实例的断言
		// BeEmpty succeeds if actual is empty. Actual must be of type string, array, map, chan, or slice.
		// Expect(dbClient.Books()).To(BeEmpty())

		// 对 array, slice or map 的比较 , 其中的 成员 必须要全部出现在 ConsistOf 中
		// ConsistOf https://pkg.go.dev/github.com/onsi/gomega@v1.18.1#ConsistOf
		// 注意，ConsistOf 中的值的 顺序 不重要
		// By default ConsistOf() uses Equal() to match the elements
		Expect([]string{"Foo", "FooBar"}).Should(ConsistOf("FooBar", "Foo"))
		// 如下 是失败
		// Expect([]string{"Foo", "FooBar"}).Should(ConsistOf("FooBar"))
		// 字符串子集
		Expect([]string{"Foo", "FooBarTest"}).Should(ConsistOf(ContainSubstring("Bar"), "Foo"))
		//  对于 map的比较， 比的是 values， 而忽略 mao 的 key
		Expect(map[string]string{"1": "Foo", "2": "FooBarTest"}).Should(ConsistOf(ContainSubstring("Bar"), "Foo"))

		// 对 array, slice or map 的比较 , ContainElements中的成员 出现 即可， 即子集即可
		// ContainElements() uses Equal() to match the elements
		// 注意，ContainElements 中的值的 顺序 不重要
		//  对于 map的比较， 比的是 values， 而忽略 mao 的 key
		Expect([]string{"Foo1", "Foo2Bar", "cat"}).Should(ContainElements("Foo1", ContainSubstring("Bar")))

		// 测试文件存在
		Expect("les-miserables.epub").NotTo(BeAnExistingFile())
	})

	//
	Context("test_channel", func() {

		It("test channel date and close", func() {
			c := make(chan int, 10)
			var result int

			go func() {
				defer GinkgoRecover()
				time.Sleep(2 * time.Second)
				c <- 10
				time.Sleep(2 * time.Second)
				c <- 20
				time.Sleep(2 * time.Second)
				close(c)
			}()

			// https://pkg.go.dev/github.com/onsi/gomega@v1.18.1#Eventually
			// https://pkg.go.dev/github.com/onsi/gomega@v1.18.1#Receive

			// Eventually 默认值会 等待 done 为 1s ， 所以，如果协程中 运行过久，会使得 断言失败, 可以使用如下方式来 指定 等待接口 的 超时时间
			// 使用 Receive() 来接收数据，一次获取一个数据
			// Actual must be a channel (and cannot be a send-only channel) -- anything else is an error.
			// 如果channel 被关闭了，Receive 会失败
			Eventually(c, "10s").Should(Receive(&result))
			fmt.Printf("get result1: %v \n", result)

			// 用这种方式 更加简洁，直接判断 读取的值 是否 符合预期
			Eventually(c, "10s").Should(Receive(BeNumerically("==", 20)))

			// https://pkg.go.dev/github.com/onsi/gomega@v1.18.1#Consistently
			// 在100ms内，持续 不断尝试 读取channel， 希望期间内 接收不到数据
			Consistently(c, "100ms").ShouldNot(Receive())

			// 确认关闭
			Eventually(c, "10s").Should(BeClosed())

		})
	})

	// https://pkg.go.dev/github.com/onsi/gomega#Eventually
	
	// https://pkg.go.dev/github.com/onsi/gomega@v1.18.1/gexec
	Context("test_call_bin", func() {

		// gexec 和 gbytes packages 配合，执行外部bin，并断言其 返回 内容
		It("call1", func() {
			cmd := exec.Command("/bin/bash", "-c", `sleep 2 && echo "this line1"  && echo "that line2" `)

			// gexec.Start 会以异步方式 来启动该命令
			session, err := gexec.Start(cmd, GinkgoWriter, GinkgoWriter)
			Expect(err).To(Succeed())

			// 如果 cmd 会话 在 3s 内未执行完毕，会失败
			// //In addition to teeing to GinkgoWriter gexec will capture any stdout/stderr output to gbytes buffers
			// 如果cmd 返回的 STD_OUT/STD_ERR 不符合预期，也会失败
			// gbytes.Say 只要输出STD_OUT/STD_ERR中 有这个数据，即成功，这个匹配 是 正则式的
			Eventually(session, "3s").Should(gbytes.Say(`line1`))
			// 如果上一句 gbytes.Say 已经把 line2 匹配了，那么 这句 line2 会匹配不中
			Eventually(session, "3s").Should(gbytes.Say(`that line2`))

			// We can also assert the session has exited
			Eventually(session).Should(gexec.Exit(0)) // with exit code 0
		})
	})

	Context("test_call_FUNC", func() {

		// 校验协程的运行时间
		It("method1", func() {
			done := make(chan interface{})
			go func() {
				defer GinkgoRecover()
				defer close(done)
				time.Sleep(2 * time.Second)
			}()
			// 指定时间内，希望读取成功 channel ，且结果符合预期
			Eventually(done, "3s").Should(BeClosed())

		})

		// 函数要在指定 超时时间内完成 , 且 校验了 各个返回值
		// https://onsi.github.io/gomega/#category-3-making-assertions-eminem-the-function-passed-into-codeeventuallycode
		It("method2", Label("shanghai") , func() {
			var resultInt int

			// 指定时间内，希望完成 函数调用，且结果符合预期
			p := func(g Gomega) {
				defer GinkgoRecover()
				// resultInt , err:=func(....) 在 此 调用 其他 被测试函数
				var err error
				// 函数无错误发生
				g.Expect(err).To(Succeed())
				// 模拟返回值 ， 对其校验
				resultInt = 1
				g.Expect(resultInt).To(BeNumerically("<=", 17))
			}
			// https://onsi.github.io/gomega/#category-2-making-codeeventuallycode-assertions-on-functions
			// 指定 函数 运行的超时时间
			// Eventually supports accepting functions that take a single Gomega argument and return zero or more values
			Eventually(p, "5s").Should(Succeed())

			// 获取 被调用函数的 返回数据，供后续使用
			fmt.Printf("method2 get result of caller func: %v \n", resultInt)

			//或者简写为
			// 在 指定时间(10s)内 ，会尝试一直 间隔1s  不断 反复 调用 函数， 只要有一次 满足期待(包括满足函数内的各种断言语句)，就提前结束成功 ！！！！！！！！！！！！！！！！！
			Eventually(func(g Gomega) bool {
				defer GinkgoRecover()
				var err error = nil
				GinkgoWriter.Println("Eventually  1 ")
				time.Sleep(time.Second)
				// do something , or call another funciton

				//函数内的所有断言也需要满足成功，才能说本次成功
				// 函数内请使用 传入的 g 的断言，而不是 Expect(err).To(Succeed())
				g.Expect(err).NotTo(HaveOccurred())
				return true
			},"10s", "1s").Should(BeTrue())
			GinkgoWriter.Println("finish Eventually ")


			// 另一种 超时写法
			Eventually(func(g Gomega) bool {
				defer GinkgoRecover()
				// ...
				return true
			}).WithTimeout(10*time.Second).WithPolling(1 * time.Second).Should(BeTrue())
			
			// 在指定时间内，会尝试一直 不断 反复 调用 函数， 每次都 希望 是 符合预期的结果(包括满足函数内的各种断言语句)，期间 只要有一次不对，就失败 ！！！！！！！！！！！！！！！！！
			Consistently(func(g Gomega) bool {
				defer GinkgoRecover()
				var err error = nil
				GinkgoWriter.Println("Consistently  1 ")
				time.Sleep(time.Second)
				// do something , or call another funciton

				//函数内的所有断言也需要满足成功，才能说本次成功
				// 函数内请使用 传入的 g 的断言，而不是 Expect(err).To(Succeed())
				g.Expect(err).To(Succeed())
				return true
			}, "10s", "1s" ).Should(BeTrue())
			GinkgoWriter.Println("finish Consistently ")

		})

	})

})

// =============== 性能测试 gomega/gmeasure =================
// https://onsi.github.io/ginkgo/#benchmarking-code
// https://onsi.github.io/gomega/#codegmeasurecode-benchmarking-code

/*
gmeasure 能对被调函数的  反复 执行时间 或 返回值 进行 统计

输出的性能数据 样例：
      Name                  | N | Min     | Median  | Mean    | StdDev | Max
      ==========================================================================
      time-test1 [duration] | 3 | 1.0005s | 1.0011s | 1.0009s | 300µs  | 1.0011s
      --------------------------------------------------------------------------
      time-test2 [duration] | 3 | 100µs   | 1.0012s | 1.0008s | 817ms  | 2.0012s
      --------------------------------------------------------------------------
      value-test1           | 3 | 0.000   | 1.000   | 1.000   | 0.816  | 2.000

      名字 / 运行的重复次数 / 花费的最短时间 / 花费的平均时间 / * / * / 花费的最长时间
      名字 / 运行的重复次数 / 返回组小值 / 返回平均值 / * / * / 返回最大值
*/

// 反复调用 测试函数
var _ = Describe("repeated call func for performance test", func() {

	// 方式1：（使用起来 不方便） 使用 experiment.MeasureDuration 或者 experiment.MeasureValue 来调用 被测量业务
	// 性能测试，建议都打上 Serial， 不与其它测试 并行
	It("method by experiment.MeasureXXX", Serial, Label("measurement"), func() {

		// we create a new experiment
		// Experiments are thread-safe
		// https://pkg.go.dev/github.com/onsi/gomega@v1.18.1/gmeasure#NewExperiment
		experiment := gmeasure.NewExperiment("performance test sample")
		// 把性能数据 追加到 测试报告中
		AddReportEntry(experiment.Name, experiment)

		// 设置 性能测试 参数
		// https://pkg.go.dev/github.com/onsi/gomega@v1.18.1/gmeasure#SamplingConfig
		// 重复运行 20次，且 每次最多运行 10s
		sampleConf := gmeasure.SamplingConfig{N: 3, MinSamplingInterval: 10 * time.Second}
		// 重复运行 20次，且 总计运行时间 不能超过 1min (若时间提前超时，则直接结束)
		// sampleConf := gmeasure.SamplingConfig{N: 3, Duration: 1 * time.Second}
		// 每次循环 都是 串行执行的，可设置 NumParallel 实现 并发数量
		// sampleConf := gmeasure.SamplingConfig{N: 3, Duration: 20 * time.Second, NumParallel: 3}

		// ------- 开始 测量 性能 ---------
		// https://pkg.go.dev/github.com/onsi/gomega@v1.18.1/gmeasure#Experiment.Sample
		// pass callback func(idx int)
		// 以下是 被重复 运行的 被测试代码 , idx 是 第几次运行，其序号从 0 开始
		experiment.Sample(func(idx int) {
			fmt.Printf("I am sample at %v time , initialize env for each sample here\n", idx)

			// MeasureDuration测试  是 测试  被调函数的 执行时间
			// https://pkg.go.dev/github.com/onsi/gomega@v1.18.1/gmeasure#Experiment.MeasureDuration
			// 传入回调格式： func()
			experiment.MeasureDuration("time-test1", func() {
				// 在此 书写 被性能测试的代码
				fmt.Printf("performance test1 code here\n")
				for i := 0; i <= 1; i++ {
					t := i * i
					if t < 0 {
						break
					}
				}
				time.Sleep(30 * time.Millisecond)
			})

			// An experiment can record multiple Measurement
			// 可以追加一些 修饰符
			experiment.MeasureDuration("time-test2", func() {
				// 在此 书写 被性能测试的代码
				fmt.Printf("performance test2 code here\n")
				// idx 是 第几次运行，其序号从 0 开始
				time.Sleep(time.Duration((idx+1)*30) * time.Millisecond)
			}, gmeasure.Annotation("this is annotation for time test2"))

			// MeasureValue 测试 返回值
			// https://pkg.go.dev/github.com/onsi/gomega@v1.18.1/gmeasure#Experiment.MeasureValue
			// 传入回调格式： func() float64
			experiment.MeasureValue("value-test1", func() float64 {
				// 在此 书写 被性能测试的代码
				fmt.Printf("here, code who return float64 \n")
				return float64(idx)
			})

		}, sampleConf)

		// ------ 时间数据处理
		// 可以 获取 性能数据
		// https://pkg.go.dev/github.com/onsi/gomega@v1.18.1/gmeasure#Stats
		stats := experiment.GetStats("time-test2")
		// 校验是否在规定时间内跑完了 规定的测试次数
		runCount := stats.N
		Expect(runCount).To(BeNumerically("==", 3))
		// 判断消耗时间，是否在 预期值的 浮动范围内
		// https://pkg.go.dev/github.com/onsi/gomega@v1.18.1/gmeasure#Stat
		median := stats.DurationFor(gmeasure.StatMedian)
		Expect(median).To(BeNumerically("~", 60*time.Millisecond, 10*time.Millisecond))
		max := stats.DurationFor(gmeasure.StatMax)
		Expect(max).To(BeNumerically("<=", 4*time.Second))

		// ------ 数值数据处理
		pstats := experiment.GetStats("value-test1")
		vm := pstats.ValueFor(gmeasure.StatMedian)
		Expect(vm).To(BeNumerically("<=", 10))

		// ------ 可以比较 多组数据
		// https://pkg.go.dev/github.com/onsi/gomega@v1.18.1/gmeasure#RankStats
		ranking := gmeasure.RankStats(gmeasure.LowerMedianIsBetter,
			experiment.GetStats("time-test1"),
			experiment.GetStats("time-test2"))
		AddReportEntry("Ranking", ranking)
		// assert the winner
		Expect(ranking.Winner().MeasurementName).To(Equal("time-test1"))

	})

	// 方式2：使用 Stopwatch 测量 时间, 书写 方便 , 还能 暂停计时 ， 推荐
	It("method by Stopwatch ", Serial, Label("measurement"), func() {
		experiment := gmeasure.NewExperiment("performance test sample")
		AddReportEntry(experiment.Name, experiment)

		// https://pkg.go.dev/github.com/onsi/gomega@v1.18.1/gmeasure#SamplingConfig
		sampleConf := gmeasure.SamplingConfig{N: 3, MinSamplingInterval: 10 * time.Second}

		// ------- 测量 性能 ---------
		// https://pkg.go.dev/github.com/onsi/gomega@v1.18.1/gmeasure#Experiment.Sample
		experiment.Sample(func(idx int) {
			defer GinkgoRecover() // necessary since NewStopwatch will launch as goroutines and contain assertions
			// 开始启动一个后端协程 计时
			stopwatch := experiment.NewStopwatch() // we make a new stopwatch for each sample.  Experiments are threadsafe, but Stopwatches are not.

			fmt.Printf("do round %v : business 1: step1 here\n", idx)
			time.Sleep(100 * time.Millisecond)
			// 对已经消耗的时间 生成一个记录
			stopwatch.Record("bussiness1").Reset()

			fmt.Printf("do round %v : business 2 we not care\n", idx)
			time.Sleep(200 * time.Millisecond)

			// reset to record new one
			stopwatch.Reset() // Subsequent recorded durations will measure the time elapsed from the moment Reset was called
			fmt.Printf("do round %v : business 3: step1 here\n", idx)
			time.Sleep(200 * time.Millisecond)

			// 暂停一些 我们不在乎的步骤
			stopwatch.Pause()
			fmt.Printf("do round %v : business 3: step2 we not care\n", idx)
			time.Sleep(50 * time.Millisecond)
			stopwatch.Resume()

			fmt.Printf("do round %v : business 3: step3 here\n", idx)
			time.Sleep(100 * time.Millisecond)
			// 可以追加若干 修饰符
			stopwatch.Record("bussiness2", gmeasure.Annotation("this is b2 annotation")).Reset()

		}, sampleConf)

		// ------ 性能数据处理
		// 可以 获取 性能数据
		// https://pkg.go.dev/github.com/onsi/gomega@v1.18.1/gmeasure#Stats
		stats := experiment.GetStats("bussiness1")
		fmt.Printf("bussiness1 stats: %+v \n", stats)
		stats = experiment.GetStats("bussiness2")
		fmt.Printf("bussiness2 stats: %+v \n", stats)

	})

	// 创建自定义的 性能记录
	It("test other", Serial, Label("measurement"), func() {
		// we create a new experiment
		// https://pkg.go.dev/github.com/onsi/gomega@v1.18.1/gmeasure#NewExperiment
		experiment := gmeasure.NewExperiment("performance test other")
		// 把性能数据 追加到 测试报告中
		AddReportEntry(experiment.Name, experiment)

		// 创建自定义的时间 Sample
		// https://pkg.go.dev/github.com/onsi/gomega@v1.18.1/gmeasure#Experiment.RecordDuration
		// 我们可以使用其它手段，进行进行测试，然后把 时间测试结果  创建到  experiment 中
		experiment.RecordDuration("my other performance test1", 3*time.Second)
		// 可追加修饰符
		experiment.RecordDuration("my other performance test1", 5*time.Second, gmeasure.Annotation("sample 2nd"))

		// 创建自定义的数值 Sample
		// 我们可以使用其它手段，进行进行测试，然后把 数值测试结果的样本  创建到  experiment 中
		// https://pkg.go.dev/github.com/onsi/gomega@v1.18.1/gmeasure#Experiment.RecordValue
		experiment.RecordValue("length", 6.5)
		experiment.RecordValue("length", 7)
		// 可以追加一些修饰符，gmeasure.Units表示数值的计量单位，gmeasure.Style表示显示颜色，gmeasure.Precision表示小数点后的个数，gmeasure.Annotation 可以添加注释
		experiment.RecordValue("length", 3.141, gmeasure.Units("cm"), gmeasure.Style("{{blue}}"), gmeasure.Precision(2), gmeasure.Annotation("box A"))

	})

	// 测量 一段 业务过程中的 性能
	It("test other", Serial, Label("measurement"), func() {
		experiment := gmeasure.NewExperiment("end-to-end web-server performance")
		AddReportEntry(experiment.Name, experiment)
		stopwatch := experiment.NewStopwatch() // start the stopwatch

		stopwatch.Reset() // reset the stopwatch
		fmt.Printf("do bussiness1 here\n")
		time.Sleep(1 * time.Second)
		stopwatch.Record("step1")

		fmt.Printf("do bussiness2 here\n")
		time.Sleep(2 * time.Second)
		stopwatch.Record("step2")

	})

})
