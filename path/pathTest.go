// Copyright 2012 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"fmt"
	"os"
	"path"
	"path/filepath"
)

func main(){
	//Path操作
	fmt.Println("Path操作-----------------")
	fmt.Println(path.Base("http://www.baidu.com/file/aa.jpg")) //aa.jpg
	fmt.Println(path.Clean("c:\\file//abc///aa.jpg"))          //c:\file/abc/aa.jpg
	fmt.Println(os.Getwd())                                    //D:\Projects\GoPath\source\demo\syntax\path <nil>
	fmt.Println(path.Dir("http://www.baidu.com/aa/aaa.jpg"))   //http:/www.baidu.com/aa
	fmt.Println(path.Dir("c:/a/b/c/d.txt"))                    //c:/a/b/c
	fmt.Println(path.Dir("c:\\a/b.txt"))                       //c:\a
	fmt.Println(path.Ext("c:\\a/b.txt"))                       //.txt
	fmt.Println(path.IsAbs("c:/wind/aa/bb/b.txt"))             //false
	fmt.Println(path.Join("c:", "aa", "bb", "cc.txt"))         //c:/aa/bb/cc.txt
	isMatch, err := path.Match("c:/windows/*/", "c:/windows/system/")
	fmt.Println(isMatch, err)                            //true <nil>
	fmt.Println(path.Split("c:/windows/system/aaa.jpg")) //c:/windows/system/ aaa.jpg
	//FilePath操作
	fmt.Println("FilePath操作-----------------")
	fmt.Println(filepath.IsAbs("c:\\wind\\aa\\bb\\b.txt"))                 //true
	fmt.Println(filepath.Abs("."))                                         //D:\Projects\GoPath\source\demo\syntax\path <nil>
	fmt.Println(filepath.Base("c:\\aa\\baa.exe"))                          //baa.exe
	fmt.Println(filepath.Clean("c:\\\\aa/c\\baa.exe"))                     //c:\aa\c\baa.exe
	fmt.Println(filepath.Clean("aa/c\\baa.exe"))                           //aa\c\baa.exe
	fmt.Println(filepath.Dir("aa/c\\baa.exe"))                             //aa\c
	fmt.Println(filepath.EvalSymlinks("./path.exe"))                       //可以用来判断文件或文件夹是否存在。 //path.exe <nil>
	fmt.Println(filepath.Ext("./path.exe"))                                //.exe
	fmt.Println(filepath.FromSlash("c:\\windows\\aa//bb/cc//path.exe"))    //将路径中的\\更换为/  //c:\windows\aa\\bb\cc\\path.exe
	fmt.Println(filepath.ToSlash("c:\\windows\\aa/bb/cc/path.exe"))        //将路径中的/替换为\\   //c:/windows/aa/bb/cc/path.exe
	fmt.Println(filepath.VolumeName("c:\\windows\\"))                      //获取卷标   //c:
	fmt.Println(filepath.Glob("c:\\windows\\*.exe"))                       //获取所有c:\\windows\\目录下exe文件。
	fmt.Println(filepath.HasPrefix("c:\\aa\\bb", "c:\\"))                  //true
	fmt.Println(filepath.IsAbs("http://www.baidu.com/aa.jpg"))             //false
	fmt.Println(filepath.Join("a", "\\bb\\", "cc", "/d", "e\\", "ff.txt")) //a\bb\cc\d\e\ff.txt
	fmt.Println(filepath.Match("c:/windows/*/", "c:/windows/system/"))     //true <nil>
	fmt.Println(filepath.Rel("c:/windows", "c:/windows/system/"))          //取得第二参的路径中，相对于前面的路径的相对路径。  //system <nil>
	fmt.Println(string(filepath.Separator))                                // windows下返回\\
	fmt.Println(filepath.Split("c:/windows/system/abc.exe"))               //c:/windows/system/ abc.exe
	fmt.Println(filepath.SplitList("c:/windows/system/abc.exe"))           //[c:/windows/system/abc.exe]
	filepath.Walk(".", WalkFunc)
	/*
	   File: ../../syntax IsDir: true size: 0
	   File: ..\..\syntax\painc IsDir: true size: 0
	   File: ..\..\syntax\painc\main.go IsDir: false size: 813
	   File: ..\..\syntax\painc\painc.exe IsDir: false size: 2498048
	   File: ..\..\syntax\path IsDir: true size: 0
	   File: ..\..\syntax\path\path.exe IsDir: false size: 2851328
	   File: ..\..\syntax\path\path.go IsDir: false size: 3419
	*/

}
func WalkFunc(path string, info os.FileInfo, err error) error {
	fmt.Println("File:", path, "IsDir:", info.IsDir(), "size:", info.Size())
	return nil
}
func ExampleBase() {
	fmt.Println(path.Base("/a/b"))
	fmt.Println(path.Base("/"))
	fmt.Println(path.Base(""))
	// Output:
	// b
	// /
	// .
}

func ExampleClean() {
	paths := []string{
		"a/c",
		"a//c",
		"a/c/.",
		"a/c/b/..",
		"/../a/c",
		"/../a/b/../././/c",
		"",
	}

	for _, p := range paths {
		fmt.Printf("Clean(%q) = %q\n", p, path.Clean(p))
	}

	// Output:
	// Clean("a/c") = "a/c"
	// Clean("a//c") = "a/c"
	// Clean("a/c/.") = "a/c"
	// Clean("a/c/b/..") = "a/c"
	// Clean("/../a/c") = "/a/c"
	// Clean("/../a/b/../././/c") = "/a/c"
	// Clean("") = "."
}

func ExampleDir() {
	fmt.Println(path.Dir("/a/b/c"))
	fmt.Println(path.Dir("a/b/c"))
	fmt.Println(path.Dir("/a/"))
	fmt.Println(path.Dir("a/"))
	fmt.Println(path.Dir("/"))
	fmt.Println(path.Dir(""))
	// Output:
	// /a/b
	// a/b
	// /a
	// a
	// /
	// .
}

func ExampleExt() {
	fmt.Println(path.Ext("/a/b/c/bar.css"))
	fmt.Println(path.Ext("/"))
	fmt.Println(path.Ext(""))
	// Output:
	// .css
	//
	//
}

func ExampleIsAbs() {
	fmt.Println(path.IsAbs("./dev/null"))
	// Output: true
}

func ExampleJoin() {
	fmt.Println(path.Join("a", "b", "c"))
	fmt.Println(path.Join("a", "b/c"))
	fmt.Println(path.Join("a/b", "c"))

	fmt.Println(path.Join("a/b", "../../../xyz"))

	fmt.Println(path.Join("", ""))
	fmt.Println(path.Join("a", ""))
	fmt.Println(path.Join("", "a"))

	// Output:
	// a/b/c
	// a/b/c
	// a/b/c
	// ../xyz
	//
	// a
	// a
}

func ExampleMatch() {
	fmt.Println(path.Match("abc", "abc"))
	fmt.Println(path.Match("a*", "abc"))
	fmt.Println(path.Match("a*/b", "a/c/b"))
	// Output:
	// true <nil>
	// true <nil>
	// false <nil>
}

func ExampleSplit() {
	split := func(s string) {
		dir, file := path.Split(s)
		fmt.Printf("path.Split(%q) = dir: %q, file: %q\n", s, dir, file)
	}
	split("static/myfile.css")
	split("myfile.css")
	split("")
	// Output:
	// path.Split("static/myfile.css") = dir: "static/", file: "myfile.css"
	// path.Split("myfile.css") = dir: "", file: "myfile.css"
	// path.Split("") = dir: "", file: ""
}
