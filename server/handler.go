package server

import (
	"fmt"
	//"github.com/fsnotify/fsnotify"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"github.com/hpcloud/tail"
	log "github.com/sirupsen/logrus"
	"imageview/config"
	"io"
	"net/http"
	"os"
	//"time"
)

func init() {
	log.SetFormatter(&log.JSONFormatter{})

}
type FileWatcher struct {
	FileName string
	Size int64
	PrevSize int64
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

func (handler *Handler) TailLog(c *gin.Context) {
	conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		log.Error("conn failed err: %v", err)
		return
	}
	timeWriter(conn)
	go func(conn *websocket.Conn) {
		for {
			msgType, msg, err := conn.ReadMessage()
			if err != nil {
				log.Error("logtail err: %v", err)
				return
			}
			err = conn.WriteMessage(msgType, msg)
			if err != nil {
				return
			}
		}
	}(conn)
}

func timeWriter(conn *websocket.Conn) {
	//go func() {
		t, err := tail.TailFile("/tmp/be.log", tail.Config{Follow: true})
		if err != nil {
			return
		}
		for line := range t.Lines {
			conn.WriteMessage(websocket.TextMessage, []byte(line.Text+"\n"))
			//fmt.Println(line.Text)
		}

	//}()
}

func Add(x, y int) int {
	return x + y
}
