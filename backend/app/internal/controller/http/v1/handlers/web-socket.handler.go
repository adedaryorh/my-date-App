package handlers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

// getUpGrader is used to upgrade HTTP connections to WebSocket connections.
func getUpGrader() websocket.Upgrader {
	return websocket.Upgrader{
		CheckOrigin: func(r *http.Request) bool {
			// origin := r.Header.Get("Origin")
			// return origin == "<http://yourdomain.com>"
			return true
		},
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
	}
}

// @Tags WebSocket
// @Summary Connects the client to the server via web sockets
// @Schemes
// @Description Connects the client to the server via web sockets
// @Accept json
// @Produce json
// @Success 200 {object} dtos.ResponseObject{} "desc"
// @Router /v1/ws [get]
func (h *Handler) WebSocketHandler(c *gin.Context) {
	upgrader := getUpGrader()
	// Upgrade the HTTP connection to a WebSocket connection
	conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		fmt.Println("Error upgrading:", err)
		return
	}
	// Listen for incoming messages
	go h.core.HandleWebsocketConnection(c, conn)
}
