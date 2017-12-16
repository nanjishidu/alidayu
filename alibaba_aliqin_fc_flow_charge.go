package alidayu

import (
	"encoding/json"
)

var ALIBABAALIQINFCFLOWCHARGE = "alibaba.aliqin.fc.flow.charge"

func (m *Alidayu) AlibabaAliqinFcFlowCharge(phone_num, grade, out_recharge_id string,
	params ...string) (response AlibabaAliqinFcFlowChargeRespone, err error) {
	m.SetMethod(ALIBABAALIQINFCFLOWCHARGE)
	var reason, scope, is_province string
	if len(params) >= 1 {
		reason = params[0]
	}
	if len(params) >= 2 {
		scope = params[1]
	}
	// 当scope=0时，is_province=true为指定分省通道。默认值为false
	if len(params) >= 3 {
		is_province = params[2]
	}
	resp, err := m.DoRequest(map[string]string{
		"phone_num":       phone_num,
		"grade":           grade,
		"out_recharge_id": out_recharge_id,
		"reason":          reason,
		"scope":           scope,
		"is_province":     is_province,
	})
	err = json.Unmarshal(resp, &response)
	return response, nil
}

type AlibabaAliqinFcFlowChargeRespone struct {
	Values        SuccessResponseValue `json:"alibaba_aliqin_fc_flow_charge_response"`
	ErrorResponse ErrorResponse        `json:"error_response"`
}
