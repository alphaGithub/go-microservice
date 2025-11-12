package middlewares

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool { return true },
}

func SocketHandler(httpContext *gin.Context) {
	conn, err := upgrader.Upgrade(httpContext.Writer, httpContext.Request, nil)

	if err != nil {
		fmt.Println("[err] >>>>> Error in socket handler", err)
		return
	}
	defer conn.Close()

	for {
		mt, msg, err := conn.ReadMessage()
		fmt.Println("[socket] >>>> msg >>>>> ", mt, msg)
		if err != nil {
			break
		}
		// echo back
		var reply = "reply:hello-world"
		conn.WriteMessage(mt, []byte(reply))
	}
}
