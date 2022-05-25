package openrtb

import (
	"encoding/json"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Native", func() {
	It("should parse correctly", func() {
		var subject *Native
		Expect(fixture("native", &subject)).To(Succeed())
		Expect(subject).To(Equal(&Native{
			Request: json.RawMessage(`"PAYLOAD"`),
			Version: "2",
		}))
	})

	It("should parse asian correctly", func() {
		var (
			subject *Native
			expect  json.RawMessage
		)
		Expect(fixture("native_inline", &subject)).To(Succeed())
		Expect(fixture("native_inline", &expect)).To(Succeed())
		Expect(subject).To(Equal(&Native{
			Request: expect,
		}))
	})
})
