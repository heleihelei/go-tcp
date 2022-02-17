package main

import (
	"fmt"
	"net"
)

func main() {
	//监听
	listen, err := net.Listen("tcp", "127.0.0.1:8000") // 如果是本地的ip地址 则可以不用写 直接省略掉 即 "  : 8000"
	if err != nil {
		fmt.Println("err = ", err)
		return
	}

	defer listen.Close() //最后程序结束后才执行listen.Close( )进行关闭

	//阻塞等待用户链接
	conn, err := listen.Accept()
	if err != nil {
		fmt.Println("err = ", err)
		return
	}

	//接受用户的请求
	buf := make([]byte, 1024) //创建切片缓冲区，其大小为1024
	n, err1 := conn.Read(buf)
	if err1 != nil {
		fmt.Println("err = ", err)
		return
	}

	fmt.Println("buf = ", string(buf[:n])) //指定打印的切片，即读了多少就打印多少
	defer conn.Close()                     //关闭当前用户链接
}
