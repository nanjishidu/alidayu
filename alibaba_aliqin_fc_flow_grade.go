package alidayu

import (
	"encoding/json"
)

var ALIBABAALIQINFCFLOWGRADE = "alibaba.aliqin.fc.flow.grade"

func (m *Alidayu) AlibabaAliqinFcFlowGrade(phone_num string) (response AlibabaAliqinFcFlowGradeRespone,
	err error) {
	m.SetMethod(ALIBABAALIQINFCFLOWGRADE)
	resp, err := m.DoRequest(map[string]string{
		"phone_num": phone_num,
	})
	err = json.Unmarshal(resp, &response)
	return response, nil
}

type AlibabaAliqinFcFlowGradeRespone struct {
	Values        SuccessResponseValue `json:"alibaba_aliqin_fc_flow_grade_response"`
	ErrorResponse ErrorResponse        `json:"error_response"`
}
