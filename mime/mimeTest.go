package main

import(
	"fmt"
	"mime"
)


func main() {
	mineType1 := mime.TypeByExtension(".svg")
	fmt.Println(mineType1)  //image/svg+xml
	mineType2 := mime.TypeByExtension(".svv") //一开始是没有与该扩展名相关的mineType
	fmt.Println(mineType2) //为空
	err := mime.AddExtensionType(".svv", "mytype/none")//在这里添加后在查找就能够查找到了
	if err != nil{
		fmt.Println(err)
	}
	mineType3 := mime.TypeByExtension(".svv")
	fmt.Println(mineType3) //mytype/none
}