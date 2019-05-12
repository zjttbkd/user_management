package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
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
	// create img directory
	if _, err := os.Stat("./img"); err != nil {
		err := os.MkdirAll("./img", 0777)

		if err != nil {
			log.Fatalln("create directory err")
		}
	}

	// create log directory
	if _, err := os.Stat("./log"); err != nil {
		err := os.MkdirAll("./log", 0777)

		if err != nil {
			log.Fatalln("create directory err")
		}
	}

	// create grpc client
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	client = pb.NewUsrmgnClient(conn)

	// set gin config
	reqlog, _ := os.Create("./log/gin.log")
	errlog, _ := os.Create("./log/err.log")
	gin.DefaultWriter = reqlog
	gin.DefaultErrorWriter = errlog
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
		log.Println(err.Error())
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
		"username": username,
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

func main() {
	defer conn.Close()

	r := gin.Default()
	r.MaxMultipartMemory = 8 << 20 // 8 MiB
	r.LoadHTMLGlob("../html/*")
	r.Static("/img", "./img")

	r.GET("/", index)
	r.GET("/index", index)
	r.GET("/login", signIn)
	r.POST("/login", login)
	r.POST("/upload", uploadProfile)
	r.POST("/change", changeNickname)

	// Run http_server
	if err := r.Run(":8052"); err != nil {
		log.Fatalf("could not run http_server: %v", err)
	}
}
