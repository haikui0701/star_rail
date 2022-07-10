package main

import (
	"example.com/m/game"
	"fmt"
	"golang.org/x/net/websocket"
	"net"
	"net/http"
	"time"
)

func main() {
	http.Handle("/", websocket.Handler(WebsocketHandler))
	http.ListenAndServe("0.0.0.0:6666", nil)
}

func WebsocketHandler(ws *websocket.Conn) {
	var player  *game.Player

	for {
		var msg []byte
		ws.SetReadDeadline(time.Now().Add(3 * time.Second))
		err := websocket.Message.Receive(ws, &msg)
		fmt.Println(err)
		if err != nil {
			netErr, ok := err.(net.Error)
			if ok && netErr.Timeout() {
				continue
			}
			break
		}
		fmt.Println(string(msg))
		if player==nil{
			player=game.GetManagePlayer().PlayerLogin(ws)
		}
		if player!=nil{
			msgTest:=fmt.Sprintf("%d",player.UserId)
			game.GetManagePlayer().BoardCast([]byte(msgTest))
		}
	}
}
