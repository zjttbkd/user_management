package main

import (
	log "github.com/cihub/seelog"
)

func init() {
	logger, err := log.LoggerFromConfigAsFile("config/seelog.xml")
	if err != nil {
		panic(err)
	}
	log.ReplaceLogger(logger)
}

func main() {
	defer conn.Close() // close grpc server conn
	defer log.Flush()

	// Run http_server
	r := getRouter()
	if err := r.Run(":8052"); err != nil {
		log.Errorf("could not run http_server: %v", err)
		panic(err)
	}
}
