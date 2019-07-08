package comment_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/ginkgo/extensions/table"
	. "github.com/onsi/gomega"

	. "github.com/dunbit/project-gadgets/pkg/comment"
)

var _ = Describe("Comment", func() {

	Describe("Add", func() {

		DescribeTable("Should add the comment at the beginning of the line",
			func(comment string, data string, expected string) {
				s := Add(comment, data)

				Expect(s).To(Equal(expected))
			},
			Entry("// Hello", "//", "Hello", "// Hello"),
			Entry("# Hello", "#", "Hello", "# Hello"),
			Entry("//", "//", "", "//"),
			Entry("#", "#", "", "#"),
			Entry("//     Hello", "//", "    Hello", "//     Hello"),
			Entry("#      Hello", "#", "    Hello", "#     Hello"),
		)
	})

	Describe("Strip", func() {

		DescribeTable("Should remove the comment from the beginning of the line",
			func(comment string, data string, expected string) {
				s := Strip(comment, data)

				Expect(s).To(Equal(expected))
			},
			Entry("// Hello", "//", "// Hello", "Hello"),
			Entry("# Hello", "#", "# Hello", "Hello"),
			Entry("//", "//", "//", ""),
			Entry("#", "#", "#", ""),
			Entry("//     Hello", "//", "//     Hello", "    Hello"),
			Entry("#      Hello", "#", "#     Hello", "    Hello"),
			Entry("Hello", "//", "Hello", "Hello"),
			Entry("Hello", "#", "Hello", "Hello"),
		)
	})
})
