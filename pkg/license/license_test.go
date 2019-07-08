package license_test

import (
	"fmt"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	. "github.com/dunbit/project-gadgets/pkg/license"
)

var _ = Describe("License", func() {

	Describe("AppendLine", func() {

		It("Should be nil initially", func() {
			license := License{}
			Expect(len(license.Data)).To(Equal(0))
		})

		It("Should be one, when add a single line", func() {
			license := License{}
			license.AppendLine("a line")

			Expect(len(license.Data)).To(Equal(1))
		})

		It("Should be ten, when adding 10 lines", func() {
			license := License{}
			for i := 0; i < 10; i++ {
				license.AppendLine("a line")
			}
			Expect(len(license.Data)).To(Equal(10))
		})
	})

	Describe("Lines", func() {

		It("Should return the amount of lines", func() {
			lines := 10
			data := []string{}
			for i := 0; i < lines; i++ {
				data = append(data, fmt.Sprintf("line %d", i))
			}

			license := License{
				Data: data,
			}

			Expect(license.Lines()).To(Equal(lines))
		})
	})
})
