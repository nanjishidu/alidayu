package alidayu

import "testing"

var (
	m                 = NewAlidayu("", "", "")
	sms_template_code = ""
	rec_num           = ""
	sms_param         = map[string]string{
		"name":    "测试用户",
		"product": "",
		"code":    "",
	}
)

func TestAlibabaAliqinFcSmsNumSend(t *testing.T) {
	m.SetHttpsEnable(true)
	resp, err := m.AlibabaAliqinFcSmsNumSend(sms_template_code, sms_param, rec_num)
	if err != nil {
		t.Fatal(err)
	}
	if resp.ErrorResponse.Msg != "" {
		t.Fatal(resp.ErrorResponse)
	}
}
