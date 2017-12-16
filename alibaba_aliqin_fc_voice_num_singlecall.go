package alidayu

import (
	"encoding/json"
)

var ALIBABAALIQINFCVOICENUMSINGLECALL = "alibaba.aliqin.fc.voice.num.singlecall"

func (m *Alidayu) AlibabaAliqinFcVoiceNumSinglecall(voice_code, called_num, called_show_num string,
	params ...string) (response AlibabaAliqinFcVoiceNumSinglecallRespone, err error) {
	m.SetMethod(ALIBABAALIQINFCVOICENUMSINGLECALL)
	var extend string
	if len(params) >= 1 {
		extend = params[0]
	}
	resp, err := m.DoRequest(map[string]string{
		"voice_code":      voice_code,
		"called_num":      called_num,
		"called_show_num": called_show_num,
		"extend":          extend,
	})
	err = json.Unmarshal(resp, &response)
	return
}

type AlibabaAliqinFcVoiceNumSinglecallRespone struct {
	SuccessResponse SuccessResponse `json:"alibaba_aliqin_fc_voice_num_singlecall_response"`
	ErrorResponse   ErrorResponse   `json:"error_response"`
}
