package main

import (
	"context"
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"github.com/golang/protobuf/ptypes/timestamp"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"grpc_demo/services"
	"io/ioutil"
	"time"
)

func main() {
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
	//----
	client := services.NewProductClient(conn)
	res, _ := client.ProductInfo(context.Background(), &services.QuerySize{Size: 10})
	fmt.Println(res)
	//----
	res1, _ := client.SomeProduct(context.Background(), &services.QuerySize{Size: 10})
	fmt.Println(res1)
	//----
	orderClient := services.NewOrderClient(conn)
	time := timestamp.Timestamp{Seconds: time.Now().Unix()}
	response, _ := orderClient.NewOrder(context.Background(), &services.OrderRequest{
		Order:&services.ProductInfo{
			Pid:   10,
			Pname: "kk",
			Price: 99.10,
			Time:  &time,
		},
	})
	fmt.Println(response)

}
