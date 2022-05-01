package print_test

import (
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	//"go.uber.org/goleak"
	"time"
)

func TestPrint(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Print Suite")

}

var _ = BeforeSuite(func() {
	GinkgoWriter.Printf("welan  BeforeSuite: %+v  , on process %v \n", time.Now() , GinkgoParallelProcess() )
	
	
	/*
	// 当 ginkgo 并发时， goleak 同样不适合在  BeforeSuite 运行，因为会在大量的 process 中 并发多协程（不是it），会误报 大量 的协程
	currentGRs := goleak.IgnoreCurrent()
	DeferCleanup(func() {
		Eventually(func() bool {
			e := goleak.Find(currentGRs)
			if e != nil {
				GinkgoWriter.Printf("goroutine leak : %+v \n", e)
				return false
			}
			return true
		}).Should(BeTrue())
	})
	*/
	 
})

var _ = AfterSuite(func() {
	GinkgoWriter.Printf("welan  AfterSuite: %+v  , on process %v \n", time.Now() , GinkgoParallelProcess() )
})