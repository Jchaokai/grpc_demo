package main

import (
	"context"
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"grpc_demo/services"
	"io/ioutil"
)

func main() {
	//Credentials, e := credentials.NewClientTLSFromFile("keys/server.crt", "greet.com")
	//if e != nil {
	//	log.Fatal(e)
	//}
	cert, _ := tls.LoadX509KeyPair("crt/client.pem", "crt/client.key")
	certPool := x509.NewCertPool()
	ca, _ := ioutil.ReadFile("crt/ca.pem")
	certPool.AppendCertsFromPEM(ca)
	creds := credentials.NewTLS(&tls.Config{
		Certificates: []tls.Certificate{cert}, //服务端证书
		ServerName:   "localhost",
		RootCAs:      certPool,
	})

	conn, _ := grpc.Dial(":8081", grpc.WithTransportCredentials(creds))
	defer conn.Close()
	client := services.NewHelloClient(conn)
	resp, _ := client.Hello(context.Background(), &services.Request{Name:"jjb",Gender:"male"})
	fmt.Println(resp.Res)
}
