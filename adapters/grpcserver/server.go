// SPDX-License-Identifier: MIT
// SPDX-FileCopyright-Text: 2025 Cristopher James
// SPDX-FileCopyright-Text: 2025 kurth4cker

package grpcserver

import (
	"context"

	"github.com/kurth4cker/go-specs-greet/domain/interactions"
)

type GreetServer struct {
	UnimplementedGreeterServer
}

func (g GreetServer) Curse(ctx context.Context, request *CurseRequest) (*CurseReply, error) {
	return &CurseReply{
		Message: interactions.Curse(request.Name),
	}, nil
}

func (g GreetServer) Greet(ctx context.Context, request *GreetRequest) (*GreetReply, error) {
	return &GreetReply{
		Message: interactions.Greet(request.Name),
	}, nil
}
