package main

import (
    "fmt"
    "imageview/server"
    "imageview/config"
    "syscall"
    "os"
    "os/signal"
    "flag"
    "github.com/sirupsen/logrus"
)

func main(){
    var (
        port             = flag.Int("port", 8088, "server listen port")
        downloadDir      = flag.String("dir", "download", "download dir")
    
    )
    flag.Parse()
    logger := logrus.New()
    logger.SetFormatter(&logrus.TextFormatter{DisableColors: true, FullTimestamp: true})
    s := make(chan os.Signal, 1)
    signal.Notify(s, syscall.SIGUSR1)
    go func () {
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

    logger.Info("start ....")
    logger.WithFields(logrus.Fields{
    "animal": "walrus",
  }).Info("A walrus appears")

    fmt.Println(cfg.DB.Server)
    fmt.Println(cfg.Owner.Name)
    server.Start(*port)
    fmt.Println("started...")

}
