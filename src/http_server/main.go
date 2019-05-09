package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
	pb "shopee/entrytask-kuangdi.bao/src/pb"
)

// TODO v1 use only one conn
var conn *grpc.ClientConn
var client pb.UsrmgnClient

func init() {
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	client = pb.NewUsrmgnClient(conn)
}

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

	c.JSON(http.StatusOK, gin.H{
		"status":   "SUCCESS",
		"nickname": res.Nickname,
		"profile":  res.Profile,
	})
}

func main() {
	defer conn.Close()

	r := gin.Default()
	r.LoadHTMLGlob("../html/*")
	r.GET("/*index", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", nil)
	})
	r.POST("/form_post", login)

	// Run http http_server
	if err := r.Run(":8052"); err != nil {
		log.Fatalf("could not run http_server: %v", err)
	}
}
