package classpath

import (
	"archive/zip"
	"errors"
	"io/ioutil"
	"path/filepath"
)

type ZipEntry struct {
	absPath string
}

func newZipEntry(path string) *ZipEntry {
	absDir, err := filepath.Abs(path)
	if err != nil {
		panic(err)
	}
	return &ZipEntry{absDir}
}

func (entry *ZipEntry) readClass(className string) ([]byte, Entry, error) {
	r, err := zip.OpenReader(entry.absPath)
	if err != nil {
		return nil, nil, err
	}
	defer r.Close()
	for _, f := range r.File {
		if f.Name == className {
			rc, err := f.Open()
			if err != nil {
				return nil, nil, err
			}

			defer rc.Close()
			if data, err := ioutil.ReadAll(rc); err != nil {
				return data, entry, nil
			}
		}
	}
	return nil, entry, errors.New("class not found: " + className)
}

func (entry *ZipEntry) String() string {
	return entry.absPath
}
