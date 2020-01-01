package ding_talk

import (
	"testing"
)

var (
	accessToken = ""
	secret      = ""
)

func TestMessageTextSend(t *testing.T) {
	text := "无人唱彻大风歌"
	r := NewRobot(accessToken, secret)
	err := r.SendMessage(NewMessageBuilder(TypeText).Text(text).Build())
	if err != nil {
		t.Error(err)
	}
	t.Log("success!")
}

func TestMessageMarkdownSend(t *testing.T) {
	title := "标题"
	text := `
### 无题
> 相见时难别亦难  
  东风无力百花残
`
	r := NewRobot(accessToken, secret)
	err := r.SendMessage(NewMessageBuilder(TypeMarkdown).Markdown(title, text).Build())
	if err != nil {
		t.Error(err)
		return
	}
	t.Log("success!")
}
