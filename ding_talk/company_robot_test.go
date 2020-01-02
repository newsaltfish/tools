package ding_talk

import (
	"encoding/json"
	"errors"
	"net/http"
	"testing"

	"github.com/labstack/echo"
)

func MiddlewareCheck(next echo.HandlerFunc) echo.HandlerFunc {
	return func(ctx echo.Context) error {
		h := DingRobotMsgHeader{
			TimeStamp: ctx.Request().Header.Get("timestamp"),
			Sign:      ctx.Request().Header.Get("sign"),
		}
		if !VerifySign(&h, appSecret) {
			return errors.New("sign not match")
		}
		return next(ctx)
	}
}
func HandleDingMessage(ctx echo.Context) error {
	// 解析消息
	m := DingRobotMessage{}
	err := ctx.Bind(&m)
	if err != nil {
		return err
	}
	// 回复消息
	replyText := ""
	msg := NewMessageBuilder(TypeText).Text(replyText).Build()
	bd, err := json.Marshal(msg)
	if err != nil {
		return err
	}
	ctx.Response().WriteHeader(http.StatusOK)
	_, _ = ctx.Response().Write(bd)
	return nil
}

func TestSeverDingRobot(t *testing.T) {
	SetAppSecret("")
	e := echo.New()
	dGroup := e.Group("/dingding")
	dGroup.Use(MiddlewareCheck)
	dGroup.GET("/callback", func(ctx echo.Context) error {
		t.Log("for http check")
		return nil
	})
	dGroup.POST("/callback", HandleDingMessage, MiddlewareCheck)
	err := e.Start(":12306")
	if err != nil {
		t.Error(err)
	}
}
