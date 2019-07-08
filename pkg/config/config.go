package config

import (
	"bytes"
	"encoding/json"
	"errors"
	"io"
	"os"
	"path/filepath"

	"github.com/hashicorp/go-multierror"
	"gopkg.in/yaml.v2"
)

// Error definitions
var (
	ErrFilePathInvalid = errors.New("File Path invalid")
	ErrFileTypeInvalid = errors.New("File Type invalid")
	ErrDataInvalid     = errors.New("Data invalid")
)

// License ...
type License struct {
	Path    string `yaml:"path" json:"path"`
	Content string `yaml:"content" json:"content"`
}

// FileType ...
type FileType struct {
	Match   string `yaml:"match" json:"match"`
	Comment string `yaml:"comment" json:"comment"`
}

// Config ...
type Config struct {
	License  License    `yaml:"license" json:"license"`
	Files    []FileType `yaml:"files" json:"files"`
	Excludes []string   `yaml:"excludes" json:"excludes"`
}

var types = [3]string{".json", ".yml", ".yaml"}

func isSupported(ext string) bool {
	for _, t := range types {
		if t == ext {
			return true
		}
	}
	return false
}

// ReadFile ...
func ReadFile(path string) (c *Config, rerr error) {
	if path == "" {
		return nil, ErrFilePathInvalid
	}

	if !isSupported(filepath.Ext(path)) {
		return nil, ErrFileTypeInvalid
	}

	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer func() {
		err = file.Close()
		if err != nil {
			rerr = multierror.Append(rerr, err)
		}
	}()

	var config *Config

	switch filepath.Ext(path) {
	case ".json":
		config, err = ReadIOJSON(file)
	case ".yml":
		fallthrough
	case ".yaml":
		config, err = ReadIOYAML(file)
	}
	if err != nil {
		return nil, err
	}

	return config, nil
}

// ReadString ...
func ReadString(data string) (*Config, error) {
	if data == "" {
		return nil, ErrDataInvalid
	}

	buffer := bytes.NewBufferString(data)

	return ReadIOJSON(buffer)
}

// ReadIOYAML ...
func ReadIOYAML(reader io.Reader) (*Config, error) {
	config := new(Config)
	err := yaml.NewDecoder(reader).Decode(config)

	return config, err
}

// ReadIOJSON ...
func ReadIOJSON(reader io.Reader) (*Config, error) {
	config := new(Config)
	err := json.NewDecoder(reader).Decode(config)

	return config, err
}
