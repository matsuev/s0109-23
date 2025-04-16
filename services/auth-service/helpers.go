package main

import "s0109-23/internal/proxyproto"

func respondError(code uint32, msg string) *proxyproto.ConnectResponse {
	return &proxyproto.ConnectResponse{
		Error: &proxyproto.Error{
			Code:    code,
			Message: msg,
		},
	}
}

func respondResult(user string) *proxyproto.ConnectResponse {
	return &proxyproto.ConnectResponse{
		Result: &proxyproto.ConnectResult{
			User: user,
		},
	}
}
