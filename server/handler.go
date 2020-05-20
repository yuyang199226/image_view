package server

import (
    "log"
    "net/http"
    "github.com/gin-gonic/gin"
)

type Handler struct {

}



func (handler *Handler) DisplayTime(c *gin.Context) {
    log.Println("display time")
    c.String(http.StatusOK, "Hello time")
}


func (handler *Handler) Ping(c *gin.Context) {
    log.Println("Ping")
    c.JSON(http.StatusOK, gin.H{"message": "ping"})
}
