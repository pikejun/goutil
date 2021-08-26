package main

import (
	"crypto/dsa"
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"math/big"
	"strings"
)

func getNewKeyPair() *dsa.PrivateKey {
	// parameters 是私钥的参数
	var param dsa.Parameters
	// L1024N160是一个枚举，根据L1024N160来决定私钥的长度（L N）
	dsa.GenerateParameters(&param, rand.Reader, dsa.L1024N160)
	// 定义私钥的变量
	var privateKey dsa.PrivateKey
	// 设置私钥的参数
	privateKey.Parameters = param
	// 生成密钥对
	dsa.GenerateKey(&privateKey, rand.Reader)

	return &privateKey
}

func Sign(message string, privateKey *dsa.PrivateKey) string {
	//privateKey:=getNewKeyPair()
	//公钥是存在在私钥中的，从私钥中读取公钥
	//publicKey := privateKey.PublicKey
	str := []byte(message)
	//进入签名操作
	r, s, _ := dsa.Sign(rand.Reader, privateKey, str)
	return base64.StdEncoding.EncodeToString(r.Bytes()) + "@#@" + base64.StdEncoding.EncodeToString(s.Bytes())
}

func Verify(message string, sign string, publicKey *dsa.PublicKey) bool {

	s2 := strings.Split(sign, "@#@")
	if len(s2) != 2 {
		return false
	}

	r := new(big.Int)
	s := new(big.Int)
	b, _ := base64.StdEncoding.DecodeString(s2[0])
	r.SetBytes(b)
	c, _ := base64.StdEncoding.DecodeString(s2[1])
	s.SetBytes(c)

	// 进入验证
	flag := dsa.Verify(publicKey, []byte(message), r, s)

	return flag
}

/*
JN3fg5lzUKJLEWKPNJvSa6fSq50=
MMNI8PqRiIpGChuBP9ZsI2C1adnS2s00xRYlxAMMhOyHfU75CBG0UHBIc1TN08RaAhByQQS67GdnkMWys5SBzwWvZpg3sMJIXa0FNg4DFalZf0ZAnkWqCs+7DFJafAbhyVN6H59blmw5Hhg3JgvWdLOoTbmkA/9Lw6Q6nIbhKmw=
lsIVvv9DkmU/GI3PDrBo2RYCLzjOdl6RS9x/UGrQajfi4iY6hEQJG0cUMMlZiiT7SQ7gr7hvPH6R1fYDXGpjMlfpk66AqBBaAuUjxu2oLkk1QCjF9TW8BaTa1HFtfNthZ5StM5E7DxFknF0ufsx3aTKJ3X0mvBdaqm6XMehxdoc=
iRvoOEhNZ6b6h2DbY7COFzlQhDs=
FwnJ2hpr5OMiSHu4+gDn4urHZiBj4T/sJ2Aew56qZADjnVav3XmGhi4x85Clm7nPfJqnukKgegcAbbDUOBulTYiPQfkeE13WX4G+KlTZ9zi/UjFf4lphEQN2eQ5eGpF9wDMpFxNhVXaxS7c34ajGXplR0l+LsHvqHJH31CVYE80=
*/
func main3() {
	kp := getNewKeyPair()
	fmt.Println(base64.StdEncoding.EncodeToString(kp.X.Bytes()))
	fmt.Println(base64.StdEncoding.EncodeToString(kp.PublicKey.Y.Bytes()))
	fmt.Println(base64.StdEncoding.EncodeToString(kp.PublicKey.P.Bytes()))
	fmt.Println(base64.StdEncoding.EncodeToString(kp.PublicKey.Q.Bytes()))
	fmt.Println(base64.StdEncoding.EncodeToString(kp.PublicKey.G.Bytes()))
}
func main() {
	message := "okkkkkkkkkkkkkkkkokkkkkkkkkkkkkkkkokkkkkkkkkkkkkkkkokkkkkkkkkkkkkkkkokkkkkkkkkkkkkkkk"
	p := new(dsa.PrivateKey)
	b, _ := base64.StdEncoding.DecodeString("JN3fg5lzUKJLEWKPNJvSa6fSq50=")
	p.X = new(big.Int).SetBytes(b)
	b, _ = base64.StdEncoding.DecodeString("MMNI8PqRiIpGChuBP9ZsI2C1adnS2s00xRYlxAMMhOyHfU75CBG0UHBIc1TN08RaAhByQQS67GdnkMWys5SBzwWvZpg3sMJIXa0FNg4DFalZf0ZAnkWqCs+7DFJafAbhyVN6H59blmw5Hhg3JgvWdLOoTbmkA/9Lw6Q6nIbhKmw=")
	p.Y = new(big.Int).SetBytes(b)
	b, _ = base64.StdEncoding.DecodeString("lsIVvv9DkmU/GI3PDrBo2RYCLzjOdl6RS9x/UGrQajfi4iY6hEQJG0cUMMlZiiT7SQ7gr7hvPH6R1fYDXGpjMlfpk66AqBBaAuUjxu2oLkk1QCjF9TW8BaTa1HFtfNthZ5StM5E7DxFknF0ufsx3aTKJ3X0mvBdaqm6XMehxdoc=")
	p.P = new(big.Int).SetBytes(b)
	b, _ = base64.StdEncoding.DecodeString("iRvoOEhNZ6b6h2DbY7COFzlQhDs=")
	p.Q = new(big.Int).SetBytes(b)
	b, _ = base64.StdEncoding.DecodeString("FwnJ2hpr5OMiSHu4+gDn4urHZiBj4T/sJ2Aew56qZADjnVav3XmGhi4x85Clm7nPfJqnukKgegcAbbDUOBulTYiPQfkeE13WX4G+KlTZ9zi/UjFf4lphEQN2eQ5eGpF9wDMpFxNhVXaxS7c34ajGXplR0l+LsHvqHJH31CVYE80=")
	p.G = new(big.Int).SetBytes(b)

	v := Sign(message, p)
	fmt.Println(v)

	vb := Verify("okkkkkkkkkkkkkkkkokkkkkkkkkkkkkkkkokkkkkkkkkkkkkkkkokkkkkkkkkkkkkkkkokkkkkkkkkkkkkkkk", v, &p.PublicKey)

	fmt.Println(vb)
}
