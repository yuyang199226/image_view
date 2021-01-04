package main

import (
	//"flag"
	//"fmt"
	"github.com/sirupsen/logrus"
	//"imageview/config"
	//"imageview/database"
	//"imageview/server"
	//"imageview/utils"
	"imageview/cmd"
	//"os"
	//"os/signal"
	//"syscall"
)

func main() {
	logrus.Info("start...")
	cmd.Execute()
	//var (
	//	port        = flag.Int("port", 8088, "server listen port")
	//	downloadDir = flag.String("dir", "download", "download dir")
	//)
	//fmt.Println(utils.CurDatetime())
	//flag.Parse()
	////logger.SetFormatter(&logrus.JSONFormatter{})
	//s := make(chan os.Signal, 1)
	//signal.Notify(s, syscall.SIGUSR1)
	//go func() {
	//	for {
	//		<-s
	//		config.ReloadConfig()
	//		fmt.Println("Reloaded config")
	//	}
	//}()
    //db, err := database.NewMySQL()
    //if err != nil {
    //    panic(err)
    //}
    //a, err := database.Query(db)
    //if err != nil {
    //    panic(err)
    //}
    //fmt.Printf("%+v\n", a)
	//err = os.MkdirAll(*downloadDir, 0777)
	//if err != nil {
	//	fmt.Printf("create dir failed,err :%s", err)
	//	logrus.Fatal(err)
	//}
	//cfg := config.Config()
	//if cfg.Debug {
	//	fmt.Println("Debug mode not auth")
	//}
	//server.Start(*port)
	//fmt.Println("started...")

}
