package alidayu

import (
	"encoding/json"
	"strings"
)

var ALIBABAALIQINFCSMSNUMSEND = "alibaba.aliqin.fc.sms.num.send"

// sms_template_code:  短信模板ID，传入的模板必须是在阿里大于“管理中心-短信模板管理”中的可用模板。示例：SMS_585014
// sms_param:          短信模板变量，传参规则{"key":"value"}，key的名字须和申请模板中的变量名一致，多个变量之间以逗号隔开。
// rec_num:            短信接收号码。支持单个或多个手机号码，传入号码为11位手机号码，不能加0或+86。
// 		               群发短信需传入多个号码，以英文逗号分隔，一次调用最多传入200个号码。
func (m *Alidayu) AlibabaAliqinFcSmsNumSend(sms_template_code string, sms_param map[string]string,
	rec_num ...string) (response AlibabaAliqinFcSmsNumSendRespone, err error) {
	m.SetMethod(ALIBABAALIQINFCSMSNUMSEND)
	b, _ := json.Marshal(sms_param)
	resp, err := m.DoRequest(map[string]string{
		"sms_type":          "normal",
		"rec_num":           strings.Join(rec_num, ","),
		"sms_template_code": sms_template_code,
		"sms_param":         string(b),
	})
	err = json.Unmarshal(resp, &response)
	return
}

type AlibabaAliqinFcSmsNumSendRespone struct {
	SuccessResponse SuccessResponse `json:"alibaba_aliqin_fc_sms_num_send_response"`
	ErrorResponse   ErrorResponse   `json:"error_response"`
}
