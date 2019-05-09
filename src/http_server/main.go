package main

import (
	"log"
	"net/http"

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

// TODO token
func login(c *gin.Context) {
	username := c.PostForm("username")
	password := c.PostForm("password")

	req := &pb.LoginRequest{Username: username, Password: password}
	res, err := client.Login(c, req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.HTML(http.StatusOK, "usrinfo.html", gin.H{
		"username": username,
		"nickname": res.Nickname,
		"profile":  res.Profile,
	})
}

func main() {
	defer conn.Close()

	r := gin.Default()
	r.LoadHTMLGlob("../html/*")
	r.Static("../img", "./img")

	r.GET("/", index)
	r.GET("/index", index)
	r.POST("/login", login)

	// Run http http_server
	if err := r.Run(":8052"); err != nil {
		log.Fatalf("could not run http_server: %v", err)
	}
}
