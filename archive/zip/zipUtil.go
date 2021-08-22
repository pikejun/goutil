package zip

import (
	"archive/zip"
	"io"
	"os"
	"path/filepath"
)

func ZipFiles(fromFile []string, toFile string) {
	// 创建一个缓冲区来写入我们的存档。
	zipfile, _ := os.Create(toFile)
	defer zipfile.Close()
	archive := zip.NewWriter(zipfile)
	defer archive.Close()

	for _, file := range fromFile {
		filepath.Walk(file, func(path string, info os.FileInfo, err error) error {
			header, err := zip.FileInfoHeader(info)

			file, err := os.Open(path)
			if err != nil {
				return err
			}

			if info.IsDir() {
				header.Name += "/"
			} else {
				header.Method = zip.Deflate
			}

			writer, err := archive.CreateHeader(header)

			defer file.Close()
			_, err = io.Copy(writer, file)
			return err
		})
	}
}

func Unzip(zipFile string, destDir string) error {
	zipReader, err := zip.OpenReader(zipFile)
	if err != nil {
		return err
	}
	defer zipReader.Close()

	for _, f := range zipReader.File {
		fpath := filepath.Join(destDir, f.Name)
		if f.FileInfo().IsDir() {
			os.MkdirAll(fpath, os.ModePerm)
		} else {
			if err = os.MkdirAll(filepath.Dir(fpath), os.ModePerm); err != nil {
				return err
			}

			inFile, err := f.Open()
			if err != nil {
				return err
			}
			defer inFile.Close()

			outFile, err := os.OpenFile(fpath, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, f.Mode())
			if err != nil {
				return err
			}
			defer outFile.Close()

			_, err = io.Copy(outFile, inFile)
			if err != nil {
				return err
			}
		}
	}
	return nil
}
