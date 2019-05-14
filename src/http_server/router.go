package main

import (
	"fmt"
	log "github.com/cihub/seelog"
	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
	"net/http"
	"path"
	"path/filepath"
	pb "shopee/entrytask-kuangdi.bao/src/pb"
	"strconv"
	"time"
)

var conn *grpc.ClientConn
var client pb.UsrmgnClient

// token binding post from json
type token struct {
	Token string `json: "token" binding: required`
}

func init() {
	// create grpc client
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		panic(err)
	}
	client = pb.NewUsrmgnClient(conn)

	// set gin config
	gin.SetMode(gin.ReleaseMode)
}

// user info page
func index(c *gin.Context) {
	cookie, err := c.Cookie("gin_cookie")
	if err != nil {
		c.Redirect(http.StatusFound, "login")
		return
	}

	req := &pb.QueryRequest{Username: cookie}
	res, err := client.Query(c, req)
	if err != nil {
		log.Error(err.Error())
		c.Redirect(http.StatusFound, "login")
		return
	}

	c.HTML(http.StatusOK, "usrinfo.html", gin.H{
		"username": res.Username,
		"nickname": res.Nickname,
		"profile":  res.Profile,
	})
}

// sign in page
func signIn(c *gin.Context) {
	c.HTML(http.StatusOK, "login.html", nil)
}

// login action
func login(c *gin.Context) {
	username := c.PostForm("username")
	password := c.PostForm("password")

	req := &pb.LoginRequest{Username: username, Password: password}
	res, err := client.Login(c, req)
	if err != nil {
		c.HTML(http.StatusInternalServerError, "result.html", gin.H{
			"result": err.Error(),
		})
		return
	}

	c.SetCookie("gin_cookie", username, 360, "/", "192.168.33.10", false, false)

	c.HTML(http.StatusOK, "usrinfo.html", gin.H{
		"username": res.Username,
		"nickname": res.Nickname,
		"profile":  res.Profile,
	})
}

// upload profile action
func uploadProfile(c *gin.Context) {
	cookie, err := c.Cookie("gin_cookie")
	if err != nil {
		c.Redirect(http.StatusFound, "login")
	}

	file, err := c.FormFile("profile")
	if err != nil {
		c.HTML(http.StatusBadRequest, "result.html", gin.H{
			"result": fmt.Sprintf("get form err: %s", err.Error()),
		})
		return
	}

	imgName := fmt.Sprint(cookie, "_", strconv.FormatInt(time.Now().Unix(), 10), path.Ext(filepath.Base(file.Filename)))
	profile := path.Join("./img/", imgName)
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

// change nickname action
func changeNickname(c *gin.Context) {
	cookie, err := c.Cookie("gin_cookie")
	if err != nil {
		c.Redirect(http.StatusFound, "login")
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

func getUserInfo(c *gin.Context) {
	var t token
	if c.BindJSON(&t) != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"message": "no token",
		})
		return
	}

	if t.Token != "test_token" {
		c.JSON(http.StatusUnauthorized, gin.H{
			"message": "token authorize failed",
		})
		return
	}

	username := c.Param("username")

	req := &pb.QueryRequest{Username: username}
	res, err := client.Query(c, req)
	if err != nil {
		log.Error(err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "query user info failed",
			"error":   err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message":  "ok",
		"profile":  res.Profile,
		"nickname": res.Nickname,
	})

}

func logger() gin.HandlerFunc {
	return func(c *gin.Context) {
		log.Debugf("%v %v", c.Request.Method, c.Request.RequestURI)
	}
}

func getRouter() *gin.Engine {
	r := gin.New()
	r.Use(logger(), gin.Recovery())
	r.MaxMultipartMemory = 8 << 20 // 8 MiB
	r.LoadHTMLGlob("./html/*")
	r.Static("/img", "./img")

	r.GET("/", index)
	r.GET("/index", index)
	r.GET("/login", signIn)
	r.POST("/login", login)
	r.POST("/upload", uploadProfile)
	r.POST("/change", changeNickname)
	r.POST("/usrinfo/:username", getUserInfo)

	return r
}
