package main_test

import (
	"github.com/yinlogn/bleat/deke"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	. "github.com/onsi/gomega/gexec"

	"net/http"
	"os/exec"
	"time"
)

var _ = Describe("Bleat functionality", func() {
	var bleatSession *Session

	BeforeEach(func() {
		deke.RespondWith("You should not see me")

		bleatStartCommand := exec.Command(bleatPath)

		session, err := Start(bleatStartCommand, GinkgoWriter, GinkgoWriter)
		Expect(err).ToNot(HaveOccurred())
		bleatSession = session

		time.Sleep(time.Millisecond * 50)
	})

	AfterEach(func() {
		deke.Close()

		bleatSession.Terminate()
		Eventually(bleatSession).Should(Exit())
	})

	It("should call the deke server and return modified data", func() {
		resp, err := http.Get("http://127.0.0.1:5000")

		Expect(resp).ToNot(BeNil())
		Expect(err).ToNot(HaveOccurred())
	})
})
