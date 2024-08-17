package endpoints

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

func RunServer(ip string, port uint) {
	// Set up gin
	r := gin.Default()
	r.GET("/publicKey", GetServerPublicKey)
	r.PUT("/register", Register)
	r.POST("/login", Login)
	r.PUT("/send", SendMessage)
	r.GET("/new", NewMessages)
	r.GET("/newWS", NewMessagesWebsocket)
	r.GET("/fetch/:from", FetchMessagesFromUser)

	// Run API server
	if err := r.Run(fmt.Sprintf("%s:%d", ip, port)); err != nil {
		panic(err)
	}
}
