package tar

import (
	"archive/tar"
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"path"
)

func TarFiles(fromFile []string, toFile string) {
	// 创建一个缓冲区来写入我们的存档。
	buf := new(bytes.Buffer)
	// 创建一个新的tar存档。
	tw := tar.NewWriter(buf)

	for _, file := range fromFile {
		fi, err := os.Stat(file)
		if err != nil {
			continue
		}
		hdr, err := tar.FileInfoHeader(fi, "")
		if err != nil {
			continue
		}

		if err := tw.WriteHeader(hdr); err != nil {
			log.Fatalln(err)
			continue
		}

		dat, err := ioutil.ReadFile(file)
		if _, err := tw.Write(dat); err != nil {
			log.Fatalln(err)
			continue
		}
	}
	// 确保在Close时检查错误。
	if err := tw.Close(); err != nil {
		log.Fatalln(err)
	}

	ioutil.WriteFile(toFile, buf.Bytes(), 0644)
}

func UnTarFiles(tarFile string, untarPath string) error {
	//打开要解包的文件，tarFile是要解包的 .tar 文件的路径
	fr, er := os.Open(tarFile)
	if er != nil {
		return er
	}
	defer fr.Close()
	// 创建 tar.Reader，准备执行解包操作
	tr := tar.NewReader(fr)
	//用 tr.Next() 来遍历包中的文件，然后将文件的数据保存到磁盘中
	for hdr, er := tr.Next(); er != io.EOF; hdr, er = tr.Next() {
		if er != nil {
			return er
		}
		//先创建目录
		fileName := untarPath + "/" + hdr.Name
		dir := path.Dir(fileName)
		_, err := os.Stat(dir)
		//如果err 为空说明文件夹已经存在，就不用创建
		if err != nil {
			err = os.MkdirAll(dir, os.ModePerm)
			if err != nil {
				fmt.Print(err)
				return err
			}
		}
		//获取文件信息
		fi := hdr.FileInfo()
		//创建空文件，准备写入解压后的数据
		fw, er := os.Create(fileName)
		if er != nil {
			return er
		}
		defer fw.Close()
		// 写入解压后的数据
		_, er = io.Copy(fw, tr)
		if er != nil {
			return er
		}
		// 设置文件权限
		os.Chmod(fileName, fi.Mode().Perm())
	}
	return nil
}
