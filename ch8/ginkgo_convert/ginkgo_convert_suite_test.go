package main_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestGinkgoConvert(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "GinkgoConvert Suite")
}
