package main

import (
	"flag"
	"fmt"
	log "github.com/sirupsen/logrus"
	"imageview/config"
	"imageview/server"
	"imageview/utils"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	var (
		port        = flag.Int("port", 8088, "server listen port")
		downloadDir = flag.String("dir", "download", "download dir")
	)
	fmt.Println(utils.CurDatetime())
	flag.Parse()
	//logger.SetFormatter(&logrus.JSONFormatter{})
	s := make(chan os.Signal, 1)
	signal.Notify(s, syscall.SIGUSR1)
	go func() {
		for {
			<-s
			config.ReloadConfig()
			fmt.Println("Reloaded config")
		}
	}()
	err := os.MkdirAll(*downloadDir, 0777)
	if err != nil {
		fmt.Printf("create dir failed,err :%s", err)
	}
	cfg := config.Config()
	if cfg.Debug {
		fmt.Println("Debug mode not auth")

	} else {
		fmt.Println("Open auth")
	}
	var img = make([]string, 0)
	img = append(img, "ssss")
	fmt.Println(img)
	resp, err := Get("http://example.com/")
	if err != nil {
		log.Error("get heep://example.com failed", err)
	}
	fmt.Println(resp)
	log.Info("start ....")
	log.WithFields(log.Fields{
		"animal": "walrus",
	}).Info("A walrus appears")

	fmt.Println(cfg.DB.Server)
	fmt.Println(cfg.Owner.Name)
	server.Start(*port)
	fmt.Println("started...")

}
