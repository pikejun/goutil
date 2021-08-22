package tar_test

import (
	"../tar"
	"fmt"
	"testing"
)

func Test1_aaaa(t *testing.T) {
	fmt.Println("123")
	tar.TarFiles([]string{"D:/400/1_maizhi2.txt", "D:/400/1_maizhi3.txt", "D:/400/2_maizhi1.txt"}, "D:/400/ddd.tar")
}

func TestUnTarFiles(t *testing.T) {
	tar.UnTarFiles("D:/400/ddd.tar", "D:/400/ddd3")
}
