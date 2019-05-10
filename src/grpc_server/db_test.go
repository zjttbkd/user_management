package main

import (
	"math/rand"
	"strconv"
	"testing"
	"time"
)

func init() {
	rand.Seed(time.Now().Unix())
}

func TestQueryInfo(t *testing.T) {
	ui, err := queryInfo("test")
	if err != nil {
		t.Error(err)
	} else {
		t.Log("nickname: ", ui.nickname, "; profile: ", ui.profile)
	}
}

func TestUploadProfile(t *testing.T) {
	username := "test"
	profile := "test_" + strconv.Itoa(rand.Int()) + ".img"
	err := uploadProfile(&username, &profile)
	if err != nil {
		t.Error(err)
	}
}

func TestChangeNickname(t *testing.T) {
	username := "test"
	nickname := "test_change_" + strconv.Itoa(rand.Int())
	err := changeNickname(&username, &nickname)
	if err != nil {
		t.Error(err)
	}
}
