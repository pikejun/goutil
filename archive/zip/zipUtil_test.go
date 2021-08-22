package zip_test

import (
	"../zip"
	"fmt"
	"testing"
)

func Test_zip(test *testing.T) {
	fmt.Println("123")
	zip.ZipFiles([]string{"D:/400/1_maizhi2.txt", "D:/400/1_maizhi3.txt", "D:/400/2_maizhi1.txt"}, "D:/400/ddd.zip")
}

func Test_unzip(test *testing.T) {
	zip.Unzip("D:/400/ddd.zip", "D:/400/ddd2")
}
