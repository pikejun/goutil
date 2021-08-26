package main

import (
	"crypto/ed25519"
	"crypto/rand"
	"encoding/pem"
	"fmt"
	"os"
)

func GenerateEd25519Key() {
	pubKey, privKey, err := ed25519.GenerateKey(rand.Reader)
	if err != nil {
		panic(err)
	}
	//Step3:组织一个pem的block结构体
	block := pem.Block{
		Type:  "ED25519 Private Key",
		Bytes: privKey,
	}
	//Step4:进行pem编码
	file, err := os.Create("ed25519Private.pem")
	if err != nil {
		panic(err)
	}
	defer file.Close()
	pem.Encode(file, &block)
	//----------获取公匙并写入磁盘----------
	//Step1:获取公匙
	//Step2:采用x509序列化
	//Step3:组织一个pem的block结构体
	block = pem.Block{
		Type:  "ED25519 Public Key",
		Bytes: pubKey,
	}
	//Step4:进行pem编码
	file, err = os.Create("ed25519Public.pem")
	if err != nil {
		panic(err)
	}
	defer file.Close()
	pem.Encode(file, &block)
}

//------------ECDSA私匙执行数字签名------------
func Ed25519Signature(plainText []byte, privName string) (r []byte) {
	//------1.获取私匙------
	//Step1:打开文件获取原始私匙
	file, err := os.Open(privName)
	if err != nil {
		panic(err)
	}
	defer file.Close()
	fileinfo, err := file.Stat()
	if err != nil {
		panic(err)
	}
	buf := make([]byte, fileinfo.Size())
	file.Read(buf)

	//Step2:私匙的反pem编码化
	block, _ := pem.Decode(buf)

	r = ed25519.Sign(block.Bytes, plainText)
	return
}

//------------ECDSA公匙验证数字签名------------
func Ed25519Verify(plainText, sign []byte, pubFile string) bool {
	//------1.获取公钥------
	//Step1:打开文件获取公匙
	file, err := os.Open(pubFile)
	if err != nil {
		panic(err)
	}
	defer file.Close()
	fileinfo, err := file.Stat()
	if err != nil {
		panic(err)
	}
	buf := make([]byte, fileinfo.Size())
	file.Read(buf)
	//Step2:私匙的反pem编码化
	block, _ := pem.Decode(buf)
	return ed25519.Verify(block.Bytes, plainText, sign)
}

//主函数
func main() {
	GenerateEd25519Key()
	src := []byte("AAAAAAAAAAAAAAAAAAAABlock中的Bytes变量中的数据进行解析 ")
	src2 := []byte("AAAAAAAAAAAAAAAAAAAABlock中的Bytes变量中的数据进行解析 ")
	sText := Ed25519Signature(src, "ed25519Private.pem")
	fmt.Println("sText=", sText)
	bl := Ed25519Verify(src2, sText, "ed25519Public.pem")
	fmt.Println(bl)
}
