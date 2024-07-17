package main

import (
	"flag"
	"fmt"
	"github.com/hashicorp/consul/api"
	"go-blog/global"
	"go-blog/handler"
	"go-blog/initialize"
	"go-blog/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/health"
	"google.golang.org/grpc/health/grpc_health_v1"
	"net"
)

func main() {

	Ip := flag.String("ip", "0.0.0.0", "ip地址")
	Port := flag.Int("port", 50051, "端口号")

	//初始化
	initialize.InitLogger()
	initialize.InitConfig()
	initialize.InitDB()

	flag.Parse()
	server := grpc.NewServer()
	proto.RegisterUserServer(server, &handler.UserServer{})

	listen, err := net.Listen("tcp", fmt.Sprintf("%s:%d", *Ip, *Port))
	if err != nil {
		panic("服务启动失败")
	}
	//注册健康服务检查
	grpc_health_v1.RegisterHealthServer(server, health.NewServer())
	//服务注册
	cfg := api.DefaultConfig()
	cfg.Address = fmt.Sprintf("%s:%d", global.ServerConfig.ConsulConfig.Host, global.ServerConfig.ConsulConfig.Port)

	client, err := api.NewClient(cfg)
	if err != nil {
		panic(err)
	}
	check := &api.AgentServiceCheck{
		GRPC:                           fmt.Sprintf("192.168.44.15:50051"),
		Timeout:                        "5s",
		Interval:                       "5s",
		DeregisterCriticalServiceAfter: "15s",
	}

	//生成注册对象
	registration := new(api.AgentServiceRegistration)
	registration.Name = global.ServerConfig.Name
	registration.ID = global.ServerConfig.Name
	registration.Port = *Port
	registration.Tags = []string{"scx", "user", "srv"}
	registration.Address = "192.168.44.15"
	registration.Check = check

	err = client.Agent().ServiceRegister(registration)

	if err != nil {
		panic(err)
	}

	err = server.Serve(listen)
	if err != nil {
		panic("服务启动失败")
	}

}
