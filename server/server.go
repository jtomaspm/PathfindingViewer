package server

import (
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"golang.org/x/net/websocket"
)

type ServerConfig struct {
	address string
	port    string
}

type Server struct {
	config *ServerConfig
	cons   map[*websocket.Conn]bool
	eng    *gin.Engine
}

func NewServer(address string, port string) Server {
	c := ServerConfig{}
	c.address = address
	c.port = port

	return Server{
		config: &c,
		cons:   make(map[*websocket.Conn]bool),
		eng:    gin.Default(),
	}
}

func (s *Server) setupRoutes() {
	http.Handle("/ws", websocket.Handler(s.handleWS))

	s.eng.GET("/", func(c *gin.Context) {
		pb, err := os.ReadFile("./server/static/home.html")
		if err != nil {
			c.JSON(500, "Error getting resource")
			return
		}
		c.Data(200, "text/html", pb)
	})
}

func (s *Server) handleWS(ws *websocket.Conn) {
	log.Println("New incoming socket connection:", ws.RemoteAddr())
	for {
		data, err := os.ReadFile("./maps/1.txt")
		if err != nil {
			return
		}
		ws.Write(data)
		time.Sleep(500 * time.Millisecond)
	}
}

func (s *Server) Run() {
	s.setupRoutes()
	go http.ListenAndServe("0.0.0.0:3000", nil)
	s.eng.Run(s.config.address + ":" + s.config.port)
}
