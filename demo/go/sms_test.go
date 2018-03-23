package qcloudsms

import (
	"testing"
)

var (
	appid  = 1400076771
	appkey = "d68645c3af8aad3cf73941710ca70cdf"
)

var (
	phoneNumber1 = "13576666666"
	phoneNumber2 = "13576666666"
	phoneNumber3 = "13576666666"
)

func TestMydemo(t *testing.T) {
    t.Log("This is test...")
}

// 普通单发
// 注：发送的内容必须与模板内容一致，否则会提示内容与模板不匹配

func TestSmsSingleSend(t *testing.T) {
	sender := NewSmsSingleSender(appid, appkey)
	res, err := sender.Send(0, "86", phoneNumber1, "您好！欢迎您！", "", "")
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("%+v", res)
}


/*
// 指定模板单发
func TestSmsSingleSendWithParam(t *testing.T) {
	tmplId := 97868 // 欢迎短信模板ID号
	sender := NewSmsSingleSender(appid, appkey)
	params := []string{"李大锤"}
	res, err := sender.SendWithParam(0, "86", phoneNumber1, tmplId, params, "", "", "")
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("%+v", res)
}

// 普通群发
func TestSmsMultiSend(t *testing.T) {
	sender := NewSmsMultiSender(appid, appkey)
	res, err := sender.Send(0, "86", []string{phoneNumber1, phoneNumber2, phoneNumber3}, "你的验证码是9999", "", "")
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("%+v", res)
}

// 指定模板群发
// 假设短信模板 id 为 1，模板内容为：测试短信，{1}，{2}，{3}，上学。
func TestSmsMultiSendWithParam(t *testing.T) {
	tmplId := 7839
	sender := NewSmsMultiSender(appid, appkey)
	params := []string{"制定模板群发", "深圳", "小明"}
	res, err := sender.SendWithParam("86", []string{phoneNumber1, phoneNumber2, phoneNumber3}, tmplId, params, "", "", "")
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("%+v", res)
}

// 拉取短信回执和回复
func TestSmsStatusPull(t *testing.T) {
	puller := NewSmsStatusPuller(appid, appkey)
	callbackResult, err := puller.PullCallback(10)
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("%+v", callbackResult)
	replyResult, err := puller.PullReply(10)
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("%+v", replyResult)
}

// 语音验证码发送
func TestSmsVoiceVerifyCodeSend(t *testing.T) {
	sender := NewSmsVoiceVerifyCodeSender(appid, appkey)
	res, err := sender.Send("86", phoneNumber1, "9999", 2, "")
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("%+v", res)
}

// 发送通知内容
func TestSmsVoicePromptSend(t *testing.T) {
	sender := NewSmsVoicePromptSender(appid, appkey)
	res, err := sender.Send("86", phoneNumber1, 2, 2, "欢迎使用", "")
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("%+v", res)
}
*/
