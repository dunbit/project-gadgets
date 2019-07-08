package walker_test

import (
	"io/ioutil"
	"os"
	"path"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/dunbit/project-gadgets/pkg/config"
	. "github.com/dunbit/project-gadgets/pkg/walker"
)

var gofile = `package myPackage`

var gofileWithLicense = `// Copyright 2019 Authors of protobuf-gadgets
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

package myPackage
`

var gofileWithWrongLicense = `// Copyright 1792 Authors of protobuf-gadgets
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

package myPackage
`

var makeFile = `.PHONY: all
all:
	@ echo test
`
var makeFileWithLicense = `# Copyright 2019 Authors of project-gadgets
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
	@ echo test
`

var makeFileWithWrongLicense = `# Copyright 3076 Authors of project-gadgets
# limitations under the License.

.PHONY: all
all:
	@ echo test
`

var tempDir string

var testData = dir{
	Name: "testdata",
	ChildDirs: []*dir{
		// All valid licenses
		{
			Name: "dir1",
			ChildFiles: []*file{
				{
					Name: "gofile.go",
					Data: []byte(gofileWithLicense),
				},
				{
					Name: "Makefile",
					Data: []byte(makeFileWithLicense),
				},
			},
			ChildDirs: []*dir{
				{
					Name: "dir1_1",
					ChildFiles: []*file{
						{
							Name: "gofile.go",
							Data: []byte(gofileWithLicense),
						},
						{
							Name: "Makefile",
							Data: []byte(makeFileWithLicense),
						},
					},
				},
			},
		},
		// go file is missing a license
		{
			Name: "dir2",
			ChildFiles: []*file{
				{
					Name: "gofile.go",
					Data: []byte(gofile),
				},
				{
					Name: "Makefile",
					Data: []byte(makeFileWithLicense),
				},
			},
		},
		// go file has wrong license
		{
			Name: "dir3",
			ChildFiles: []*file{
				{
					Name: "gofile.go",
					Data: []byte(gofileWithWrongLicense),
				},
				{
					Name: "Makefile",
					Data: []byte(makeFileWithLicense),
				},
			},
		},
		// Makefile is missing a license
		{
			Name: "dir4",
			ChildFiles: []*file{
				{
					Name: "gofile.go",
					Data: []byte(gofileWithLicense),
				},
				{
					Name: "Makefile",
					Data: []byte(makeFile),
				},
			},
		},
		// Makefile has a wrong license
		{
			Name: "dir5",
			ChildFiles: []*file{
				{
					Name: "gofile.go",
					Data: []byte(gofileWithLicense),
				},
				{
					Name: "Makefile",
					Data: []byte(makeFileWithWrongLicense),
				},
			},
		},
		// go file and Makefile are missing a license
		{
			Name: "dir6",
			ChildFiles: []*file{
				{
					Name: "gofile.go",
					Data: []byte(gofile),
				},
				{
					Name: "Makefile",
					Data: []byte(makeFile),
				},
			},
		},
		// go file and Makefile have wrong license
		{
			Name: "dir7",
			ChildFiles: []*file{
				{
					Name: "gofile.go",
					Data: []byte(gofileWithWrongLicense),
				},
				{
					Name: "Makefile",
					Data: []byte(makeFileWithWrongLicense),
				},
			},
		},
	},
}

type file struct {
	Name string
	Data []byte
}

type dir struct {
	Name       string
	ChildDirs  []*dir
	ChildFiles []*file
}

func mkFile(wd string, f *file) error {
	file, err := os.Create(path.Join(wd, f.Name))
	if err != nil {
		return err
	}
	_, err = file.Write(f.Data)
	if err != nil {
		return err
	}
	err = file.Close()
	if err != nil {
		return err
	}

	return nil
}

func mkDir(wd string, d *dir) error {
	err := os.Mkdir(path.Join(wd, d.Name), 0777)
	if err != nil {
		return err
	}

	for _, childFile := range d.ChildFiles {
		err := mkFile(path.Join(wd, d.Name), childFile)
		if err != nil {
			return err
		}
	}

	for _, childDir := range d.ChildDirs {
		err := mkDir(path.Join(wd, d.Name), childDir)
		if err != nil {
			return err
		}
	}

	return nil
}

func createTestData() {
	var err error

	tempDir, err = ioutil.TempDir(os.TempDir(), "project_gadgets_")
	if err != nil {
		panic(err)
	}

	err = mkDir(tempDir, &testData)
	if err != nil {
		panic(err)
	}
}

func cleanTestData() {
	err := os.RemoveAll(tempDir)
	if err != nil {
		panic(err)
	}
}

var _ = Describe("Walker", func() {

	BeforeSuite(func() {
		createTestData()
	})

	AfterSuite(func() {
		cleanTestData()
	})

	Describe("Walk", func() {

		It("Should return all matching files", func() {
			c := &config.Config{
				Files: []config.FileType{
					{Match: "**/*.go", Comment: "//"},
					{Match: "**/Makefile", Comment: "#"},
				},
			}

			files, err := Walk(c, path.Join(tempDir, "testdata", "dir1"))
			Expect(err).NotTo(HaveOccurred())
			Expect(files).To(HaveLen(4))
		})

		It("Should return 0 files and err, if dir is not existing", func() {
			c := &config.Config{
				Files: []config.FileType{
					{Match: "**/*.go", Comment: "//"},
					{Match: "**/Makefile", Comment: "#"},
				},
			}

			files, err := Walk(c, path.Join(tempDir, "testdata", "somedir"))
			Expect(err).To(Equal(os.ErrNotExist))
			Expect(files).To(HaveLen(0))
		})

		It("Should return 2 files, if excludes is set", func() {
			c := &config.Config{
				Files: []config.FileType{
					{Match: "**/*.go", Comment: "//"},
					{Match: "**/Makefile", Comment: "#"},
				},
				Excludes: []string{
					"**/dir1_1",
				},
			}

			files, err := Walk(c, path.Join(tempDir, "testdata", "dir1"))
			Expect(err).NotTo(HaveOccurred())
			Expect(files).To(HaveLen(2))
		})
	})
})
