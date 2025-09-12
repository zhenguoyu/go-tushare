package gotushare

import (
	"encoding/json"
	"fmt"
	"io"
	"strings"
	"time"

	"github.com/zhenguoyu/gotushare/internal/api"
	"github.com/zhenguoyu/gotushare/model"
)

const (
	APIURL     = "https://api.tushare.pro" // 接口地址
	REQTIMEOUT = 5 * time.Second           // 请求超时时间
)

type TushareResp struct {
	Code    int             `json:"code"`
	Message any             `json:"message"`
	Data    json.RawMessage `json:"data"`
}

type TushareReq struct {
	ApiName string            `json:"api_name"`
	Token   string            `json:"token"`
	Params  map[string]string `json:"params,omitempty"`
	Fields  string            `json:"fields,omitempty"`
}

type Tushare struct {
	token      string
	httpClient *api.HttpClient
}

func NewTushare(token string) *Tushare {
	t := &Tushare{token: token}
	t.httpClient = api.NewHttpClient(api.HttpClientConfig{
		Timeout: 5 * time.Second,
	})
	return t
}

func (t *Tushare) Post(apiName string, params map[string]string, fields ...string) (*TushareResp, error) {
	// 构造请求体

	reqBody := TushareReq{
		ApiName: "stock_basic",
		Token:   t.token,
		Params:  params,
	}
	if len(fields) > 0 {
		reqBody.Fields = strings.Join(fields, ",")
	} else {
		reqBody.Fields = strings.Join(model.FieldsStockBasic, ",")
	}
	reqBodyJson, err := json.Marshal(reqBody)
	if err != nil {
		return nil, err
	}
	resp, status, err := t.httpClient.Post(APIURL, io.NopCloser(strings.NewReader(string(reqBodyJson))), nil)
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

func (t *Tushare) GetStockBasic(params map[string]string, fields ...string) (*model.StockBasic, error) {

}
