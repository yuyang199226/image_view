package main

import (
    "fmt"
    "imageview/server"
    "imageview/config"
    "syscall"
    "os"
    "os/signal"
)

func main(){
    s := make(chan os.Signal, 1)
    signal.Notify(s, syscall.SIGUSR1)
    go func () {
        for {
            <-s
            config.ReloadConfig()
            fmt.Println("Reloaded config")
        }
    }()
    cfg := config.Config()
    fmt.Println(cfg.DB.Server)
    fmt.Println(cfg.Owner.Name)
    server.Start()
    fmt.Println("started...")

}
