package main

import (
	"code/common"
	"code/model"
	c "code/src/client"
	ss "code/src/server"
	"crypto/tls"
	"crypto/x509"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

const (
	serverCertFile   = "cert/server-cert.pem"
	serverKeyFile    = "cert/server-key.pem"
	clientCACertFile = "cert/ca-cert.pem"
)

func loadTLSCredentials() (credentials.TransportCredentials, error) {
	// Load certificate of the CA who signed client's certificate
	pemClientCA, err := ioutil.ReadFile(clientCACertFile)
	if err != nil {
		return nil, err
	}

	certPool := x509.NewCertPool()
	if !certPool.AppendCertsFromPEM(pemClientCA) {
		return nil, fmt.Errorf("failed to add client CA's certificate")
	}

	// Load server's certificate and private key
	serverCert, err := tls.LoadX509KeyPair(serverCertFile, serverKeyFile)
	if err != nil {
		return nil, err
	}

	// Create the credentials and return it
	config := &tls.Config{
		Certificates: []tls.Certificate{serverCert},
		ClientAuth:   tls.RequireAndVerifyClientCert,
		ClientCAs:    certPool,
	}

	return credentials.NewTLS(config), nil
}

func main() {

	cmd := flag.String("option", "", "")
	flag.Parse()

	if *cmd == `server` {

		serverOptions := []grpc.ServerOption{}
		tlsCredentials, err := loadTLSCredentials()
		if err != nil {
			log.Printf("cannot load TLS credentials: %w", err)
		}

		serverOptions = append(serverOptions, grpc.Creds(tlsCredentials))

		ls, err := net.Listen("tcp", fmt.Sprintf("localhost:%s", common.PORT))
		if err != nil {
			log.Fatalf("failed to listen: %v", err)
		}

		grpcServer := grpc.NewServer(serverOptions...)
		model.RegisterAccountManagementServer(grpcServer, ss.NewAccountServer())
		err = grpcServer.Serve(ls)
		if err != nil {
			log.Fatalf("failed to listen: %v", err)
		}
		return
	}

	if *cmd == `simply-client` {
		c.CreateAccount()
		return
	}

	if *cmd == `stream-client` {
		c.CreateAccountBulk()
		return
	}

	fmt.Println(`please run using this: go run main.go --option={YOUR_OPTION}`)
}
