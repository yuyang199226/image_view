go 包管理 
用go mod 不需要吧 目录放到 gopath 下面了
go mod init "NAME" 
上面的命令会生成 go.mod 文件  

如果有依赖不需要了， 执行 go mod tidy 可以去掉不需要的依赖module

go build  会生成可执行文件  go.sum 也会生成
# test

切到 有后缀是 _test.go 的目录 执行 
`go test .`


编写可维护的代码
https://learnku.com/go/wikis/38430

1. 尽早return
2. 只在 main.main 或者 init 函数中使用log.Fatal
3. 
  
  ## kafka
  ### topic
  对于每个tpoic ，kafka 集群都会保留一个分区的日志
  ![pic](!https://kafka.apache.org/25/images/log_anatomy.png)
  
  每个分区都是有序的，在分区里的记录被分配了一个序列号 offset, 在这个分区是惟一的。kafka 使用持久保留published record..
  无论这个record 有没有被消费。使用配置 去设置保留时间。比如 设置两天，两天内the record 是可用的。两天后就会被kafka 删掉来释放空间。
  一个topic 可以有多个分区。分区的目的1 是 扩容。而是充当 并行单元。
  每个分区都有 一个leader server 和  0 或 多个 follower server .leader server 处理读写请求，follower 
  负责 复制
  
  ### producer
  负责将数据发布到 topic
  record
  有key ,value, timestamp ， offsetDetla,组成
  
  
  curl --include \
       --no-buffer \
       --header "Connection: Upgrade" \
       --header "Upgrade: websocket" \
       --header "Host: 127.0.0.1:8088" \
       --header "Origin: http://127.0.0.1:8088" \
       --header "Sec-WebSocket-Key: SGVsbG8sIHdvcmxkIQ==" \
       --header "Sec-WebSocket-Version: 13" \
       http://127.0.0.1:8088/v1/ws
  
  