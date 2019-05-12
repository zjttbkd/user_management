package main

import (
	"context"
	"math/rand"
	pb "shopee/entrytask-kuangdi.bao/src/pb"
	"strconv"
	"testing"
	"time"
)

func init() {
	rand.Seed(time.Now().Unix())
}

func TestServer_Login(t *testing.T) {
	s := server{}

	in := []pb.LoginRequest{{Username: "test", Password: "123456"}, {Username: "test", Password: "654321"}}
	for _, tt := range in {
		out, err := s.Login(context.Background(), &tt)
		if err != nil {
			if err.Error() != "wrong password" {
				t.Error(err)
			}
		} else {
			t.Log(out)
		}
	}
}

func TestServer_Query(t *testing.T) {
	s := server{}

	in := &pb.QueryRequest{Username: "test"}
	out, err := s.Query(context.Background(), in)
	if err != nil {
		t.Error(err)
	} else {
		t.Log(out)
	}
}

func TestServer_UploadProfile(t *testing.T) {
	s := server{}

	in := &pb.ProfileRequest{Username: "test", Profile: "/img/test.png"}
	out, err := s.UploadProfile(context.Background(), in)
	if err != nil {
		t.Error(err)
	} else {
		t.Log(out)
	}
}

func TestServer_ChangeNickname(t *testing.T) {
	s := server{}

	in := &pb.NicknameRequest{Username: "test", Nickname: "test_change_" + strconv.Itoa(rand.Int())}
	out, err := s.ChangeNickname(context.Background(), in)
	if err != nil {
		t.Error(err)
	} else {
		t.Log(out)
	}
}
