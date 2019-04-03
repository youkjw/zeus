package main

import (
	//_ "github.com/astaxie/beego/cache/redis"
	pb "zeus/cmd/grpc-linker/grpc-server/proto"
	"flag"
	"google.golang.org/grpc/reflection"
	"net"
	"log"
	"google.golang.org/grpc"
)
func main() {
	pt := flag.String("p", "8188", "port of grpc server")
	flag.Parse()
	lis, err := net.Listen("tcp", ":" + *pt)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	} else {
		log.Printf("Grpc serve on port:" + *pt)
	}
	s := grpc.NewServer()
	pb.RegisterApiauthServer(s, &ApiAuthServer{})
	//pb.RegisterCasbinServer(s, server.NewServer())
	// Register reflection service on gRPC server.
	reflection.Register(s)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}