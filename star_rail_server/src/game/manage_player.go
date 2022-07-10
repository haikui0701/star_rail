package game

import (
	"golang.org/x/net/websocket"
	"sync"
)

var managePlayer *ManagePlayer

type ManagePlayer struct {
	Players map[int64]*Player
	Id      int64
	lock    *sync.RWMutex
}

func GetManagePlayer() *ManagePlayer {
	if managePlayer == nil {
		managePlayer = new(ManagePlayer)
		managePlayer.Players = make(map[int64]*Player)
		managePlayer.lock = new(sync.RWMutex)
	}
	return managePlayer
}

func (mp *ManagePlayer)PlayerLogin(ws *websocket.Conn) *Player {
	mp.lock.Lock()
	defer mp.lock.Unlock()

	mp.Id++
	playerInfo:=NewTestPlayer(ws,mp.Id)
	if playerInfo==nil{
		return nil
	}
	mp.Players[playerInfo.UserId]=playerInfo
	return playerInfo
}

func (mp *ManagePlayer)BoardCast(msg []byte) {
	mp.lock.RLock()
	defer mp.lock.RUnlock()

	for _,p:=range mp.Players{
		p.SendNotice(msg)
	}
}

