package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func main() {
	router := gin.Default()

	// Serve HTML page to trigger connection
	router.GET("/", func(c *gin.Context) {
		http.ServeFile(c.Writer, c.Request, "index.html")
	})

	// Handle WebSocket connections
	router.GET("/ws", func(c *gin.Context) {
		conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
		if err != nil {
			panic(err)
		}

		for {
			// Read message from client
			messageType, p, err := conn.ReadMessage()
			if err != nil {
				panic(err)
			}

			// Echo message back to client
			err = conn.WriteMessage(messageType, p)
			if err != nil {
				panic(err)
			}
		}
	})

	router.Run(":7000")
}
