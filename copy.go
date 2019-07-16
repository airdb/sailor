// package main
package sailor

import (
	"fmt"
	"io"
	"log"
	"os"
)

/*
1. Perform checks to ensure that at least some form of copy will succeed (access permissions, directories exist, etc.)
2. Check to see if both files already exist and are the same using os.SameFile, return success if they are the same
3. Attempt a Link, return if success
4. Copy the bytes (all efficient means failed), return result
*/

// CopyFile copies a file from src to dst. If src and dst files exist, and are
// the same, then return success. Otherise, attempt to create a hard link
// between the two files. If that fail, copy the file contents from src to dst.
func CopyFile(src, dst string) (err error) {
	sfi, err := os.Stat(src)
	if err != nil {
		return
	}
	if !sfi.Mode().IsRegular() {
		// cannot copy non-regular files (e.g., directories,
		// symlinks, devices, etc.)
		return fmt.Errorf("CopyFile: non-regular source file %s (%q)", sfi.Name(), sfi.Mode().String())
	}

	// create dst path if not exist.
	pdir, _ := GetFilePath(dst)
	if "" != pdir {
		if _, err := os.Stat(pdir); os.IsNotExist(err) {
			err = os.MkdirAll(pdir, 0755)
			if err != nil {
				log.Printf("create dst path fail, path: %s, Mode 0755.\n", pdir)
			}
		}
	}

	dfi, err := os.Stat(dst)
	if err != nil {
		if !os.IsNotExist(err) {
			return
		}
	} else {
		if !(dfi.Mode().IsRegular()) {
			return fmt.Errorf("CopyFile: non-regular destination file %s (%q)", dfi.Name(), dfi.Mode().String())
		}
		if os.SameFile(sfi, dfi) {
			return
		}
	}

	// create link
	// if err = os.Link(src, dst); err == nil {
	// log.Println( "====", src, dst)
	//     return
	// }
	err = copyFileContents(src, dst)

	if err != nil {
		log.Printf("CopyFile failed %q, filename: %s\n", err, dst)
		panic(err)
	}
	// else {
	// 	log.Printf("CopyFile succeeded, filename: %s\n", dst)
	// }

	// copy file atime and mtime.
	err = CopyFileTime(src, dst)
	if err != nil {
		log.Println(err)
	}

	srcMd5 := ComputeMd5(src)
	dstMd5 := ComputeMd5(dst)
	if srcMd5 == dstMd5 {
		log.Printf("CopyFile succeeded, filename: %s, dstMd5: %s\n", dst, dstMd5)
	} else {
		log.Printf("CopyFile failed %q, filename: %s, srcMd5: %s, dstMd5: %s\n", err, dst, srcMd5, dstMd5)
	}

	return
}

// copyFileContents copies the contents of the file named src to the file named
// by dst. The file will be created if it does not already exist. If the
// destination file exists, all it's contents will be replaced by the contents
// of the source file.
func copyFileContents(src, dst string) (err error) {
	log.Printf("copy content from %s to %s\n", src, dst)
	in, err := os.Open(src)
	if err != nil {
		return
	}
	defer in.Close()
	out, err := os.Create(dst)
	if err != nil {
		return
	}
	defer func() {
		cerr := out.Close()
		if err == nil {
			err = cerr
		}
	}()
	if _, err = io.Copy(out, in); err != nil {
		return
	}
	err = out.Sync()
	return
}
