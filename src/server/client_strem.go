package server

import (
	"code/model"
	"fmt"
	"io"
)

func (s *AccountServer) BulkCreate(stream model.AccountManagement_BulkCreateServer) error {
	var countAccount int64
	for {
		acc, err := stream.Recv()
		if err == io.EOF {
			return stream.SendAndClose(
				&model.CreateAccountBulkResponse{
					CountAccount: countAccount,
					Message:      "Success & close",
				},
			)
		}
		if err != nil {
			return err
		}

		// if acc.Id == 4 {
		// 	return stream.SendAndClose(
		// 		&model.CreateAccountBulkResponse{
		// 			CountAccount: countAccount,
		// 			Message:      "failed id 4",
		// 		},
		// 	)
		// }

		countAccount++
		tmpAccount := *acc
		fmt.Println(tmpAccount)
	}
}
