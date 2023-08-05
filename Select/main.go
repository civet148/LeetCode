package main

import "fmt"

const (
	REQUEST_CHAN_MAX = 5
)

var request_chan = make(chan int, REQUEST_CHAN_MAX) //请求容量为5

func main() {

	for i := 0; i < REQUEST_CHAN_MAX+1; i++ {

		OnRequst(i)
	}
}

//使用select实现一个请求满了返回服务器忙的业务
func OnRequst(nReqId int) {

	select {
	case request_chan <- nReqId:
		fmt.Printf("Request %v ok \n", nReqId)
	default:
		fmt.Printf("Request %v error [channel is overflow]\n", nReqId)
	}
	return
}
