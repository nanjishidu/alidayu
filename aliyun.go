package alidayu

import (
	"crypto/tls"
	"errors"
	"fmt"
	"io/ioutil"
	"net"
	"net/http"
	"net/url"
	"sort"
	"strings"
	"time"
)

const (
	httpURL          string = "http://gw.api.taobao.com/router/rest"
	httpsURL         string = "https://eco.taobao.com/router/rest"
	ConnectTimeout          = 60 * time.Second
	ReadWriteTimeout        = 60 * time.Second
	Sign_Method_MD5         = "md5"
	Sign_Method_HDMC        = "hdmc"
)

//公共参数
type Alidayu struct {
	Method       string
	AppKey       string
	AppSecret    string
	TargetAppKey string
	SignMethod   string
	Sign         string
	Session      string
	Timestamp    string
	Format       string
	V            string
	PartnerId    string
	Simplify     string
	//
	SmsFreeSignName string
	//定义URI参数
	HttpsEnable bool
}

func NewAlidayu(appkey, appsecret string, params ...string) *Alidayu {
	var sms_free_sign_name string
	if len(params) > 0 {
		sms_free_sign_name = params[0]
	}
	return &Alidayu{
		AppKey:          appkey,
		AppSecret:       appsecret,
		SmsFreeSignName: sms_free_sign_name,
	}
}
func (m *Alidayu) SetMethod(method string) {
	m.Method = method
}
func (m *Alidayu) SetHttpsEnable(httpEnable bool) {
	m.HttpsEnable = httpEnable
}
func GetRequestUri(httpsEnable bool) (uri string) {
	if httpsEnable {
		return httpsURL
	}
	return httpURL
}
func (m *Alidayu) DoRequest(params map[string]string) ([]byte, error) {
	if m.Method == "" || m.AppKey == "" || m.AppSecret == "" {
		return nil, errors.New("param is invalid")
	}
	if m.SmsFreeSignName != "" {
		params["sms_free_sign_name"] = m.SmsFreeSignName
	}
	if m.Session != "" {
		//todo
	}
	params["method"] = m.Method
	params["app_key"] = m.AppKey

	//只允许json
	params["format"] = "json"

	var signMethod string
	switch m.SignMethod {
	case "", Sign_Method_MD5:
		signMethod = Sign_Method_MD5
	case Sign_Method_HDMC:
		signMethod = Sign_Method_HDMC
	default:
		return nil, fmt.Errorf("unsupported request sign_method: %s", m.SignMethod)

	}
	params["sign_method"] = signMethod
	if m.Timestamp == "" {
		m.Timestamp = time.Now().Format("2006-01-02 15:04:05")
	}
	params["timestamp"] = m.Timestamp
	if m.V == "" {
		m.V = "2.0"
	}
	params["v"] = m.V

	var keys []string
	for k := range params {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	var (
		signString string
		uriParam   = url.Values{}
	)

	for _, k := range keys {
		signString += k + params[k]
		uriParam.Set(k, params[k])
	}
	var sign string
	switch signMethod {
	case "md5":
		sign = Md5(m.AppSecret + signString + m.AppSecret)
	case "hmac":
		sign = Hmac(signString, m.AppSecret)
	}
	// sign => 字符串大写
	uriParam.Set("sign", strings.ToUpper(sign))

	rbody, size := ioutil.NopCloser(strings.NewReader(uriParam.Encode())), int64(len(uriParam.Encode()))

	req, err := http.NewRequest("POST", GetRequestUri(m.HttpsEnable), rbody)
	if err != nil {
		return nil, err
	}
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	req.ContentLength = size

	client := &http.Client{
		Transport: &http.Transport{
			Dial: func(netw, addr string) (net.Conn, error) {
				c, err := net.DialTimeout(netw, addr, ReadWriteTimeout)
				if err != nil {
					return nil, err
				}
				c.SetDeadline(time.Now().Add(ConnectTimeout))
				return c, nil
			},
			ResponseHeaderTimeout: ConnectTimeout,
			DisableKeepAlives:     true,
			TLSClientConfig:       &tls.Config{InsecureSkipVerify: true},
		},
	}

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	return ioutil.ReadAll(resp.Body)
}
