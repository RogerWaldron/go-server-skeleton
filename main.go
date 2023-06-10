package main

import (
	"fmt"
	"time"
)

type Server struct {
	quitCh chan struct{} // 0 bytes
	msgCh chan string
}

func newServer() *Server {
	return &Server{
		quitCh: make(chan struct{}),
		msgCh: make(chan string, 128),
	}
}

func (s *Server) start() {
	fmt.Println("Starting Server")
	s.loop() // block
}

func (s *Server) quit() {
	s.quitCh <- struct{}{}
}

func (s *Server) sendMessage(msg string) {
	s.msgCh <- msg
}

func (s *Server) loop() {
	mainloop:
		for {
			select {
			case <-s.quitCh:
				fmt.Println("Quiting Server")
				break mainloop
			case msg := <-s.msgCh:
				s.handleMessage(msg)
			}
		}
		fmt.Println("Shutting Server Down Gracefully")
}

func (s *Server) handleMessage(msg string) {
	fmt.Println("Received message: ", msg)
}

func main() {
	server := newServer()

	go func() {
		time.Sleep(time.Second * 5)
		server.quit()
	}()

	server.start()
}