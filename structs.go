package alidayu

type SuccessResponse struct {
	Result *SuccessResponseResult `json:"result"`
}
type SuccessResponseValue struct {
	Result *SuccessResponseResult `json:"value"`
}
type SuccessResponseResult struct {
	ErrCode string `json:"err_code"`
	Model   string `json:"model"`
	Success bool   `json:"success"`
	Msg     string `json:"msg"`
}
type ErrorResponse struct {
	Code    int    `json:"code"`
	Msg     string `json:"msg"`
	SubCode string `json:"sub_code"`
	SubMsg  string `json:"sub_msg"`
}
