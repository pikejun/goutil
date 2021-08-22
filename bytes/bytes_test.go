package bytes

import (
	"bytes"
	"fmt"
	"testing"
)

func Test_replace(t *testing.T) {
	a := "ok a ok b ok c"

	d := bytes.IndexByte([]byte(a), byte('c'))

	fmt.Println(d)
}
