package server
import (
    "github.com/gin-gonic/gin"
    "fmt"
)

func Start(port int) {
    router := gin.Default()
    handler := &Handler{}
    router.GET("/ping",handler.Ping)
    router.GET("/time", handler.DisplayTime)
    addr := fmt.Sprintf(":%d", port)
    router.Run(addr)
}
