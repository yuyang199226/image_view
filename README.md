go 包管理 
用go mod 不需要吧 目录放到 gopath 下面了
go mod init "NAME" 
上面的命令会生成 go.mod 文件  

如果有依赖不需要了， 执行 go mod tidy 可以去掉不需要的依赖module

go build  会生成可执行文件  go.sum 也会生成

