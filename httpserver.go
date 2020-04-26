package main

import (
	"context"
	"crypto/tls"
	"crypto/x509"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"grpc_demo/services"
	"io/ioutil"
	"net/http"
)

func main() {

	//双向验证
	cert, _ := tls.LoadX509KeyPair("crt/client.pem", "crt/client.key")
	certPool := x509.NewCertPool()
	ca, _ := ioutil.ReadFile("crt/ca.pem")
	certPool.AppendCertsFromPEM(ca)
	creds := credentials.NewTLS(&tls.Config{
		Certificates: []tls.Certificate{cert}, //服务端证书
		ServerName:   "localhost",
		RootCAs:      certPool,
	})
	serveMux := runtime.NewServeMux()
	options := []grpc.DialOption{grpc.WithTransportCredentials(creds)}
	_ = services.RegisterHelloHandlerFromEndpoint(context.Background(), serveMux, "localhost:8081", options)
	_ = services.RegisterProductHandlerFromEndpoint(context.Background(), serveMux, "localhost:8081", options)
	_ = services.RegisterOrderHandlerFromEndpoint(context.Background(), serveMux, "localhost:8081", options)
	httpServer := &http.Server{
		Addr:    ":8080",
		Handler: serveMux,
	}
	_ = httpServer.ListenAndServe()

}