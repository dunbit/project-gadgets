package comments_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/ginkgo/extensions/table"
	. "github.com/onsi/gomega"

	. "github.com/dunbit/project-gadgets/pkg/comments"
)

var _ = Describe("Comments", func() {

	Describe("Add", func() {

		DescribeTable("Should add the comments at the beginning of the line",
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

		DescribeTable("Should remove the comments from the beginning of the line",
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
		)
	})
})
