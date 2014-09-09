package main

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"net"
	"time"
)

func main() {
	// 链接服务器
	conn, err := net.Dial("tcp", "127.0.0.1:6000")
	if err != nil {
		fmt.Printf("Dial error: %s\n", err)
		return
	}

	// 客户端信息
	fmt.Printf("Client: %s\n", conn.LocalAddr())

	// 消息缓冲
	msgbuf := bytes.NewBuffer(make([]byte, 0, 1024))

	// 消息内容
	message := []byte("我是utf-8的消息")
	// 消息长度
	messageLen := uint32(len(message))
	// 消息总长度
	mlen := 4 + len(message)

	// 写入5条消息
	for i := 0; i < 10; i++ {
		binary.Write(msgbuf, binary.LittleEndian, messageLen)
		msgbuf.Write(message)
	}

	// 单包发送一条消息
	conn.Write(msgbuf.Next(mlen))
	time.Sleep(time.Second)

	// 单包发送三条消息
	conn.Write(msgbuf.Next(mlen * 3))
	time.Sleep(time.Second)

	// 发送不完整的消息头
	conn.Write(msgbuf.Next(2))
	time.Sleep(time.Second)
	// 发送消息剩下部分
	conn.Write(msgbuf.Next(mlen - 2))
	time.Sleep(time.Second)

	// 发送不完整的消息体
	conn.Write(msgbuf.Next(mlen - 6))
	time.Sleep(time.Second)
	// 发送消息剩下部分
	conn.Write(msgbuf.Next(6))
	time.Sleep(time.Second)

	// 多段发送
	conn.Write(msgbuf.Next(mlen + 2))
	time.Sleep(time.Second)
	conn.Write(msgbuf.Next(-2 + mlen - 8))
	time.Sleep(time.Second)
	conn.Write(msgbuf.Next(8 + 1))
	time.Sleep(time.Second)
	conn.Write(msgbuf.Next(-1 + mlen + mlen))
	time.Sleep(time.Second)

	// 关闭链接
	conn.Close()
}
