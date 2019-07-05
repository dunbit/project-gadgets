package config

import (
	"bytes"
	"encoding/json"
	"errors"
	"io"
	"os"
	"path/filepath"

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

// ReadFile ...
func ReadFile(path string) (*Config, error) {
	if path == "" {
		return nil, ErrFilePathInvalid
	}

	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}

	switch filepath.Ext(path) {
	case ".json":
		return ReadIOJSON(file)
	case ".yml":
		fallthrough
	case ".yaml":
		return ReadIOYAML(file)
	}

	return nil, ErrFileTypeInvalid
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
