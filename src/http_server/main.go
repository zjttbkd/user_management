package main

import (
	"fmt"
	"log"
	"net/http"
	"path"
	"path/filepath"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
	pb "shopee/entrytask-kuangdi.bao/src/pb"
)

// TODO connection pool
var conn *grpc.ClientConn
var client pb.UsrmgnClient

func init() {
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	client = pb.NewUsrmgnClient(conn)
}

func index(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", nil)
}

func login(c *gin.Context) {
	var username, password string
	var authorized bool

	cookie, err := c.Cookie("gin_cookie")
	if err != nil || cookie == "" {
		username = c.PostForm("username")
		password = c.PostForm("password")
	} else {
		username = cookie
		authorized = true
	}

	req := &pb.LoginRequest{Username: username, Password: password, Authorized: authorized}
	res, err := client.Login(c, req)
	if err != nil {
		c.HTML(http.StatusInternalServerError, "result.html", gin.H{
			"result": err.Error(),
		})
		return
	}

	if !authorized {
		c.SetCookie("gin_cookie", username, 360, "/", "192.168.33.10", false, false)
	}

	c.HTML(http.StatusOK, "usrinfo.html", gin.H{
		"username": username,
		"nickname": res.Nickname,
		"profile":  res.Profile,
	})
}

func uploadProfile(c *gin.Context) {
	cookie, err := c.Cookie("gin_cookie")
	if err != nil {
		c.HTML(http.StatusOK, "index.html", nil)
	}

	file, err := c.FormFile("profile")
	if err != nil {
		c.HTML(http.StatusBadRequest, "result.html", gin.H{
			"result": fmt.Sprintf("get form err: %s", err.Error()),
		})
		return
	}

	profile := "./img/" + cookie + "_" + strconv.FormatInt(time.Now().Unix(), 10) + path.Ext(filepath.Base(file.Filename))
	if err := c.SaveUploadedFile(file, profile); err != nil {
		c.HTML(http.StatusBadRequest, "result.html", gin.H{
			"result": fmt.Sprintf("upload file err: %s", err.Error()),
		})
		return
	}

	req := &pb.ProfileRequest{Username: cookie, Profile: profile}
	_, err = client.UploadProfile(c, req)
	if err != nil {
		c.HTML(http.StatusInternalServerError, "result.html", gin.H{
			"result": fmt.Sprintf("upload file err: %s", err.Error()),
		})
		return
	}

	c.HTML(http.StatusOK, "result.html", gin.H{
		"result": fmt.Sprintf("File %s uploaded successfully.", file.Filename),
	})
}

func changeNickname(c *gin.Context) {
	cookie, err := c.Cookie("gin_cookie")
	if err != nil {
		c.HTML(http.StatusOK, "index.html", nil)
	}

	nickname := c.PostForm("nickname")

	req := &pb.NicknameRequest{Username: cookie, Nickname: nickname}
	_, err = client.ChangeNickname(c, req)
	if err != nil {
		c.HTML(http.StatusInternalServerError, "result.html", gin.H{
			"result": fmt.Sprintf("change nickname err: %s", err.Error()),
		})
		return
	}

	c.HTML(http.StatusOK, "result.html", gin.H{
		"result": fmt.Sprintln("Change nickname successfully."),
	})
}

func main() {
	defer conn.Close()

	r := gin.Default()
	r.MaxMultipartMemory = 8 << 20 // 8 MiB
	r.LoadHTMLGlob("../html/*")
	r.Static("/img", "./img")

	r.GET("/", index)
	r.GET("/index", index)
	r.GET("/login", login)
	r.POST("/login", login)
	r.POST("/upload", uploadProfile)
	r.POST("/change", changeNickname)

	// Run http_server
	if err := r.Run(":8052"); err != nil {
		log.Fatalf("could not run http_server: %v", err)
	}
}
