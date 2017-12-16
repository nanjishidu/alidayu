package alidayu

import (
	"encoding/json"
)

var AlibabaAliqinFcFlowQuery = "alibaba.aliqin.fc.flow.query"

func (m *Alidayu) AlibabaAliqinFcFlowQuery(out_id string) (response AlibabaAliqinFcFlowQueryRespone,
	err error) {
	m.SetMethod(AlibabaAliqinFcFlowQuery)
	resp, err := m.DoRequest(map[string]string{
		"out_id": out_id,
	})
	err = json.Unmarshal(resp, &response)
	return response, nil
}

type AlibabaAliqinFcFlowQueryRespone struct {
	Values        SuccessResponseValue `json:"alibaba_aliqin_fc_flow_query_response"`
	ErrorResponse ErrorResponse        `json:"error_response"`
}
