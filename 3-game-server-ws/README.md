基于websocket, 实现了玩家加入房间，玩家心跳, 心跳超时, 发送消息，广播消息，发系统消息的功能

```
命令扩展：
go get XXX                    #go get XXX 下载并安装指定的包及其依赖项
go mod tidy                   #go mod tidy 确保go.mod文件中的依赖项与当前项目的代码匹配
go mod init game-server-ws    #go mod init 初始化一个新的模块
go mod graph                  #go mod graph 显示模块依赖图
go mod download               #go mod download 下载所有依赖项
go mod verify                 #go mod verify 验证模块的完整性
go mod edit -require=...      #go mod edit -require=... 添加新的依赖项
go mod edit -droprequire=...  #go mod edit -droprequire=... 删除依赖项
go mod edit -replace=...      #go mod edit -replace=... 替换依赖项
go mod edit -dropreplace=...  #go mod edit -dropreplace=... 删除替换项
go mod vendor                 #go mod vendor 将依赖项复制到vendor目录中
go mod edit -module=...       #go mod edit -module=... 更新模块路径
go mod edit -json             #go mod edit -json 以JSON格式编辑go.mod文件
go mod why -m                 #go mod why -m 显示模块的原因

```

