package server

import "code/model"

type AccountServer struct {
	model.UnimplementedAccountManagementServer
}

func NewAccountServer() *AccountServer {
	return &AccountServer{}
}
