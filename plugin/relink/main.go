// Package relink 加/解密、链接转换
package relink

import (
	ctrl "github.com/FloatTech/zbpctrl"
	"github.com/FloatTech/zbputils/control"
	zero "github.com/wdvxdr1123/ZeroBot"
)

func init() {
	engine := control.AutoRegister(&ctrl.Options[*zero.Ctx]{
		DisableOnDefault:  false,
		Brief:             "加/解密、链接转换",
		Help:              "- (熊曰|[新]佛曰)加密",
		PrivateDataFolder: "relink",
	})
	engine.OnRegex(`(?ms)(.*)加密(.*)`).SetBlock(true).Handle(PcMoeEncode)
	engine.OnRegex(`(?ms)(.*)：(.*)`).SetBlock(true).Handle(PcMoeDecode)
}
