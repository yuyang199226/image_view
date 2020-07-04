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
