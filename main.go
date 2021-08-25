package main

import (
	"code/common"
	"code/model"
	c "code/src/client"
	ss "code/src/server"
	"flag"
	"fmt"
	"log"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

func main() {

	cmd := flag.String("option", "", "")
	flag.Parse()

	if *cmd == `server` {
		ls, err := net.Listen("tcp", fmt.Sprintf("localhost:%s", common.PORT))
		if err != nil {
			log.Fatalf("failed to listen: %v", err)
		}
		creds, err := credentials.NewClientTLSFromFile("./cert/server.crt", "./cert/server.key")
		if err != nil {
			log.Fatalf("failed to listen: %v", err)
		}

		opts := []grpc.ServerOption{grpc.Creds(creds)}

		grpcServer := grpc.NewServer(opts...)
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
