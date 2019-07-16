package sailor

import (
	"crypto/md5"
	"encoding/hex"
	"io"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"
	"syscall"
	"time"

	"github.com/airdb/sailor/check"
)

func WriteFile(filename string, content string) error {
	f, err := os.OpenFile(filename, os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Printf("open file failed, filename: %s, err: %v\n", filename, err)
	}

	defer f.Close()

	if _, err = f.WriteString(content); err != nil {
		log.Printf("write file failed, filename: %s\n", filename)
	} else {
		log.Println("write file successfully, filename: ", filename)
	}

	return err
}

func WriteByteToFile(dst string, d []byte) error {
	pdir, filename := filepath.Split(dst)

	if !check.IsFileExists(pdir) {
		err := os.MkdirAll(pdir, 0755)
		if err != nil {
			log.Printf("%s", err)
		} else {
			log.Println("Create Directory OK!")
		}
	}

	log.Printf("WriteFile filename: %s,  Size of download: %d\n", filename, len(d))
	err := ioutil.WriteFile(dst, d, 0644)
	if err != nil {
		log.Fatal(err)
	}
	return err
}

func ChangMtime(filename string, mtime time.Time) error {
	atime, _, _, err := StatTimes(filename)
	if err != nil {
		log.Println(err)
		return err
	}

	err = os.Chtimes(filename, atime, mtime)
	if err != nil {
		log.Printf("filename: %s, change mtime to [%s] failed!  %v\n", filename, mtime, err)
	} else {
		log.Printf("filename: %s, change mtime to [%s] successfully!\n", filename, mtime)
	}
	return err
}

func CopyFileTime(srcfile, dstfile string) (err error) {
	atime, mtime, _, err := StatTimes(srcfile)
	if err != nil {
		log.Println(err)
		return err
	}

	err = os.Chtimes(dstfile, atime, mtime)
	if err != nil {
		log.Printf("filename: %s, change mtime to [%s] failed!  %v\n", dstfile, mtime, err)
	} else {
		log.Printf("filename: %s, change mtime to [%s] successfully!\n", dstfile, mtime)
	}
	return err
}

func ReadAll(filename string) string {
	fi, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	defer fi.Close()
	fd, _ := ioutil.ReadAll(fi)
	// fmt.Println(string(fd))
	return string(fd)
}

func StatTimes(name string) (atime, mtime, ctime time.Time, err error) {
	fi, err := os.Stat(name)
	if err != nil {
		return
	}
	mtime = fi.ModTime()
	stat := fi.Sys().(*syscall.Stat_t)
	atime = time.Unix(int64(stat.Atim.Sec), int64(stat.Atim.Nsec))
	ctime = time.Unix(int64(stat.Ctim.Sec), int64(stat.Ctim.Nsec))
	return
}

/*
func MTime(name string) (mtime string) {
	fi, err := os.Stat(name)
	if err != nil {
		return
	}
	mtime = fi.ModTime().String()
	// log.Println(mtime)
	return
}
*/
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

// func GetPWD() (string, error) {
// 	file, err := exec.LookPath(os.Args[0])
// 	if err != nil {
// 		return "", err
// 	}
// 	path, err := filepath.Abs(file)
// 	if err != nil {
// 		return "", err
// 	}
// 	i := strings.LastIndex(path, "/")
// 	if i < 0 {
// 		i = strings.LastIndex(path, "\\")
// 	}
// 	if i < 0 {
// 		return "", errors.New(`error: Can't find "/" or "\".`)
// 	}
// 	return string(path[0 : i+1]), nil
// }
