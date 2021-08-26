package main

import (
	"crypto/tls"
	"crypto/x509"
	"io/ioutil"
	"log"
)

func main() {

	cert, err := tls.LoadX509KeyPair("./crypto/tls/client.pem", "./crypto/tls/client.key")
	if err != nil {
		log.Println(err)
		return
	}
	certBytes, err := ioutil.ReadFile("./crypto/tls/client.pem")

	clientCertPool := x509.NewCertPool()
	clientCertPool.AppendCertsFromPEM(certBytes)

	conf := &tls.Config{
		RootCAs:            clientCertPool,
		Certificates:       []tls.Certificate{cert},
		InsecureSkipVerify: true,
	}
	conn, err := tls.Dial("tcp", "127.0.0.1:443", conf)
	if err != nil {
		log.Println(err)
		return
	}
	defer conn.Close()
	n, err := conn.Write([]byte("hello\n"))
	if err != nil {
		log.Println(n, err)
		return
	}
	buf := make([]byte, 100)
	n, err = conn.Read(buf)
	if err != nil {
		log.Println(n, err)
		return
	}
	println(string(buf[:n]))
}
