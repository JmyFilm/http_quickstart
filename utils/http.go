package utils

import (
	"github.com/bytedance/sonic"
	"github.com/gofiber/fiber/v3/client"
	"time"
)

type Map map[string]string

var cc *client.Client

func init() {
	cc = client.New().SetTimeout(time.Second * 27).SetJSONMarshal(sonic.Marshal).SetJSONUnmarshal(sonic.Unmarshal)
}

func HTTP(method string, url string, header, param Map, body any) (respBody []byte, statusCode int, err error) {
	resp, err := cc.Custom(url, method, client.Config{
		Body:   body,
		Header: header,
		Param:  param,
	})
	if err != nil {
		return respBody, statusCode, err
	}

	return resp.Body(), resp.StatusCode(), nil
}
