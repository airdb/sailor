package fileutil

import (
	"crypto/md5"
	"encoding/hex"
	"io"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"
	"time"
)

func MTime(name string) (mtime time.Time) {
	fi, err := os.Stat(name)
	if err != nil {
		return
	}
	mtime = fi.ModTime()
	// log.Println(mtime)
	return
}

func ComputeMd5(filePath string) string {
	// func ComputeMd5(filePath string) ([]byte, error) {
	var result []byte
	file, err := os.Open(filePath)
	if err != nil {
		//return result, err
		return hex.EncodeToString(result)
	}
	defer file.Close()
	hash := md5.New()
	if _, err := io.Copy(hash, file); err != nil {
		// return result, err
		return hex.EncodeToString(result)
	}
	// return hash.Sum(result), nil
	return hex.EncodeToString(hash.Sum(result))
}

func FileRename(filename, dstname string) {
	err := os.Rename(filename, dstname)
	if err != nil {
		log.Printf("filename: %s rename to %s failed. err: %v\n", filename, dstname, err)
	} else {
		log.Printf("filename: %s rename to %s success.\n", filename, dstname)
	}
}

func Substr(s string, pos, length int) string {
	runes := []rune(s)
	l := pos + length
	if l > len(runes) {
		l = len(runes)
	}
	return string(runes[pos : l+1])
}

func GetFilePath(filepath string) (string, string) {
	pdir := Substr(filepath, 0, strings.LastIndex(filepath, "/"))
	filename := strings.Replace(filepath, pdir, "", -1)
	return pdir, filename
}

func GetFileSize(path string) (fileSize int64) {
	fileInfo, err := os.Stat(path)
	if err != nil {
		log.Println(err)
		return
		// panic(err)
	}
	fileSize = fileInfo.Size()
	return fileSize
}

func CountFileLine(filename string) (count int64) {
	count = 0
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		return
	}
	for i := 0; i < len(data); i++ {
		if data[i] == '\n' {
			count++
		}
	}
	log.Println("filename:", filename, ", line:", count)
	return
}

func WalkDir(dirPth, suffix string) (files []string, err error) {
	files = make([]string, 0, 30)
	suffix = strings.ToUpper(suffix)
	err = filepath.Walk(dirPth, func(filename string, fi os.FileInfo, err error) error {
		//if err != nil {
		// return err
		//}
		if fi.IsDir() {
			return nil
		}
		if strings.HasSuffix(strings.ToUpper(fi.Name()), suffix) {
			files = append(files, filename)
		}
		return nil
	})
	return
}

func WalkDirPath(dirPth string) (paths []string, err error) {
	paths = make([]string, 0, 30)
	err = filepath.Walk(dirPth, func(filename string, fi os.FileInfo, err error) error {
		if fi.IsDir() {
			paths = append(paths, filename)
		}
		return nil
	})
	return
}

// return the source filename after the last slash
func ChopPath(original string) string {
	i := strings.LastIndex(original, "/")
	if i == -1 {
		return original
	} else {
		return original[i+1:]
	}
}

