package deke_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
)

func TestDeke(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Deke Suite")
}
