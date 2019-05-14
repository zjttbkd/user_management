package main

import (
	log "github.com/cihub/seelog"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"net"
	pb "shopee/entrytask-kuangdi.bao/src/pb"
)

func init() {
	logger, err := log.LoggerFromConfigAsFile("config/seelog.xml")
	if err != nil {
		panic(err)
	}
	log.ReplaceLogger(logger)
}

func main() {
	defer db.Close()     // close mysql conn
	defer client.Close() // close redis conn
	defer log.Flush()

	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Errorf("failed to listen: %v", err)
		panic(err)
	}
	s := grpc.NewServer()
	pb.RegisterUsrmgnServer(s, &server{})

	reflection.Register(s)
	if err := s.Serve(lis); err != nil {
		log.Errorf("failed to serve: %v", err)
		panic(err)
	}
}
