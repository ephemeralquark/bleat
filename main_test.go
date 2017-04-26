package main_test

import (
	"github.com/yinlogn/bleat/deke"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	. "github.com/onsi/gomega/gexec"

	"net/http"
	"os/exec"
	"syscall"
)

var _ = Describe("Bleat functionality", func() {
	var bleatSession *Session

	BeforeEach(func() {
		deke.RespondWith("You should not see me")

		bleatStartCommand := exec.Command(bleatPath)
		session, err := Start(bleatStartCommand, GinkgoWriter, GinkgoWriter)
		Expect(err).ToNot(HaveOccurred())
		bleatSession = session
	})

	AfterEach(func() {
		deke.Close()

		err := bleatSession.Command.Process.Signal(syscall.SIGTERM)
		Expect(err).ToNot(HaveOccurred())
		Eventually(bleatSession, 5).Should(Exit(143))
	})

	It("should call the deke server and return modified data", func() {
		resp, err := http.Get("http://127.0.0.1:5000")

		Expect(err).ToNot(HaveOccurred())
		Expect(resp).ToNot(BeNil())
	})
})
