package main_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/onsi/gomega/gexec"

	"testing"
	"time"
)

var bleatPath string

func TestBleat(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Bleat Suite")
}

var _ = BeforeSuite(func() {
	path, err := gexec.Build("github.com/yinlogn/bleat")
	Expect(err).ToNot(HaveOccurred())
	bleatPath = path
	SetDefaultEventuallyTimeout(15 * time.Second)
	SetDefaultEventuallyPollingInterval(100 * time.Millisecond)
	SetDefaultConsistentlyDuration(1 * time.Second)
	SetDefaultConsistentlyPollingInterval(10 * time.Millisecond)
})

var _ = AfterSuite(func() {
	gexec.CleanupBuildArtifacts()
})
