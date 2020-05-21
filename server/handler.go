package server

import (
    "log"
    "net/http"
    "github.com/gin-gonic/gin"
    "imageview/config"
)

type Handler struct {

}



func (handler *Handler) DisplayTime(c *gin.Context) {
    log.Println("display time")
    c.String(http.StatusOK, "Hello time")
}


func (handler *Handler) Ping(c *gin.Context) {
    log.Println("Ping")
    cfg := config.Config()
    c.JSON(http.StatusOK, gin.H{"message": cfg.Owner.Name})
}
