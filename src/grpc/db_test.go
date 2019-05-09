package main

import (
	"math/rand"
	"strconv"
	"testing"
)

func TestQueryInfo_1(t *testing.T) {
	ui, err := queryInfo("test")
	if err != nil {
		t.Log(err)
	} else {
		t.Log("nickname: ", ui.nickname, "; profile: ", ui.profile)
	}
}

func TestUploadProfile(t *testing.T) {
	username := "test"
	profile := []byte{uint8(rand.Int()), uint8(rand.Int())}
	err := uploadProfile(&username, &profile)
	if err != nil {
		t.Log(err)
	}
}

func TestChangeNickname(t *testing.T) {
	username := "test"
	nickname := "test_change_" + strconv.Itoa(rand.Int())
	err := changeNickname(&username, &nickname)
	if err != nil {
		t.Log(err)
	}
}

func TestQueryInfo_2(t *testing.T) {
	ui, err := queryInfo("test")
	if err != nil {
		t.Log(err)
	} else {
		t.Log("nickname: ", ui.nickname, "; profile: ", ui.profile)
	}
}
