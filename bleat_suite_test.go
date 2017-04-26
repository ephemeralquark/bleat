package main_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
)

func TestBleat(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Bleat Suite")
}
