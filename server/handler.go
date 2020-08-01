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
	"time"
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
func (handler *Handler) TailLog(c *gin.Context) {
	// w.Write([]byte("hello"))
	conn, err := upgrader.Upgrade(c.Writer, c.Request, nil )
	if err != nil {
		log.Error("conn failed err: %v", err)
		//c.String(http.StatusBadRequest, "Bad request")
		return
	}
	go timeWriter(conn)

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
	for {
		time.Sleep(time.Second * 2)
		conn.WriteMessage(websocket.TextMessage, []byte(time.Now().Format("2006-01-02 15:04:05")))
	}
}

func Add(x, y int) int {
	return x + y
}
