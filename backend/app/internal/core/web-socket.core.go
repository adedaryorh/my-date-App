package core

import (
	"context"
	"fmt"
	"log"
	"runtime"
	"strings"
	"time"

	"backend.app/common/helpers"
	"backend.app/internal/models"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/gorilla/websocket"
)

var Clients = make(map[string]*Client)

type Client struct {
	conn   *websocket.Conn
	reason string
}

// getLoginQrCode generates the qr code to login
func (c *Core) getLoginQrCode(ctx context.Context, conn *websocket.Conn) {
	// generate unique id
	uuidString := uuid.New().String()
	newUuid := helpers.BytesToBase64([]byte(uuidString))

	if err := conn.WriteMessage(websocket.TextMessage, []byte("qr-code-made:"+newUuid)); err != nil {
		c.log.Error("error sending unique id to client: %v", err)
		return
	}

	Clients[uuidString] = &Client{
		conn:   conn,
		reason: "qr-code-login",
	}
	go cleanAfterTimeout(time.Now().Add(time.Second*3600), uuidString)
	logMemUsage()
}

// authenticate user
func (c *Core) qrCodeAuthenticateUser(ctx *gin.Context, conn *websocket.Conn, data string) error {
	// action:YUUTRR%^&*(&&)
	fmt.Println(data)
	var dto models.QrCodeLoginDto
	err := helpers.Base64ToStruct(&dto, data)

	// get connection
	webClient := Clients[dto.Id]
	fmt.Println(webClient)
	if webClient == nil {
		if err := conn.WriteMessage(websocket.TextMessage, []byte("login-failed")); err != nil {
			c.log.Error("error writing to websocket connection : %v", err)
			return err
		}
	}
	// authenticate token
	user, err := c.middleware.JwtRefreshTokenAuth(ctx, dto.Token, models.RedisKeys.AccessToken)
	if err != nil && webClient != nil {
		if err := webClient.conn.WriteMessage(websocket.TextMessage, []byte("login-failed")); err != nil {
			c.log.Error("error writing to websocket connection : %v", err)
			return err
		}
		if err := conn.WriteMessage(websocket.TextMessage, []byte("login-failed")); err != nil {
			c.log.Error("error writing to websocket connection : %v", err)
			return err
		}
	}

	if user != nil {
		result, err := c.generateTokens(ctx, user)
		if err != nil && webClient != nil {
			if err := webClient.conn.WriteMessage(websocket.TextMessage, []byte("login-failed")); err != nil {
				c.log.Error("error writing to websocket connection : %v", err)
				return err
			}
			if err := conn.WriteMessage(websocket.TextMessage, []byte("login-failed")); err != nil {
				c.log.Error("error writing to websocket connection : %v", err)
				return err
			}
		}
		// convert auth user to base64
		authUser, err := helpers.StructToBase64(result)
		if err != nil && webClient != nil {
			if err := webClient.conn.WriteMessage(websocket.TextMessage, []byte("login-failed")); err != nil {
				c.log.Error("error writing to websocket connection : %v", err)
				return err
			}
			if err := conn.WriteMessage(websocket.TextMessage, []byte("login-failed")); err != nil {
				c.log.Error("error writing to websocket connection : %v", err)
				return err
			}
		}
		if webClient != nil {
			if err := webClient.conn.WriteMessage(websocket.TextMessage, []byte("user-authenticated:"+authUser)); err != nil {
				c.log.Error("error writing to websocket connection : %v", err)
				return err
			}
		}
	}

	delete(Clients, dto.Id)
	return nil
}

func (c *Core) HandleWebsocketConnection(ctx *gin.Context, conn *websocket.Conn) {
	defer conn.Close()
	for {
		_, message, err := conn.ReadMessage()
		if err != nil {
			c.log.Error("error reading from websocket connection : %v", err)
			break
		}
		action, data := getWebSocketAction(message)
		if action == "get-login-qr-code" {
			c.getLoginQrCode(ctx, conn)
		}
		if action == "authenticate" {
			fmt.Println(message)
			c.qrCodeAuthenticateUser(ctx, conn, data)
		}
	}
}

func getWebSocketAction(message []byte) (string, string) {
	action := string(message)
	split := strings.Split(action, ":")
	size := len(split)
	if size > 0 && size < 2 {
		return split[0], "nil"
	}
	if size > 1 {
		return split[0], split[1]
	}
	return "nil", "nil"
}

func logMemUsage() {
	var m runtime.MemStats
	runtime.ReadMemStats(&m)

	log.Printf("Alloc = %v ", m.Alloc)

}
func cleanAfterTimeout(timeout time.Time, id string) {
	time.Sleep(time.Until(timeout))
	log.Println("cleanAfterTimeout:", id)
	if client, ok := Clients[id]; ok {
		client.conn.WriteMessage(websocket.TextMessage, []byte("TIMEOUT"))
		defer client.conn.Close()
		logMemUsage()
		delete(Clients, id)
		logMemUsage()
	}
	runtime.GC()
	logMemUsage()

}
