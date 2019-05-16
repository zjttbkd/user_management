package main

import (
	"database/sql"
	"errors"
	log "github.com/cihub/seelog"
	_ "github.com/go-sql-driver/mysql"
	"time"
)

var db *sql.DB

func init() {
	db, _ = sql.Open("mysql", "root:passw0rd@tcp(172.18.0.22)/user_mgn_db?charset=utf8mb4")
	db.SetMaxIdleConns(1000)
	db.SetMaxOpenConns(2000)
	db.SetConnMaxLifetime(time.Second * 60)
	if err := db.Ping(); err != nil {
		panic(err)
	}
}

func queryInfo(username string) (*userInfo, error) {
	row := db.QueryRow("select password, Nickname, Profile from user_info_tab_00000000 where username = ?", username)

	ui := userInfo{username: username}
	err := row.Scan(&ui.password, &ui.Nickname, &ui.Profile)

	return &ui, err
}

func uploadProfile(username string, profile string) error {
	stmt, err := db.Prepare("update user_info_tab_00000000 set Profile= ?, modify_time=UNIX_TIMESTAMP() where username = ?")
	if err != nil {
		return err
	}

	res, err := stmt.Exec(profile, username)
	stmt.Close()
	if err != nil {
		log.Error(err)
		return err
	}

	num, err := res.RowsAffected()
	if err != nil {
		log.Error(err)
		return err
	}

	if num != 1 {
		return errors.New("update affected row != 1")
	}

	return nil
}

func changeNickname(username string, nickname string) error {
	stmt, err := db.Prepare("update user_info_tab_00000000 set Nickname= ?, modify_time=UNIX_TIMESTAMP() where username = ?")
	if err != nil {
		return err
	}

	res, err := stmt.Exec(nickname, username)
	stmt.Close()
	if err != nil {
		log.Error(err)
		return err
	}

	num, err := res.RowsAffected()
	if err != nil {
		log.Error(err)
		return err
	}

	if num != 1 {
		return errors.New("update affected row != 1")
	}

	return nil
}
