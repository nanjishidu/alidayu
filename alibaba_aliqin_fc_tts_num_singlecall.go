package alidayu

import (
	"encoding/json"
)

var ALIBABAALIQINFCTTSNUMSINGLECALL = "alibaba.aliqin.fc.tts.num.singlecall"

func (m *Alidayu) AlibabaAliqinFcTtsNumSinglecall(tts_code, called_num, called_show_num string,
	tts_param map[string]string, params ...string) (response AlibabaAliqinFcTtsNumSinglecallResponse, err error) {
	m.SetMethod(ALIBABAALIQINFCTTSNUMSINGLECALL)
	b, _ := json.Marshal(params)
	var extend string
	if len(params) >= 1 {
		extend = params[0]
	}
	resp, err := m.DoRequest(map[string]string{
		"tts_code":        tts_code,
		"called_num":      called_num,
		"called_show_num": called_show_num,
		"tts_param":       string(b),
		"extend":          extend,
	})
	err = json.Unmarshal(resp, &response)
	return
}

type AlibabaAliqinFcTtsNumSinglecallResponse struct {
	SuccessResponse SuccessResponse `json:"alibaba_aliqin_fc_tts_num_singlecall_response"`
	ErrorResponse   ErrorResponse   `json:"error_response"`
}
