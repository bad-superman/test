package mq

import (
	"fmt"
	"log"
	"testing"
	"time"

	zmq "github.com/pebbe/zmq4"
)

func ZeroMQClient() {
	zctx, _ := zmq.NewContext()
	s, _ := zctx.NewSocket(zmq.SUB)
	s.Connect("tcp://localhost:5555")
	s.SetSubscribe("topic1")
	for {
		msg, _ := s.Recv(0)
		log.Printf("Received %s\n", msg)
	}

}

func TestZeroMQ(t *testing.T) {
	go ZeroMQClient()
	go ZeroMQClient()

	zctx, _ := zmq.NewContext()

	// Socket to talk to server
	fmt.Printf("Connecting to the server...\n")
	s, _ := zctx.NewSocket(zmq.PUB)
	// s.SetSndhwm()
	s.Bind("tcp://*:5555")

	// Do 10 requests, waiting each time for a response
	for i := 0; i < 10; i++ {
		fmt.Printf("Sending request %d...\n", i)
		// s.Send("topic1", zmq.SNDMORE)
		s.Send("Hello", 0)
		time.Sleep(time.Second)
	}
}
