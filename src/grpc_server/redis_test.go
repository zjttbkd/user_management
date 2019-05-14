package main

import (
	"fmt"
	"math/rand"
	"testing"
	"time"
)

var ui *userInfo

func init() {
	rand.Seed(time.Now().Unix())
	username := fmt.Sprintf("test_%v", rand.Int())
	ui = &userInfo{username, "", "test_for_reids", "no profile"}
}

func TestCacheUserInfo(t *testing.T) {
	cacheUserInfo(ui)
}

func TestFetchUserInfo(t *testing.T) {
	var res *userInfo
	res, err := fetchUserInfo(ui.username)
	if err != nil {
		t.Error(err)
	} else {
		t.Log("Username: ", ui.username, "; Nickname: ", res.Nickname, "; Profile: ", res.Profile)
	}
}

func TestDelUserInfo(t *testing.T) {
	delUserInfo(ui.username)
}
