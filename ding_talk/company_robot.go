package ding_talk

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"errors"
	"net/http"
	"strconv"
	"time"
)

var (
	appSecret = "" // 机器人appSecret(不是加签secret)
)

func SetAppSecret(s string) {
	appSecret = s
}

type DingRobotMessage struct {
	AtUsers []struct {
		DingtalkID string `json:"dingtalkId"`
		StaffID    string `json:"staffId"`
	} `json:"atUsers"`
	ChatbotUserID     string `json:"chatbotUserId"`
	ConversationID    string `json:"conversationId"`
	ConversationTitle string `json:"conversationTitle"`
	ConversationType  string `json:"conversationType"`
	CreateAt          int    `json:"createAt"`
	MsgID             string `json:"msgId"`
	Msgtype           string `json:"msgtype"`
	SenderCorpID      string `json:"senderCorpId"`
	SenderID          string `json:"senderId"`
	SenderNick        string `json:"senderNick"`
	SenderStaffID     string `json:"senderStaffId"`
	Text              struct {
		Content string `json:"content"`
	} `json:"text"`
}

type DingRobotMsgHeader struct {
	TimeStamp string `json:"timestamp"`
	Sign      string `json:"sign"`
}

func Check(header http.Header) error {
	h := DingRobotMsgHeader{
		TimeStamp: header.Get("timestamp"),
		Sign:      header.Get("sign"),
	}
	if !VerifySign(&h) {
		return errors.New("sign not match")
	}
	return nil
}

func VerifySign(m *DingRobotMsgHeader) bool {
	t, err := strconv.Atoi(m.TimeStamp)
	if err != nil {
		return false
	}
	limit := time.Now().Add(-time.Hour).Unix() * 1000
	if t < int(limit) {
		return false
	}
	message := m.TimeStamp + "\n" + appSecret
	h := hmac.New(sha256.New, []byte(appSecret))
	_, _ = h.Write([]byte(message))
	return m.Sign == base64.StdEncoding.EncodeToString(h.Sum(nil))
}
