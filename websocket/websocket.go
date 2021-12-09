package websocket

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/websocket"
	"net/http"
	"time"
)

type Person struct {
	Name string `json:"name"` // 태그설정 : "name"
	Age  int    `json:"age"`  // 태그설정 : "age"
}

var Wsupgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin:     func(*http.Request) bool { return true },
}

func Wshandler(w http.ResponseWriter, r *http.Request) {
	conn, err := Wsupgrader.Upgrade(w, r, nil)
	if err != nil {
		fmt.Println("Failed to set websocket upgrade: %+v", err)
		return
	}

	p := Person{"Jason", 26}
	byteArray, err := json.Marshal(p) // p 데이터 json데이터로 변환
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(p) // {Jason 26}
	fmt.Println(string(byteArray))

	for {
		time.Sleep(time.Second * 2)
		conn.WriteMessage(1, byteArray)
	}

	/*for {
		t, msg, err := conn.ReadMessage()
		if err != nil {
			break
		}
		conn.WriteMessage(t, msg)
	}*/
}
