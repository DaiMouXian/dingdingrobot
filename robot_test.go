package dingdingrobot

import "testing"

//https://oapi.dingtalk.com/robot/send?access_token=10459b225f5b25538bb2ed53beae8722ec74c318ee125f3cda6798f8fa7c78b6
var dingToken = "66a35cc4233fb4c4f9eddf138ac4bf16fd8ddaa18c4b6e7bea63b2d9b20b59f7"

func TestMarkDownMsg(t *testing.T) {
	dingTalkCli, err := NewDingRobotWithSecret(dingToken, "SECcfe46279845a07444d9c7ff7a26e0b2b3e2275830b4a0d458b0cd11cd7356c50")
	if err != nil {
		t.Errorf("TestMarkDownMsg expected be nil, but %v got", err)
	}
	//err := dingTalkCli.SendMarkDownMessage("Markdown title", "### Link test\n --- \n- <font color=#ff0000 size=6>红色文字</font> \n - content2.", "+18975759665")
	err = dingTalkCli.SendTextMessage("Markdown title", "*")
	if err != nil {
		t.Errorf("TestMarkDownMsg expected be nil, but %v got", err)
	}
}
