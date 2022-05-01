package print_test

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	//"github.com/weizhoublue/github-action-test/pkg/lock"
	"go.uber.org/goleak"
	"fmt"
	"time"
)

// =========  测试 业务 是否测试过程中 有 协程泄露

func FakeLeak(){
	go func(){
		fmt.Printf("I am sleeping")
		time.Sleep(10*time.Second)
	}()
}


var _ = Describe("Goleak",Label("leak_root"), func() {

	
	BeforeEach(func(){
		// 在 BeforeEach 中 使用时，可以检测每个IT的泄露情况，但缺点是 如果 ginkgo 使用了 -p 并发测试用例,会误报 其它IT 的并发协程
		// 当 ginkgo 并发时， goleak 同样不适合在  BeforeSuite 运行，因为会在大量的 process 中 并发多协程（不是it），会误报 大量 的协程
		currentGRs := goleak.IgnoreCurrent()
		DeferCleanup(func() {
			Eventually(func() bool { 
				e:=goleak.Find(currentGRs) 
				if e!=nil {
					GinkgoWriter.Printf("goroutine leak : %+v \n", e)
					return false
				}
				return true
			}).Should(BeTrue())
		})
	})
	
	
	
	It("test golang routine leak", Label("leak1"), func(){
		
		FakeLeak()
		time.Sleep(3*time.Second)
		
	})
	
	It("test golang routine leak", Label("leak1"), func(){

		FakeLeak()
		time.Sleep(15*time.Second)

	})
})


