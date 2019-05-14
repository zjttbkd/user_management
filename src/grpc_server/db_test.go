package main

import (
	"fmt"
	"math/rand"
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
		t.Log("Nickname: ", ui.Nickname, "; Profile: ", ui.Profile)
	}
}

func TestUploadProfile(t *testing.T) {
	username := "test"
	profile := fmt.Sprintf("test_%v.img", rand.Int())
	err := uploadProfile(username, profile)
	if err != nil {
		t.Error(err)
	}
}

func TestChangeNickname(t *testing.T) {
	username := "test"
	nickname := fmt.Sprintf("test_change_%v", rand.Int())
	err := changeNickname(username, nickname)
	if err != nil {
		t.Error(err)
	}
}
