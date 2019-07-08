package config_test

import (
	"bytes"
	"io/ioutil"
	"os"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	. "github.com/dunbit/project-gadgets/pkg/config"

	"github.com/tdewolff/minify/v2"
	"github.com/tdewolff/minify/v2/json"
)

var exampleConfig = &Config{
	License: License{
		Path: "/my/path/to/license.txt",
		Content: `
			MY AWESOME LICENSE
				ANNO 2019
			ALL RIGHTS RESERVED
		`,
	},
	Files: []FileType{
		{Comment: "#", Match: "Makefile"},
		{Comment: "//", Match: "*.go"},
	},
	Excludes: []string{
		"./vendor",
	},
}

const exampleConfigYAML = `
license:
  path: /my/path/to/license.txt
  content: "\n\t\t\tMY AWESOME LICENSE\n\t\t\t\tANNO 2019\n\t\t\tALL RIGHTS RESERVED\n\t\t"
files:
  - match: Makefile
    comment: '#'
  - match: '*.go'
    comment: //
excludes:
  - ./vendor
`

const exampleConfigJSON = `
{
	"license": {
		"path": "/my/path/to/license.txt",
		"content": "\n\t\t\tMY AWESOME LICENSE\n\t\t\t\tANNO 2019\n\t\t\tALL RIGHTS RESERVED\n\t\t"
	},
	"files": [
		{
		"match": "Makefile",
		"comment": "#"
		},
		{
		"match": "*.go",
		"comment": "//"
		}
	],
	"excludes": [
		"./vendor"
	]
}
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

	err = file.Sync()
	if err != nil {
		return nil, err
	}

	err = file.Close()
	if err != nil {
		return nil, err
	}

	return file, nil
}

func minifyJson(data string) (string, error) {
	m := minify.New()
	m.AddFunc("application/json", json.Minify)
	s, err := m.String("application/json", data)
	if err != nil {
		return "", err
	}
	return s, nil
}

var _ = Describe("Config", func() {

	Describe("ReadFile", func() {

		It("Should parse a yml file", func() {
			file, err := createTempFile(".yml", exampleConfigYAML)
			defer func() {
				err := os.Remove(file.Name())
				Expect(err).NotTo(HaveOccurred())
			}()
			Expect(err).NotTo(HaveOccurred())

			config, err := ReadFile(file.Name())
			Expect(err).NotTo(HaveOccurred())
			Expect(config).To(Equal(exampleConfig))
		})

		It("Should parse a yaml file", func() {
			file, err := createTempFile(".yaml", exampleConfigYAML)
			defer func() {
				err := os.Remove(file.Name())
				Expect(err).NotTo(HaveOccurred())
			}()
			Expect(err).NotTo(HaveOccurred())

			config, err := ReadFile(file.Name())
			Expect(err).NotTo(HaveOccurred())
			Expect(config).To(Equal(exampleConfig))
		})

		It("Should parse a json file", func() {
			file, err := createTempFile(".json", exampleConfigJSON)
			defer func() {
				err := os.Remove(file.Name())
				Expect(err).NotTo(HaveOccurred())
			}()
			Expect(err).NotTo(HaveOccurred())

			config, err := ReadFile(file.Name())
			Expect(err).NotTo(HaveOccurred())
			Expect(config).To(Equal(exampleConfig))
		})

		It("Should return err for empty path", func() {
			config, err := ReadFile("")
			Expect(err).To(Equal(ErrFilePathInvalid))
			Expect(config).To(BeNil())
		})

		It("Should return err for non existing path", func() {
			config, err := ReadFile("my/dir/config.json")
			Expect(err).To(HaveOccurred())
			Expect(config).To(BeNil())
		})

		It("Should return err for non supported file type", func() {
			config, err := ReadFile("/my/unsupported/file.ini")
			Expect(err).To(Equal(ErrFileTypeInvalid))
			Expect(config).To(BeNil())
		})

		It("Should return err if content of file is not json or yaml", func() {
			file, err := createTempFile(".json", "some unsupported data")
			defer func() {
				err := os.Remove(file.Name())
				Expect(err).NotTo(HaveOccurred())
			}()
			Expect(err).NotTo(HaveOccurred())

			config, err := ReadFile(file.Name())
			Expect(err).To(HaveOccurred())
			Expect(config).To(BeNil())
		})
	})

	Describe("ReadString", func() {

		It("Should parse a json string", func() {
			data, err := minifyJson(exampleConfigJSON)
			Expect(err).NotTo(HaveOccurred())

			config, err := ReadString(data)
			Expect(err).NotTo(HaveOccurred())
			Expect(config).To(Equal(exampleConfig))
		})

		It("Should return error with string", func() {
			data := ""

			config, err := ReadString(data)
			Expect(err).To(Equal(ErrDataInvalid))
			Expect(config).To(BeNil())
		})
	})

	Describe("ReadIOYAML", func() {

		It("Should parse a config from a reader", func() {
			buffer := new(bytes.Buffer)
			buffer.WriteString(exampleConfigYAML)

			config, err := ReadIOYAML(buffer)
			Expect(err).NotTo(HaveOccurred())
			Expect(config).To(Equal(exampleConfig))
		})
	})

	Describe("ReadIOJSON", func() {

		It("Should parse a config from a reader", func() {
			buffer := new(bytes.Buffer)
			buffer.WriteString(exampleConfigJSON)

			config, err := ReadIOJSON(buffer)
			Expect(err).NotTo(HaveOccurred())
			Expect(config).To(Equal(exampleConfig))
		})
	})
})
