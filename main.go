package main

import (
	"flag"
	"fmt"
	"go-blog/handler"
	"go-blog/proto"
	"google.golang.org/grpc"
	"net"
)

func main() {

	Ip := flag.String("ip", "0.0.0.0", "ip地址")
	Port := flag.Int("port", 50051, "端口号")
	flag.Parse()
	server := grpc.NewServer()
	proto.RegisterUserServer(server, handler.UserServer{})
	listen, err := net.Listen("tcp", fmt.Sprintf("%s:%d", *Ip, *Port))
	if err != nil {
		panic("服务启动失败")
	}
	err = server.Serve(listen)
	if err != nil {
		panic("服务启动失败")
	}

}
