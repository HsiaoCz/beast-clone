package main

import (
	"fmt"
	"io"
	"log"
	"log/slog"
	"net/http"
	"os"
	"time"

	"github.com/joho/godotenv"
	"golang.org/x/net/websocket"
)

type Server struct {
	conns map[*websocket.Conn]bool
}

func NewServer() *Server {
	return &Server{
		conns: make(map[*websocket.Conn]bool),
	}
}

func (s *Server) handleWS(ws *websocket.Conn) {
	slog.Info("new incoming connection from client", "the remote address", ws.RemoteAddr())

	s.conns[ws] = true

	s.readLoop(ws)
}

func (s *Server) readLoop(ws *websocket.Conn) {
	buf := make([]byte, 1024)
	for {
		n, err := ws.Read(buf)
		if err != nil {
			if err == io.EOF {
				break
			}
			slog.Error("read error", "error message", err)
			continue
		}
		msg := buf[:n]
		s.broadcast(msg)
	}
}

func (s *Server) broadcast(b []byte) {
	for ws := range s.conns {
		go func(ws *websocket.Conn) {
			if _, err := ws.Write(b); err != nil {
				slog.Error("websocket write error", "error message", err)
			}
		}(ws)
	}
}

func (s *Server) handleWsOrderbook(ws *websocket.Conn) {
	slog.Info("new incoming connection from client to orderbook feed", "remote address", ws.RemoteAddr())

	for {
		payload := fmt.Sprintf("orderbook data ---> %d\n", time.Now().UnixNano())
		ws.Write([]byte(payload))
		time.Sleep(time.Second)
	}
}

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatal(err)
	}
	server := NewServer()
	http.Handle("/ws", websocket.Handler(server.handleWS))
	http.Handle("/orderbookfeed", websocket.Handler(server.handleWsOrderbook))
	slog.Info("the http server is ruuning", "port", os.Getenv("PORT"))
	http.ListenAndServe(os.Getenv("PORT"), nil)
}
