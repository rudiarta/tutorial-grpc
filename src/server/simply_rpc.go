package server

import (
	"code/model"
	"context"
	"encoding/json"
	"fmt"
)

func (s *AccountServer) Create(ctx context.Context, a *model.Account) (*model.CreateAccountResponse, error) {
	fmt.Println(a)
	data, _ := json.Marshal(a)
	return &model.CreateAccountResponse{
		Message: string(data),
	}, nil
}
