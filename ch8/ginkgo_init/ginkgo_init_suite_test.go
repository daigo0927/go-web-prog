package main_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestGinkgoInit(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "GinkgoInit Suite")
}
