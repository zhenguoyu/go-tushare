package gotushare

import (
	"bytes"
	"encoding/json"
	"fmt"
	"strings"
	"time"

	"github.com/zhenguoyu/gotushare/internal/api"
)

const (
	APIURL      = "http://api.tushare.pro" // 接口地址
	REQTIMEOUT  = 5 * time.Second          // 请求超时时间
	ContentType = "application/json"
)

type TushareResp struct {
	Code    int             `json:"code,omitempty"`
	Message any             `json:"message,omitempty"`
	Data    json.RawMessage `json:"data,omitempty"`
}

type TushareReqData struct {
	ApiName string      `json:"api_name"`
	Token   string      `json:"token"`
	Params  interface{} `json:"params,omitempty"` // 接口参数，如daily接口中start_date和end_date, 大部分是map[string]string
	Fields  string      `json:"fields,omitempty"` // 逗号分隔的字段列表
}

type Tushare struct {
	token      string
	apiURL     string
	httpClient *api.HttpClient
}

func NewTushare(token string) *Tushare {
	return &Tushare{
		token:      token,
		apiURL:     APIURL,
		httpClient: api.NewHttpClientWithOptions(api.WithDialTimeout(5 * time.Second)),
	}
}

func (t *Tushare) RequestApi(apiName string, params map[string]string, fields ...string) (*TushareResp, error) {
	reqBody := TushareReqData{
		ApiName: apiName,
		Token:   t.token,
		Params:  params,
	}
	if len(fields) > 0 {
		reqBody.Fields = strings.Join(fields, ",")
	}
	reqBodyJson, err := json.Marshal(reqBody)
	if err != nil {
		return nil, err
	}
	resp, status, err := t.httpClient.Post(t.apiURL, bytes.NewBuffer(reqBodyJson), nil)
	if err != nil {
		return nil, err
	}
	if status != 200 {
		return nil, fmt.Errorf("请求失败: %s, status:%d", string(resp), status)
	}
	var tushareResp TushareResp
	err = json.Unmarshal(resp, &tushareResp)
	if err != nil {
		return nil, err
	}
	return &tushareResp, nil
}
