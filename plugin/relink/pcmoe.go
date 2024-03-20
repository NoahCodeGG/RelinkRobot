package relink

import (
	"github.com/FloatTech/floatbox/web"
	"github.com/sirupsen/logrus"
	zero "github.com/wdvxdr1123/ZeroBot"
	"github.com/wdvxdr1123/ZeroBot/message"
	"net/http"
	"net/url"
	"strings"
)

// PcMoeDecode 萌研社解密
func PcMoeDecode(ctx *zero.Ctx) {
	logrus.Infoln("[Relink] PcMoeDecode mode text: ", ctx.State["regex_matched"].([]string)[1])
	// 获取 Mode
	mode := PcMoeMode(ctx.State["regex_matched"].([]string)[1])
	logrus.Infoln("[Relink] PcMoeDecode mode: ", ctx.State["regex_matched"].([]string)[1])

	resp, err := pcMoeRequest(mode, PcMoeCodeDecode, ctx.State["regex_matched"].([]string)[2])
	if err != nil {
		ctx.SendChain(message.Text("ERROR: ", err))
		return
	}
	logrus.Infoln("[Relink] PcMoeDecode resp: ", resp)
	ctx.SendChain(message.Text(resp))
}

// PcMoeEncode 萌研社加密
func PcMoeEncode(ctx *zero.Ctx) {
	logrus.Infoln("[Relink] PcMoeEncode mode text: ", ctx.State["regex_matched"].([]string)[1])
	// 获取 Mode
	mode := PcMoeMode(ctx.State["regex_matched"].([]string)[1])
	logrus.Infoln("[Relink] PcMoeEncode mode: ", ctx.State["regex_matched"].([]string)[1])

	resp, err := pcMoeRequest(mode, PcMoeCodeEncode, ctx.State["regex_matched"].([]string)[2])
	if err != nil {
		ctx.SendChain(message.Text("ERROR: ", err))
		return
	}
	logrus.Infoln("[Relink] PcMoeEncode resp: ", resp)
	ctx.SendChain(message.Text(resp))
}

func pcMoeRequest(mode PcMoeMode, code PcMoeCode, txt string) (data string, err error) {
	// 获取 Mode 值
	modeValue, err := mode.value()
	if err != nil {
		return "", err
	}

	// 初始化请求参数
	urlValues := url.Values{
		"mode": {modeValue},
		"code": {string(code)},
		"txt":  {txt},
	}
	logrus.Infoln("[Relink] pcMoeRequest urlValues: ", urlValues.Encode())

	// 发送请求
	client := web.NewDefaultClient()
	resp, err := web.RequestDataWithHeaders(client, PcMoeApi, "POST", getPcMoeRequestHeaders, strings.NewReader(urlValues.Encode()))
	if err != nil {
		return "", err
	}

	return string(resp), nil
}

func getPcMoeRequestHeaders(req *http.Request) error {
	req.Header.Add("X-Token", "203B61D35068")
	req.Header.Add("X-Requested-With", "XMLHttpRequest")
	req.Header.Add("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/122.0.0.0 Safari/537.36")
	req.Header.Add("Referer", "http://hi.pcmoe.net/index.html")
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Add("Accept", "*/*")
	req.Header.Add("Host", "hi.pcmoe.net")
	req.Header.Add("Connection", "keep-alive")
	return nil
}
