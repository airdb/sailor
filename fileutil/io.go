package fileutil

//nolint:gosec
import (
	"archive/tar"
	"compress/gzip"
	"crypto/md5"
	"encoding/hex"
	"errors"
	"io"
	"log"
	"os"
)

func IOWriteFile(r io.Reader, filePath string) error {
	outFile, err := os.Create(filePath)
	if err != nil {
		return err
	}

	defer outFile.Close()

	_, err = io.Copy(outFile, r)

	return err
}

func GetMd5Sum(f io.Reader) string {
	h := md5.New() //nolint:gosec
	if _, err := io.Copy(h, f); err != nil {
		log.Fatal(err)

		return ""
	}

	return hex.EncodeToString(h.Sum(nil))
}

func ExtraTarFile(r io.Reader) {
	uncompressedStream, err := gzip.NewReader(r)
	if err != nil {
		log.Fatal("ExtractTarGz: NewReader failed")
	}

	tarReader := tar.NewReader(uncompressedStream)

	for {
		header, err := tarReader.Next()

		if errors.Is(err, io.EOF) {
			break
		}

		if err != nil {
			log.Fatalf("ExtractTarGz: Next() failed: %s", err.Error())
		}

		switch header.Typeflag {
		case tar.TypeDir:
			EnsureFolderExists(header.Name)
		case tar.TypeReg:
			WriteTarReaderToFile(header.Name, tarReader)
		default:
			log.Fatalf(
				"ExtractTarGz: uknown type: %v in %s",
				header.Typeflag,
				header.Name)
		}
	}
}
