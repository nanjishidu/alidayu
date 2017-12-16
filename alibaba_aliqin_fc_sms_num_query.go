package alidayu

import (
	"encoding/json"
	"time"
)

var ALIBABAALIQINFCSMSNUMQUERY = "alibaba.aliqin.fc.sms.num.query"

// current_page: 分页参数，每页数量。最大值50
// page_size: 	 分页参数,页码
// query_date:   短信发送日期，支持近30天记录查询
// rec_num:      短信接收号码
// params:       biz_id 短信发送流水

func (m *Alidayu) AlibabaAliqinFcSmsNumQuery(rec_num string, query_date time.Time, current_page,
	page_size int64, params ...string) (response AlibabaAliqinFcSmsNumQueryRespone, err error) {
	m.SetMethod(ALIBABAALIQINFCSMSNUMQUERY)
	if page_size > 50 {
		page_size = 50
	}
	var biz_id string
	if len(params) >= 1 {
		biz_id = params[0]
	}
	resp, err := m.DoRequest(map[string]string{
		"rec_num":      rec_num,
		"query_date":   query_date.Format("20060102"),
		"current_page": GetInt64Str(current_page),
		"page_size":    GetInt64Str(page_size),
		"biz_id":       biz_id,
	})
	err = json.Unmarshal(resp, &response)
	return response, nil
}

type AlibabaAliqinFcSmsNumQueryRespone struct {
	AlibabaAliqinFcSmsNumQueryResponePageInfo AlibabaAliqinFcSmsNumQueryResponePageInfo `json:"alibaba_aliqin_fc_sms_num_query_response"`
	ErrorResponse                             ErrorResponse                             `json:"error_response"`
}
type AlibabaAliqinFcSmsNumQueryResponePageInfo struct {
	CurrentPage int64                                   `json:"current_page"`
	PageSize    int64                                   `json:"page_size"`
	TotalCount  int64                                   `json:"total_count"`
	TotalPage   int64                                   `json:"total_page"`
	Values      AlibabaAliqinFcSmsNumQueryResponeValues `json:"values"`
	RequestId   string                                  `json:"request_id"`
}
type AlibabaAliqinFcSmsNumQueryResponeValues struct {
	FcPartnerSmsDetailDto []FcPartnerSmsDetailDto `json:"fc_partner_sms_detail_dto"`
}
type FcPartnerSmsDetailDto struct {
	Extend           string `json:"extend"`            //公共回传参数
	RecNum           string `json:"rec_num"`           //短信接收号码
	ResultCode       string `json:"result_code"`       //短信错误码
	SmsCode          string `json:"sms_code"`          //模板编码
	SmsContent       string `json:"sms_content"`       //短信内容短信发送内容
	SmsReceiver_time string `json:"sms_receiver_time"` //2015-12-12 12:12:12短信接收时间
	SmsSendTime      string `json:"sms_send_time"`     //2015-12-12 12:12:12短信发送时间
	SmsStatus        int    `json:"sms_status"`        //发送状态 1：等待回执，2：发送失败，3：发送成功
}
