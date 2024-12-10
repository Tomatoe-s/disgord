package main 
import (
"fmt"
"net/http"

"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize: 4096,
	WriteBufferSize: 4096,
	
}

//create connection

conn,err:= websocket.Dial("wss://gateway.discord.gg","",op.Origin)
fmt.Println(conn,err)