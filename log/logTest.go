package main

import (
	"bytes"
	"fmt"
	"log"
	"os"
)

func main() {

	//1.日志写入一段缓存里
	var buf bytes.Buffer
	logger := log.New(&buf, "logger: ", log.Llongfile)//三个参数对应输出对象，前缀，文件名然后整合成前缀
	logger.Print("Hello, log file!") //实际log信息
	fmt.Println(&buf)

	//2.日志写入一段文件里,执行完刷新工程，会有文件debug.log打开来看
	fileName := "debug.log"  //在工程路径下和src同级，也可以写绝对路径，不过要注意转义符
	logFile,err  := os.Create(fileName)  //创建该文件，返回句柄
	defer logFile.Close()      //确保文件在该函数执行完以后关闭
	if err != nil {
		log.Fatalln("open file error !")
	}

	//库里的log包不存在级别，只有Print，Fatal，Panic三种，想要7种级别，可以用syslog包或者框架
	debugLog := log.New(logFile,"",log.Llongfile)
	debugLog.Print("找到一个低级bug,低级别的日志报告，执行debug记录后继续执行")
	debugLog.Fatal("发现一个严重bug，日志已记录，程序在此终止")
	debugLog.Panic("哎呀，我没有被执行，好惨")

	//3.使用beego框架的log包进行log处理 ，对应7种级别

}