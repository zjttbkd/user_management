package main

import (
	"errors"
	"log"
	"net"

	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	pb "shopee/entrytask-kuangdi.bao/src/pb"
)

type server struct{}

func (s *server) Login(ctx context.Context, in *pb.LoginRequest) (*pb.LoginReply, error) {
	log.Printf("in username: %v", in.Username)

	ui, err := queryInfo(in.Username)
	if err != nil {
		return &pb.LoginReply{}, errors.New("user not existed")
	}

	if ui.password != in.Password {
		return &pb.LoginReply{}, errors.New("wrong password")
	}

	return &pb.LoginReply{Username: ui.username, Nickname: ui.nickname, Profile: ui.profile}, nil
}

func main() {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterUsrmgnServer(s, &server{})

	reflection.Register(s)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
