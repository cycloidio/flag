package flag

import (
	"fmt"
	"os"
	"path/filepath"
)

type File string

func (f *File) UnmarshalFlag(value string) error {
    // backup value used in case of refresh
	*f = File(value)

	stat, err := os.Stat(value)
	if err != nil {
		return err
	}

	if stat.IsDir() {
		return fmt.Errorf("path '%s' is a directory, not a file", value)
	}

	abs, err := filepath.Abs(value)
	if err != nil {
		return err
	}

	*f = File(abs)

	return nil
}

func (f File) Path() string {
	return string(f)
}

// Reload reloads the value of the Keys
func (f File) Reload() error {
	return f.UnmarshalFlag(f.Path())
}
