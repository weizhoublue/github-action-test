package print_test

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"time"

	"github.com/weizhoublue/github-action-test/pkg/print"
)

var _ = Describe("Print", Label("unitest"), func() {
	It("test output", func() {
		print.MyPrint()
		print.MyPrint()

		time.Sleep(10 * time.Second)
		Expect(1).To(Equal(1))
	})
})
