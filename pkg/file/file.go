package file

import (
	"bufio"
	"os"
	"strings"

	"github.com/dunbit/project-gadgets/pkg/license"
	"github.com/hashicorp/go-multierror"
)

// File represents a file that could contain a license, or will need a license
type File struct {
	Path    string
	Comment string
}

// New ...
func New(path string, comment string) *File {
	return &File{
		path,
		comment,
	}
}

// ReadLicense ...
func (f *File) ReadLicense() (l *license.License, rerr error) {
	license := new(license.License)

	file, err := os.Open(f.Path)
	if err != nil {
		return nil, err
	}
	defer func() {
		err = file.Close()
		if err != nil {
			rerr = multierror.Append(rerr, err)
		}
	}()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()

		if line == "" {
			break
		}

		line = strings.TrimPrefix(line, f.Comment)
		line = strings.TrimPrefix(line, " ")

		license.AppendLine(line)
	}

	return license, nil
}
