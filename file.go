package sailor

import (
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
)

const (
	FilePerm600 os.FileMode = 0600 // For secret files.
	FilePerm644 os.FileMode = 0644 // For normal files.
	FilePerm755 os.FileMode = 0755 // For directory or execute files.
)

func Exists(path string) bool {
	_, err := os.Stat(path)
	return !os.IsNotExist(err)
}

func EnsureFolderExists(folder string) {
	if Exists(folder) {
		return
	}

	err := os.MkdirAll(folder, FilePerm755)
	if err != nil {
		log.Fatal("directory not exists, err: ", err)
	}
}

func EnsureFileExists(path string) {
	if Exists(path) {
		return
	}

	EnsureFolderExists(filepath.Dir(path))

	err := ioutil.WriteFile(path, nil, FilePerm644)
	if err != nil {
		log.Fatal("file not exists, err: ", err)
	}
}

func WriteFile(path string, content string) error {
	EnsureFileExists(path)

	err := ioutil.WriteFile(path, StringToBytes(content), FilePerm644)
	if err != nil {
		return err
	}

	return nil
}
