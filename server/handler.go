package server

import (
	"fmt"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"imageview/config"
	"io"
	"net/http"
	"os"
	"github.com/gorilla/websocket"
)

func init() {
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
var upgrader = websocket.Upgrader{}
func (handler *Handler) TailLog(w http.ResponseWriter, r *http.Request) {
	conn, _ := upgrader.Upgrade(w, r, nil )
	go func(conn *websocket.Conn) {
		for {
			msgType, msg, _ := conn.ReadMessage()
			err := conn.WriteMessage(msgType, msg)
			if err != nil {
				break
			}
		}
	}(conn)
}

func Add(x, y int) int {
	return x + y
}
