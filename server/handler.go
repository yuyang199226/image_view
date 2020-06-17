package server

import (
    log "github.com/sirupsen/logrus"
    "net/http"
    "github.com/gin-gonic/gin"
    "imageview/config"
    "os"
    "io"
    "fmt"
)
func init () {
log.SetFormatter(&log.JSONFormatter{})

}
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


func (handler *Handler) Upload(c *gin.Context) {
    file, header, err := c.Request.FormFile("upload")
    if err != nil {
        c.String(http.StatusBadRequest, "Bad request")
        return
    }
    filename := header.Filename
    fmt.Println(file, err, filename)
    // 创建文件，文件名filename
    out, err := os.Create(filename)
    if err != nil {
        log.Fatal(err)
    }
    defer out.Close()
    _, err = io.Copy(out, file)
    if err != nil {
        log.Fatal(err)
    }
    c.String(http.StatusCreated, "upload successfule \n")
}


func Add(x , y int) int {
    return x+y
}
