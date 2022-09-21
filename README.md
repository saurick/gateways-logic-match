# 游戏匹配网关

## 目标
实现网关匹配系统，进行玩家鉴权，向服务器进行匹配请求调用，使用redis进行事件发布订阅，使用protobuf进行消息转发 

## 启动redis容器服务
执行命令`docker-compose up -d`启动相关容器服务

## 构建和运行
执行命令`make build-gateways-logic-match`生成bin目录和可执行文件，终端输入`./bin/./gateways-logic-match`来运行程序（注意端口是否冲突，默认端口是`30052`）

## 测试
要同时启动[匹配服务器](https://git.sysfun.cn/projects/CB/repos/services-match/browse)，本地安装redis-cli
1. 先clone[模拟游戏客户端](https://git.sysfun.cn/projects/CB/repos/testclients-csharp-match/browse)
```
$ git clone ssh://git@git.sysfun.cn:7999/cb/testclients-csharp-match.git
$ cd testclients-csharp-match
```
2. 开一个终端模拟1个玩家请求匹配
```
$ redis-cli -p 63790 set playerToken1 1
$ dotnet run Program.cs match playerToken1 30052
```
3. 开一个终端模拟玩家11个玩家请求匹配
```

$ cat > match.sh <<EOF
# bash/bin

max=11

for i in \`seq 1 \$max\`
do
    redis-cli -p 63790 set playerToken\$i \$i
done

for i in \`seq 1 \$max\`
do
    player="playerToken\$i"
    dotnet run Program.cs "match" \$player 30052 &
    sleep 1
done
EOF

$ chmod 700 match.sh

$ ./match.sh
```
4. 模拟1个玩家取消匹配
```
$ redis-cli -p 63790 set playerToken1 1
$ dotnet run Program.cs cancel playerToken1 30052
```
5. 模拟客户端网络中断

在终端执行命令`dotnet run Program.cs match playerToken1 30052`的之后，按ctrl + c退出程序

## TODO
* 单元测试
* 性能测试
