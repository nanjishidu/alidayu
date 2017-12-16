package alidayu

import (
	"encoding/json"
)

var ALIBABAALIQINFCVOICENUMDOUBLECALL = "alibaba.aliqin.fc.voice.num.doublecall"

func (m *Alidayu) AlibabaAliqinFcVoiceNumDoublecall(caller_num, caller_show_num, called_num, called_show_num string,
	params ...string) (response AlibabaAliqinFcVoiceNumDoublecallRespone, err error) {
	m.SetMethod(ALIBABAALIQINFCVOICENUMDOUBLECALL)
	//params[0] session_time_out
	//params[1] extend
	var session_time_out, extend string
	if len(params) >= 1 {
		session_time_out = params[0]
	}
	if len(params) >= 2 {
		extend = params[1]
	}
	resp, err := m.DoRequest(map[string]string{
		"caller_num":       caller_num,
		"caller_show_num":  caller_show_num,
		"called_num":       called_num,
		"called_show_num":  called_show_num,
		"session_time_out": session_time_out,
		"extend":           extend,
	})
	err = json.Unmarshal(resp, &response)
	return
}

type AlibabaAliqinFcVoiceNumDoublecallRespone struct {
	SuccessResponse SuccessResponse `json:"alibaba_aliqin_fc_voice_num_doublecall_response"`
	ErrorResponse   ErrorResponse   `json:"error_response"`
}
