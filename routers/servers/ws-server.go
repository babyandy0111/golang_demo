package main

import (
	"io"
	"log"
	"net"
	"strconv"

	"google.golang.org/grpc"

	proto "golang_demo/proto/chat" // 自动生成的 proto代码
)

// Streamer 服务端
type Streamer struct{}

// BidStream 实现了 ChatServer 接口中定义的 BidStream 方法
func (s *Streamer) BidStream(stream proto.Chat_BidStreamServer) error {
	ctx := stream.Context()
	for {
		select {
		case <-ctx.Done():
			log.Println("收到客户端通过context发出的终止信号")
			return ctx.Err()
		default:
			// 接收从客户端发来的消息
			input, err := stream.Recv()
			if err == io.EOF {
				log.Println("客户端发送的数据流结束")
				return nil
			}
			if err != nil {
				log.Println("接收数据出错:", err)
				return err
			}

			// 如果接收正常，则根据接收到的 字符串 执行相应的指令
			switch input.Input {
			case "结束对话\n", "结束对话": //⚠️ 此处增加了匹配条件
				log.Println("收到'结束对话'指令")
				if err := stream.Send(&proto.Response{Output: "收到结束指令"}); err != nil {
					return err
				}
				// 收到结束指令时，通过 return nil 终止双向数据流
				return nil

			case "返回数据流\n", "返回数据流": //⚠️ 此处增加了匹配条件
				log.Println("收到'返回数据流'指令")
				// 收到 收到'返回数据流'指令， 连续返回 10 条数据
				for i := 0; i < 10; i++ {
					if err := stream.Send(&proto.Response{Output: "数据流 #" + strconv.Itoa(i)}); err != nil {
						return err
					}
				}

			default:
				// 缺省情况下， 返回 '服务端返回: ' + 输入信息
				log.Printf("[收到消息]: %s", input.Input)
				if err := stream.Send(&proto.Response{Output: "服务端返回: " + input.Input}); err != nil {
					return err
				}
			}
		}
	}
}

func init() {
	log.Println("启动服务端...")
	server := grpc.NewServer()

	// 注册 ChatServer
	proto.RegisterChatServer(server, &Streamer{})

	address, err := net.Listen("tcp", ":3333")
	if err != nil {
		log.Println(err)
	}

	if err := server.Serve(address); err != nil {
		log.Println(err)
	}
}
