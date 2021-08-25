package main

import (
	"code/model"
	c "code/src/client"
	ss "code/src/server"
	"flag"
	"fmt"
	"log"
	"net"

	"google.golang.org/grpc"
)

func main() {

	cmd := flag.String("option", "", "")
	flag.Parse()

	if *cmd == `server` {
		ls, err := net.Listen("tcp", ":9000")
		if err != nil {
			log.Fatalf("failed to listen: %v", err)
		}
		grpcServer := grpc.NewServer()
		model.RegisterAccountManagementServer(grpcServer, ss.NewAccountServer())
		grpcServer.Serve(ls)
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
