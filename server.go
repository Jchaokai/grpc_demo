package main

import (
	"crypto/tls"
	"crypto/x509"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"grpc_demo/services"
	"io/ioutil"
	"net"
)

func main() {
	//加载证书
	//Credentials, e := credentials.NewServerTLSFromFile("keys/server.crt", "keys/server_no_password.key")
	//if e != nil {
	//	log.Fatal(e)
	//}
	//server := grpc.NewServer(grpc.Creds(Credentials))

	//双向验证
	cert, _ := tls.LoadX509KeyPair("crt/server.pem", "crt/server.key")
	certPool := x509.NewCertPool()
	ca, _ := ioutil.ReadFile("crt/ca.pem")
	certPool.AppendCertsFromPEM(ca)
	creds := credentials.NewTLS(&tls.Config{
		Certificates: []tls.Certificate{cert},	//服务端证书
		ClientAuth:   tls.RequireAndVerifyClientCert, //双向验证
		ClientCAs:    certPool,
	})
	server := grpc.NewServer(grpc.Creds(creds))
	services.RegisterHelloServer(server,new(services.Hello_Service))
	services.RegisterProductServer(server,new(services.Product_Service))
	services.RegisterOrderServer(server,new(services.Order_Service))
	listener, _ := net.Listen("tcp", ":8081")
	_ = server.Serve(listener)


	//http
	//mux := http.NewServeMux()
	//mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
	//	server.ServeHTTP(w,r)
	//})
	//httpServer := &http.Server{
	//	Addr:    ":8081",
	//	Handler: mux,
	//}
	//_ = httpServer.ListenAndServeTLS("keys/server.crt", "keys/server_no_password.key")
}
