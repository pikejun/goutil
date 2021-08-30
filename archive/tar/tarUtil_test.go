package tar

import (
	"fmt"
	"testing"
)

func Test1_aaaa(t *testing.T) {
	fmt.Println("123")
	TarFiles([]string{"D:/400/1_maizhi2.txt", "D:/400/1_maizhi3.txt", "D:/400/2_maizhi1.txt"}, "D:/400/ddd.tar")
}

func TestUnTarFiles(t *testing.T) {
	UnTarFiles("D:/400/ddd.tar", "D:/400/ddd3")
}
