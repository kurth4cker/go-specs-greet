// SPDX-License-Identifier: MIT
// SPDX-FileCopyright-Text: 2025 Cristopher James
// SPDX-FileCopyright-Text: 2025 kurth4cker

package main

import (
	"context"
	"log"
	"net"

	"github.com/kurth4cker/go-specs-greet/adapters/grpcserver"
	"github.com/kurth4cker/go-specs-greet/domain/interactions"
	"google.golang.org/grpc"
)

func check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	lis, err := net.Listen("tcp", ":50051")
	check(err)
	s := grpc.NewServer()
	grpcserver.RegisterGreeterServer(s, &GreetServer{})

	err = s.Serve(lis)
	check(err)
}

type GreetServer struct {
	grpcserver.UnimplementedGreeterServer
}

func (g GreetServer) Greet(ctx context.Context, request *grpcserver.GreetRequest) (*grpcserver.GreetReply, error) {
	return &grpcserver.GreetReply{
		Message: interactions.Greet(request.Name),
	}, nil
}
