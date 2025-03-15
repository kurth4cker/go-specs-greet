// SPDX-License-Identifier: MIT
// SPDX-FileCopyright-Text: 2025 Cristopher James
// SPDX-FileCopyright-Text: 2025 kurth4cker

package main

import (
	"log"
	"net"

	"github.com/kurth4cker/go-specs-greet/adapters/grpcserver"
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
	grpcserver.RegisterGreeterServer(s, &grpcserver.GreetServer{})

	err = s.Serve(lis)
	check(err)
}
