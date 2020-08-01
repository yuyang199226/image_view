package server

import (
	"fmt"
	"github.com/gin-gonic/gin"
	// "net/http"
)

func Start(port int) {
	router := gin.Default()
	handler := &Handler{}
	router.Use(Logger(), gin.Recovery())
	router.GET("/ping", handler.Ping)
	router.GET("/time", handler.DisplayTime)
	router.POST("/upload", handler.Upload)
	// http.HandleFunc("/v1/ws", handler.TailLog)
	router.GET("/v1/ws", handler.TailLog)
	addr := fmt.Sprintf("0.0.0.0:%d", port)
	router.Run(addr)
}
