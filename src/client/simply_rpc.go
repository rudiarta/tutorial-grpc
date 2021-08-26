package client

import (
	"code/common"
	"code/model"
	"context"
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"io/ioutil"
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

func loadTLSCredentials() (credentials.TransportCredentials, error) {
	// Load certificate of the CA who signed server's certificate
	pemServerCA, err := ioutil.ReadFile("cert/ca-cert.pem")
	if err != nil {
		return nil, err
	}

	certPool := x509.NewCertPool()
	if !certPool.AppendCertsFromPEM(pemServerCA) {
		return nil, fmt.Errorf("failed to add server CA's certificate")
	}

	// Load client's certificate and private key
	clientCert, err := tls.LoadX509KeyPair("cert/client-cert.pem", "cert/client-key.pem")
	if err != nil {
		return nil, err
	}

	// Create the credentials and return it
	config := &tls.Config{
		Certificates: []tls.Certificate{clientCert},
		RootCAs:      certPool,
	}

	return credentials.NewTLS(config), nil
}

func CreateAccount() {
	transportOption := grpc.WithInsecure()
	tlsCredentials, err := loadTLSCredentials()
	if err != nil {
		log.Fatal("cannot load TLS credentials: ", err)
	}

	transportOption = grpc.WithTransportCredentials(tlsCredentials)

	conn, err := grpc.Dial(fmt.Sprintf("localhost:%s", common.PORT), transportOption)
	if err != nil {
		log.Fatalf("failed to Dial: %v", err)
	}
	defer conn.Close()

	client := model.NewAccountManagementClient(conn)
	res, err := client.Create(context.Background(), &model.Account{
		Id:       1,
		Username: "rudi",
		Password: `PasswordSupersecret`,
	})
	if err != nil {
		log.Fatalf("err response: %v", err)
	}

	fmt.Println(res)
}

func CreateAccountBulk() {
	transportOption := grpc.WithInsecure()
	tlsCredentials, err := loadTLSCredentials()
	if err != nil {
		log.Fatal("cannot load TLS credentials: ", err)
	}

	transportOption = grpc.WithTransportCredentials(tlsCredentials)

	conn, err := grpc.Dial(fmt.Sprintf("localhost:%s", common.PORT), transportOption)
	if err != nil {
		log.Fatalf("failed to Dial: %v", err)
	}
	defer conn.Close()

	client := model.NewAccountManagementClient(conn)
	stream, err := client.BulkCreate(context.Background())
	if err != nil {
		log.Fatalf("%v.BulkCreate(_) = _, %v", client, err)
	}

	accountOne := &model.Account{
		Id:       1,
		Username: "user satu",
		Password: `PasswordSupersecret`,
	}
	accountTwo := &model.Account{
		Id:       2,
		Username: "user dua",
		Password: `PasswordSupersecret`,
	}
	accountThree := &model.Account{
		Id:       3,
		Username: "user tiga",
		Password: `PasswordSupersecret`,
	}
	accountFour := &model.Account{
		Id:       4,
		Username: "user empat",
		Password: `PasswordSupersecret`,
	}
	accountFive := &model.Account{
		Id:       5,
		Username: "user lima",
		Password: `PasswordSupersecret`,
	}
	accounts := []model.Account{
		*accountOne, *accountTwo, *accountThree, *accountFour, *accountFive,
	}

	for _, acc := range accounts {
		err = stream.Send(&acc)
		if err != nil {
			log.Fatalf("%v.Send(%v) = %v", stream, accountOne, err)
		}
	}

	reply, err := stream.CloseAndRecv()
	if err != nil {
		log.Fatalf("%v.CloseAndRecv() got error %v, want %v", stream, err, nil)
	}
	log.Printf("Account summary: %v", reply)
}
