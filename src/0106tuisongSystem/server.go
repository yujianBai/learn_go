package main

import (
	"net/http"
	"fmt"
	"github.com/gorilla/websocket"
	impl "0106tuisongSystem/impl"
	"time"
)

var (
	upgrade = websocket.Upgrader{ 
		//允许跨域
		CheckOrigin: func(r *http.Request)bool{
			return true
		}, 
	}
)
func wsHandler(w http.ResponseWriter, r *http.Request){
	fmt.Printf("in wsHandler %v\n", w)
	// w.Write([]byte("hello"))
	var (
		wsConn *websocket.Conn
		err error
		data []byte
		conn *impl.Connection
	)

	if wsConn, err = upgrade.Upgrade(w, r, nil); err != nil{
		// goto ERR
		return 
	}

	if conn, err = impl.InitConnection(wsConn); err != nil{
		goto ERR
	}

	go func(){
		var(
			err error
		)
		for {
			if err = conn.WriteMessage([]byte("heartbeat")); err != nil{
				return 
			}
			time.Sleep(1 * time.Second)
		}

	}()

	for{
		if data, err = conn.ReadMessage(); err != nil{
			goto ERR
		}
		if err = conn.WriteMessage(data); err != nil{
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