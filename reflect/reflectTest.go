package main

import (
	"reflect"

	"fmt"
)

type User struct {
	Name  string
	Age   int
	Sex   bool
	Phone *string
	Qian  float64
	Atest uint
	Group interface{}
	Btest interface{}
}

func (u *User) Hello() {
	fmt.Println("hello world 你好世界")
}

func main() {

	a := "hello world 你好世界"
	user := &User{"张三", 25, true, &a, 88.8, 9, 99, nil}

	var obj interface{} = user
	v := reflect.ValueOf(obj)

	method := v.MethodByName("Hello") //返回v的名为Hello的方法
	method.Call([]reflect.Value{})    //执行反射的方法

	fmt.Println(v.IsValid()) //返回v是否持有值，如果v是value零值会返回假，此时v除了IsValid String Kind之外的方法都会导致panic
	fmt.Println(v.Kind())    //返回v持有值的分类，如果v是value零值，返回值为invalid
	fmt.Println(v.Type())    //返回v持有值的类型Type表示

	v = v.Elem() //返回持有的接口的值，或者指针的值，如果不是interface{}或指针会panic,实际上是从 *User到User
	var u User
	fmt.Println(v.Convert(reflect.TypeOf(u)).FieldByName("Name")) //转换为其他类型的值,如果无法使用标准Go转换规则来转换，那么panic

	fmt.Println(v.FieldByName("Name").CanSet())   //是否可以设置Name的值
	v.FieldByName("Name").SetString("把Name值修改一下") //设置v的持有值，如果v的kind不是string或者v.Canset()返回假，会panic
	v.FieldByName("Name").Set(reflect.ValueOf(a)) //将v的持有值修改为a的反射值，如果Canset返回假，会panic

	fmt.Println(v.FieldByName("Group").Elem())     //返回持有的接口的值，或者指针的值，如果不是interface{}或指针会panic
	fmt.Println(v.FieldByName("Phone").Elem())     //或者指针的值
	fmt.Println(v.FieldByName("Name").Interface()) //把Name当做interface{}值

	fmt.Println(v.FieldByName("Name").String()) //返回v持有的值的字符串表示，如果v的值不是string也不会panic
	fmt.Println(v.FieldByName("Sex").Bool())    //返回持有的布尔值，如果v的kind不是bool会panic
	fmt.Println(v.FieldByName("Age").Int())     //返回持有的int64，如果v的kind不是int int8-int64会panic

	var x int64
	fmt.Println(v.FieldByName("Age").OverflowInt(x)) //如果v持有值的类型不能无一出的表示x，会返回真，如果v的kind不是int int8-int64会panic
	fmt.Println(v.FieldByName("Atest").Uint())       //返回v持有的无符号整数,如果v的kind不是uint uintptr uint8 uint16 uint32 uint64会panic

	var x2 uint64
	fmt.Println(v.FieldByName("Atest").OverflowUint(x2)) //如果v持有的值的类型不能无溢出的表示x2，会返回真，如果v的kind不是uint uintptr uint8 uint16 uint32 uint64会panic
	fmt.Println(v.FieldByName("Qian").Float())           //返回v持有的浮点数float64,如果v的kind不是float32 float64会panic

	var x3 float64
	fmt.Println(v.FieldByName("Qian").OverflowFloat(x3)) //如果v持有值的类型不能无溢出的表示x3，会返回真，如果v的kind不是float32 float64会panic
	fmt.Println(v.FieldByName("Btest").IsNil())          //如果v持有值是否为nil，如果v的值不是通道 函数 接口 映射 指针 切片之一会panic

	fmt.Println(v.NumField())             //返回v持有的结构体类型值的字段数，如果v的kind不是struct会panic
	fmt.Println(v.Field(0))               //返回结构体的第i个字段，如果v的kind不是struct或i出界会panic
	fmt.Println(v.FieldByIndex([]int{0})) //和上面一样，没明白有啥用

}