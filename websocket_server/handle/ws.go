package handle

import (
	"flag"
	"github.com/gin-gonic/gin"
	"log"
)

var addr = flag.String("127.0.0.1", ":8080", "http websocket server address")

func Run() {
	flag.Parse()
	//信任客户端所有的代理
	_ = r.SetTrustedProxies(nil)
	hub := newHub()
	go hub.run()

	r.GET("ws", func(c *gin.Context) {

		serveWs(hub, c.Writer, c.Request)
	})
	err := r.Run(*addr)
	if err != nil {
		log.Panic("无法启动ws服务")
	}

}
