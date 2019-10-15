package main

import (
	"io"
	"log"
	"net"
	"strconv"

	"google.golang.org/grpc"

	proto "golang_demo/proto/chat" // 自动生成的 proto代码
)

// Streamer 服務端
type Streamer struct{}

// BidStream 實作了 ChatServer
func (s *Streamer) BidStream(stream proto.Chat_BidStreamServer) error {
	ctx := stream.Context()
	for {
		select {
		case <-ctx.Done():
			log.Println("client context 發出的終止訊號")
			return ctx.Err()
		default:
			// 接收client的訊息
			input, err := stream.Recv()
			if err == io.EOF {
				log.Println("client 發送的stream结束")
				return nil
			}
			if err != nil {
				log.Println("接收訊息出錯:", err)
				return err
			}

			// 如果接收正常，則根據收到的 字符串 執行相應邏輯
			switch input.Input {
			case "結束對話\n", "結束對話": //⚠️ 此处增加了匹配条件
				log.Println("收到'結束對話'指令")
				if err := stream.Send(&proto.Response{Output: "收到结束指令"}); err != nil {
					return err
				}
				// 收到结束指令时，通过 return nil 终止双向数据流
				return nil

			case "返回數據流\n", "返回數據流": //⚠️ 此处增加了匹配条件
				log.Println("收到'返回數據流'指令")
				// 收到 收到'返回数据流'指令， 连续返回 10 条数据
				for i := 0; i < 10; i++ {
					if err := stream.Send(&proto.Response{Output: "數據流 #" + strconv.Itoa(i)}); err != nil {
						return err
					}
				}

			default:
				log.Printf("[收到消息]: %s", input.Input)
				if err := stream.Send(&proto.Response{Output: "服務端返回: " + input.Input}); err != nil {
					return err
				}
			}
		}
	}
}

func init() {
	log.Println("啟動...")
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
