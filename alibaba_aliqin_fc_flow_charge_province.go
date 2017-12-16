package alidayu

import (
	"encoding/json"
)

var ALIBABAALIQINFCFLOWCHARGEPROVINCE = "alibaba.aliqin.fc.flow.charge.province"

func (m *Alidayu) AlibabaAliqinFcFlowChargeProvince(phone_num, grade, out_recharge_id string,
	params ...string) (response AlibabaAliqinFcFlowChargeProvinceRespone, err error) {
	m.SetMethod(ALIBABAALIQINFCFLOWCHARGEPROVINCE)
	var reason string
	if len(params) >= 1 {
		reason = params[0]
	}
	resp, err := m.DoRequest(map[string]string{
		"phone_num":       phone_num,
		"grade":           grade,
		"out_recharge_id": out_recharge_id,
		"reason":          reason,
	})
	err = json.Unmarshal(resp, &response)
	return response, nil
}

type AlibabaAliqinFcFlowChargeProvinceRespone struct {
	Values        SuccessResponseValue `json:"alibaba_aliqin_fc_flow_charge_province_response"`
	ErrorResponse ErrorResponse        `json:"error_response"`
}
