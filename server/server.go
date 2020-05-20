package server
import (
    "github.com/gin-gonic/gin"
)

func Start() {
    router := gin.Default()
    handler := &Handler{}
    router.GET("/ping",handler.Ping)
    router.GET("/time", handler.DisplayTime)
    router.Run(":8888")
}
