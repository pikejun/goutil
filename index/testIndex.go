// Copyright 2019 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// +build ignore

// Gen generates sais2.go by duplicating functions in sais.go
// using different input types.
// See the comment at the top of sais.go for details.
package main

import (
	"fmt"
	"index/suffixarray"
)
func init() {

}
func main() {
	index := suffixarray.New([]byte("banana"))
	offsets := index.Lookup([]byte("ana"), -1)
	for _, off := range offsets {
		fmt.Println(off)
	}
}

