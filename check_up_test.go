package mel

import (
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestDataValidator(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Test")
}

var _ = Describe("Data Validator tests", func() {
	When("Testing function ValidCpfFormat", func() {
		Context("and is correct fields", func() {
			It("when the function receives a valid cpf with the WITH_FORMATTING type, it must return true", func() {
				cpfmock := "123.456.789-11"
				res := ValidCpfFormat(cpfmock, WITH_FORMATTING)
				Expect(res).Should(BeTrue())
			})

			It("when the function receives a valid cpf with the WITHOUT_FORMATTING type, it must return true", func() {
				cpfmock := "12345678910"
				res := ValidCpfFormat(cpfmock, WITHOUT_FORMATTING)
				Expect(res).Should(BeTrue())
			})
		})

		Context("and is wrong fields", func() {
			It("when the function receives a invalid data with type WITH_FORMATTING, it must return false", func() {
				cpfmock := "123456.789-11"
				res := ValidCpfFormat(cpfmock, WITH_FORMATTING)
				Expect(res).Should(BeFalse())
			})

			It("when the function receives a invalid data with type WITHOUT_FORMATTING, it must return false", func() {
				cpfmocksFail := []string{"12345678g11", "1234", "1234527483912"}

				for _, v := range cpfmocksFail {
					res := ValidCpfFormat(v, WITHOUT_FORMATTING)
					Expect(res).Should(BeFalse())
				}
			})
		})
	})

	When("Testing function CleanCpf", func() {
		Context("and is correct fields", func() {
			It("when the function receives a valid cpf with the WITH_FORMATTING type, it must return true", func() {
				cpfmock := "123.456.789-11"
				res, err := CleanCpf(cpfmock)
				Expect(res).Should(Equal("12345678911"))
				Expect(err).Should(BeNil())
			})
		})

		Context("and is wrong fields", func() {
		})
	})
})
