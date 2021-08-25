package client

import (
	"code/common"
	"code/model"
	"context"
	"fmt"
	"log"

	"google.golang.org/grpc"
)

func CreateAccount() {
	conn, err := grpc.Dial(fmt.Sprintf(":%s", common.PORT), grpc.WithInsecure())
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
	conn, err := grpc.Dial(fmt.Sprintf(":%s", common.PORT), grpc.WithInsecure())
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
