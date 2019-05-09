package main

import (
	"context"
	pb "shopee/entrytask-kuangdi.bao/src/pb"
	"testing"
)

func TestLogin(t *testing.T) {
	s := server{}

	in := &pb.LoginRequest{Username: "test", Password: "123456"}
	out, err := s.Login(context.Background(), in)
	if err != nil {
		t.Log(err)
	} else {
		t.Log(out)
	}
}
