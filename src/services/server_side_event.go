package services

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

var clients []chan string

func SSEHandler(c *gin.Context) {
	flusher, ok := c.Writer.(http.Flusher)
	if !ok {
		http.Error(c.Writer, "Streaming unsupported!", http.StatusInternalServerError)
		return
	}

	messageChan := make(chan string)
	clients = append(clients, messageChan)

	c.Writer.Header().Set("Content-Type", "text/event-stream")
	c.Writer.Header().Set("Cache-Control", "no-cache")
	c.Writer.Header().Set("Connection", "keep-alive")

	defer func() {
		for i, client := range clients {
			if client == messageChan {
				clients = append(clients[:i], clients[i+1:]...)
				break
			}
		}
	}()

	for {
		select {
		case msg := <-messageChan:
			fmt.Fprintf(c.Writer, "data: %s\n\n", msg)
			flusher.Flush()
		case <-c.Request.Context().Done():
			return
		}
	}
}

func BroadcastMessage(message string) {
	for _, client := range clients {
		client <- message
	}
}
