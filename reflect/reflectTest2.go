package main

import (
	"reflect"

	"fmt"
)

type lx interface {
	SayHi()
}

type User struct {
	Name string
	Age  int64
	Sex  string
}

func (u *User) SayHi() {
	fmt.Println("hello world")
}

func main() {
	user := User{"张三", 25, "男"}
	FillStruct(user)
}

func FillStruct(obj interface{}) {
	t := reflect.TypeOf(obj)       //反射出一个interface{}的类型
	fmt.Println("11111111",reflect.ValueOf(obj).Field(0).String())
	fmt.Println(t.Name())          //类型名
	fmt.Println(t.Kind().String()) //Type类型表示的具体分类
	fmt.Println(t.PkgPath())       //反射对象所在的短包名
	fmt.Println(t.String())        //包名.类型名
	fmt.Println(t.Size())          //要保存一个该类型要多少个字节
	fmt.Println(t.Align())         //返回当从内存中申请一个该类型值时，会对齐的字节数
	fmt.Println(t.FieldAlign())    //返回当该类型作为结构体的字段时，会对齐的字节数

	var u User
	fmt.Println(t.AssignableTo(reflect.TypeOf(u)))  // 如果该类型的值可以直接赋值给u代表的类型，返回真
	fmt.Println(t.ConvertibleTo(reflect.TypeOf(u))) // 如该类型的值可以转换为u代表的类型，返回真

	fmt.Println(t.NumField())             // 返回struct类型的字段数（匿名字段算作一个字段），如非结构体类型将panic
	fmt.Println(t.Field(0).Name)          // 返回struct类型的第i个字段的类型，如非结构体或者i不在[0, NumField())内将会panic
	fmt.Println(t.FieldByName("Age"))     // 返回该类型名为name的字段（会查找匿名字段及其子字段），布尔值说明是否找到，如非结构体将panic
	fmt.Println(t.FieldByIndex([]int{0})) // 返回索引序列指定的嵌套字段的类型，等价于用索引中每个值链式调用本方法，如非结构体将会panic
}