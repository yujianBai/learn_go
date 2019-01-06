package main

import (
	"net/http"
	// "fmt"
	"github.com/gorilla/websocket"
)

var(
	upgrade = websocket.Upgrade{
		CheckOrgin: func(r *http.Request)bool{
			return true
		},
	}
)
func wsHandler(w http.ResponseWriter, r *http.Request){
	// fmt.Println("in wsHandler")
	// w.Write([]byte("hello"))
	var (
		conn *websocket.Conn
		err error
		data []byte
	)

	if conn, err = upgrade.Upgrade(w, r, nil); err != nil{
		return 	
	}

	for{
		if _, data, err = conn.ReadMessage(); err != nil{
			goto ERR
		}
		if err = conn.WriteMessage(websocket.TextMessage, data); err != nil{
			goto ERR
		}
	}
	ERR:
		conn.Close()
}


func main()  {
	http.HandleFunc("/ws", wsHandler)
	http.ListenAndServe("0.0.0.0:7777", nil)
}