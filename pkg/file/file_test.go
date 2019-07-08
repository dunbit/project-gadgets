package file_test

import (
	"io/ioutil"
	"os"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/ginkgo/extensions/table"
	. "github.com/onsi/gomega"

	. "github.com/dunbit/project-gadgets/pkg/file"
	"github.com/dunbit/project-gadgets/pkg/license"
)

var exampleLicense = license.License{
	Data: []string{
		"Copyright 2019 Authors of project-gadgets",
		"",
		"Licensed under the Apache License, Version 2.0 (the \"License\");",
		"you may not use this file except in compliance with the License.",
		"You may obtain a copy of the License at",
		"",
		"    http://www.apache.org/licenses/LICENSE-2.0",
		"",
		"Unless required by applicable law or agreed to in writing, software",
		"distributed under the License is distributed on an \"AS IS\" BASIS,",
		"WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.",
		"See the License for the specific language governing permissions and",
		"limitations under the License.",
	},
}

var exampleGoFile = `// Copyright 2019 Authors of project-gadgets
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package main

func main() {
	
}
`

var exampleMakeFile = `# Copyright 2019 Authors of project-gadgets
#
# Licensed under the Apache License, Version 2.0 (the "License");
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at
#
#     http://www.apache.org/licenses/LICENSE-2.0
#
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.

.PHONY: all
all:
	@ echo hello world
`

func createTempFile(ext string, data string) (*os.File, error) {
	file, err := ioutil.TempFile(os.TempDir(), "*"+ext)
	if err != nil {
		return nil, err
	}

	_, err = file.WriteString(data)
	if err != nil {
		return nil, err
	}

	err = file.Close()
	if err != nil {
		return nil, err
	}

	return file, nil
}

var _ = Describe("File", func() {

	Describe("New", func() {

		It("Should create a new File object", func() {
			f := New("/my/path/to/file.txt", "//")

			Expect(f).NotTo(BeNil())
		})
	})

	DescribeTable("ReadLicense",
		func(ext string, comment string, data string, expected *license.License) {
			file, err := createTempFile(ext, data)
			defer func() {
				err := os.Remove(file.Name())
				Expect(err).NotTo(HaveOccurred())
			}()
			Expect(err).NotTo(HaveOccurred())

			sourceFile := New(file.Name(), comment)
			license, err := sourceFile.ReadLicense()
			Expect(err).NotTo(HaveOccurred())

			Expect(license).To(Equal(expected))
		},
		Entry(".go file", ".go", "//", exampleGoFile, &exampleLicense),
		Entry("Makefile file", "Makefile", "#", exampleMakeFile, &exampleLicense),
	)

	Describe("ReadLicense", func() {

		It("Should fail if path does not exixts", func() {
			sourceFile := New("my/fake/path.txt", "#")
			license, err := sourceFile.ReadLicense()
			Expect(err).To(HaveOccurred())
			Expect(license).To(BeNil())
		})
	})
})
