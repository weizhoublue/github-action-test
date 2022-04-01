package print_test

import (
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestPrint(t *testing.T) {

	RegisterFailHandler(Fail)

	RunSpecs(t, "Print Suite")

}
