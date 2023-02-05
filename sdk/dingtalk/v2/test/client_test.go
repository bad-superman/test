package test

import (
	"testing"

	v2 "github.com/bad-superman/test/sdk/dingtalk/v2"
)

func TestDtalk(t *testing.T) {
	manager := v2.New("7cb929f2787d27719e3c81970112ebcf2f3d036ce9ac236ea2373589e5deba16")
	//im := []string{"18600321498"}
	//talk := v2.NewText()
	//talk.Link.Content = "今日监控 , 是不一样的烟火"
	//
	//talk.At.AtMobiles=im
	//manager.SendMsg(nil,talk)
	mark := v2.NewMarkDown()
	mark.Markdown.Title = "那就不拖拉拉夫斯基"
	mark.Markdown.Text = "### 持仓信息\n #### long:11 short:0\n### 挂单信息\n #### ask pos_slide:long price:25620\n#### bid pos_slide:long price:23180\n"

	manager.SendMsg(nil, mark)
}
