package print_test

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"github.com/weizhoublue/github-action-test/pkg/print"
)

var _ = Describe("Print", Label("unitest"), func() {
	It("test output", func() {
		print.MyPrint()
		print.MyPrint()

		print.MyPrint()

		Expect(1).To(Equal(1))
	})
})
