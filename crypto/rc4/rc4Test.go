package main

import (
	"crypto/rc4"
	"fmt"
)

func main() {
	// 密钥
	key := []byte("abcdefg")
	// 要加密的源数据
	str := []byte("this is my test!")

	// 加密方式1：加密/解密后的数据单独存放
	{
		// 加密操作
		dest1 := make([]byte, len(str))
		fmt.Printf("方法1加密前:%s \n", str)
		cipher1, _ := rc4.NewCipher(key)
		cipher1.XORKeyStream(dest1, str)
		fmt.Printf("方法1加密后:%s \n", dest1)

		// 解密操作
		dest2 := make([]byte, len(dest1))
		cipher2, _ := rc4.NewCipher(key) // 切记：这里不能重用cipher1，必须重新生成新的
		cipher2.XORKeyStream(dest2, dest1)
		fmt.Printf("方法1解密后:%s \n\n", dest2)
	}

	// 加密方式2：加密后的数据直接存放在源数据那里，不需额外申请空间
	{
		// 加密操作
		fmt.Printf("方法2加密前:%s \n", str)
		cipher1, _ := rc4.NewCipher(key)
		cipher1.XORKeyStream(str, str) // 加密后的数据直接覆盖到str中
		fmt.Printf("方法2加密后:%s \n", str)

		// 解密操作
		cipher2, _ := rc4.NewCipher(key)
		cipher2.XORKeyStream(str, str) // 解密后的数据直接覆盖到str中
		fmt.Printf("方法2解密后:%s \n\n", str)
	}
}
