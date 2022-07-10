package game

import (
	"fmt"
	"github.com/lxn/walk"
	. "github.com/lxn/walk/declarative"
	"github.com/lxn/win"
	"golang.org/x/net/websocket"
	"log"
	"net"
	"time"
)

const (
	TYPE_TEST = "test"
)

var starRailClient *StarRailClient = nil

type StarRailClient struct {
	mw   *walk.MainWindow
	test *MyPushButton
	ws   *websocket.Conn
}

type MyPushButton struct {
	*walk.PushButton
}

func GetStarRailClient() *StarRailClient {
	if starRailClient == nil {
		starRailClient = new(StarRailClient)
		if err := (MainWindow{
			AssignTo: &starRailClient.mw,
			Title:    "client:star_rail_client",
			Size:     Size{400, 300},
			Layout:   HBox{},
		}).Create(); err != nil {
			log.Fatal(err)
		}

		mpb, err := NewMyPushButton(starRailClient.mw)
		if err != nil {
			log.Fatal(err)
		}
		mpb.SetText(TYPE_TEST)
		starRailClient.test = mpb
	}

	return starRailClient
}

func NewMyPushButton(parent walk.Container) (*MyPushButton, error) {
	pb, err := walk.NewPushButton(parent)
	if err != nil {
		return nil, err
	}

	mpb := &MyPushButton{pb}

	if err := walk.InitWrapperWindow(mpb); err != nil {
		return nil, err
	}

	return mpb, nil
}

func WebsocketHandler(ws *websocket.Conn) {
	defer ws.Close()
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
		fmt.Println("收到消息开始")
		fmt.Println(string(msg))
		GetStarRailClient().test.SetText(string(msg))
		fmt.Println("收到消息结束")
	}
	return
}

func (s *StarRailClient) Start() {
	ws, err := websocket.Dial("ws://127.0.0.1:6666", "", "http://127.0.0.1:6666")
	if err != nil {
		log.Fatal(err)
	}
	s.ws = ws
	go WebsocketHandler(s.ws)
	starRailClient.mw.Run()
}

func (mpb *MyPushButton) WndProc(hwnd win.HWND, msg uint32, wParam, lParam uintptr) uintptr {
	switch msg {
	case win.WM_LBUTTONDOWN:
		log.Printf("%s: WM_LBUTTONDOWN", mpb.Text())
		mpb.LogicMsg()
	}

	return mpb.PushButton.WndProc(hwnd, msg, wParam, lParam)
}

func (mpb *MyPushButton) LogicMsg() {
	GetStarRailClient().Test()
}

func (s *StarRailClient) Test() {
	s.ws.Write([]byte("test"))
}
