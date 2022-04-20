package handle

import (
	"encoding/json"
	"fmt"
)

type CallFunctionsStrResponseData struct {
	FunctionName string      `json:"type"`
	Param        interface{} `json:"params"`
}

func DayMarket() {
	fmt.Println("侧额")
}

type messageHandle struct {
	c *Client
}

func MessageInit(c *Client) *messageHandle {

	return &messageHandle{
		c: c,
	}
}

type loginResponse struct {
	Type string `json:"type"`
	Code int    `json:"code"`
	Msg  string `json:"msg"`
}

func (handle messageHandle) login() {

	rp := loginResponse{
		Type: "login_result",
		Code: 1,
		Msg:  "登录成功",
	}
	marshal, err := json.Marshal(rp)

	if err != nil {
		return
	}

	handle.c.hub.broadcast <- marshal
}
