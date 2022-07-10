package game

import (
	"fmt"
	"golang.org/x/net/websocket"
)

type Player struct {
	UserId       int64
	ws           *websocket.Conn
}

func NewTestPlayer(ws *websocket.Conn, userId int64) *Player {
	player := new(Player)
	player.UserId = userId
	player.ws = ws
	return player
}


func (p *Player)SendNotice(msg []byte) {
	err :=websocket.Message.Send(p.ws, msg)
	if err!=nil{
		fmt.Println(err)
	}
}