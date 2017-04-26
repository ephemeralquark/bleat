package deke_test

import (
	decoyServer "github.com/yinlogn/bleat/deke"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"io/ioutil"
	"net/http"
)

var _ = Describe("Deke", func() {
	Context("when set up to respond with a fake response", func() {
		BeforeEach(func() {
			decoyServer.RespondWith("General fake response")
		})

		AfterEach(func() {
			decoyServer.Close()
		})

		It("should return the instantiated server's URL", func() {
			Expect(decoyServer.Url()).NotTo(Equal(nil))
		})

		It("should return the set up response", func() {
			resp, err := http.Get(decoyServer.Url())
			Expect(err).To(BeNil())
			Expect(resp.StatusCode).To(Equal(200))

			body, err := ioutil.ReadAll(resp.Body)
			Expect(err).To(BeNil())

			Expect(string(body)).To(Equal("General fake response"))
		})
	})

})
