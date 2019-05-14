package main

import (
	"errors"
	"golang.org/x/net/context"
	pb "shopee/entrytask-kuangdi.bao/src/pb"
)

type userInfo struct {
	username string
	password string
	Nickname string
	Profile  string
}

type server struct{}

func (s *server) Login(ctx context.Context, in *pb.LoginRequest) (*pb.UsrInfoReply, error) {
	ui, err := queryInfo(in.Username)
	if err != nil {
		return &pb.UsrInfoReply{}, errors.New("user not existed")
	}

	if ui.password != in.Password {
		return &pb.UsrInfoReply{}, errors.New("wrong password")
	}

	cacheUserInfo(ui) // store user info into redis

	return &pb.UsrInfoReply{Username: ui.username, Nickname: ui.Nickname, Profile: ui.Profile}, nil
}

func (s *server) Query(ctx context.Context, in *pb.QueryRequest) (*pb.UsrInfoReply, error) {
	ui, err := fetchUserInfo(in.Username) // fetch from redis first
	if err != nil {
		ui, err = queryInfo(in.Username)
		if err != nil {
			return &pb.UsrInfoReply{}, errors.New("user not existed")
		}
		cacheUserInfo(ui) // store user info into redis
	}

	return &pb.UsrInfoReply{Username: ui.username, Nickname: ui.Nickname, Profile: ui.Profile}, nil
}

func (s *server) UploadProfile(ctx context.Context, in *pb.ProfileRequest) (*pb.UsrInfoReply, error) {
	delUserInfo(in.Username) // delete from redis

	err := uploadProfile(in.Username, in.Profile)
	if err != nil {
		return &pb.UsrInfoReply{}, err
	}
	return &pb.UsrInfoReply{Username: in.Username}, nil
}

func (s *server) ChangeNickname(ctx context.Context, in *pb.NicknameRequest) (*pb.UsrInfoReply, error) {
	delUserInfo(in.Username) // delete from redis

	err := changeNickname(in.Username, in.Nickname)
	if err != nil {
		return &pb.UsrInfoReply{}, err
	}
	return &pb.UsrInfoReply{Username: in.Username}, nil
}
