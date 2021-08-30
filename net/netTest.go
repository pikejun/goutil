package main

import (
	"bytes"
	"errors"
	"fmt"
	"net"
	"time"
)

func main() {
	Ping("47.93.33.216")
}

func Ping(ip string) (bool, error) {
	recv := make([]byte, 1024)                //保存响应数据
	raddr, err := net.ResolveIPAddr("ip", ip) //raddr为目标主机的地址
	if err != nil {
		fmt.Printf("resolve ip: %s fail:", ip)
		return false, err
	}
	laddr := net.IPAddr{IP: net.ParseIP("0.0.0.0")} //源地址
	if ip == "" {
		return false, errors.New("ip or domain is null")
	}

	conn, err := net.DialIP("ip4:icmp", &laddr, raddr)
	if err != nil {
		return false, err
	}
	defer conn.Close()

	buffer := assemblyIcmp()
	conn.Write(buffer.Bytes())

	conn.SetReadDeadline((time.Now().Add(time.Second * 5)))
	n, _ := conn.Read(recv)

	fmt.Println(recv[:n])

	return true, nil
}

func assemblyIcmp() *bytes.Buffer{
	b:=bytes.NewBufferString("ping -t 47.93.33.216\n")
	return b
}