package main

import (
	"context"
	"encoding/json"
	"s0109-23/internal/demobase"
	"s0109-23/internal/proxyproto"
	"strconv"
)

// Service ...
type Service struct {
	proxyproto.UnimplementedCentrifugoProxyServer
	store demobase.Querier
}

// NewService ...
func NewService(store demobase.Querier) *Service {
	return &Service{
		store: store,
	}
}

// Connect ...
func (s *Service) Connect(ctx context.Context, req *proxyproto.ConnectRequest) (*proxyproto.ConnectResponse, error) {
	type AuthRequest struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}

	authRequest := &AuthRequest{}

	if err := json.Unmarshal(req.Data, authRequest); err != nil {
		return respondError(107, "bad request"), nil
	}

	account, err := s.store.UserLogin(ctx, demobase.UserLoginParams{
		Username: authRequest.Username,
		Password: authRequest.Password,
	})
	if err != nil {
		return respondError(101, "unauthorized"), nil
	}

	if !account.Enabled {
		return respondError(101, "unauthorized"), nil
	}

	return respondResult(strconv.Itoa(int(account.ID))), nil

}
