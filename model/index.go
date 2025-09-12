// index.go 指数接口相关数据结构
package model

import (
	json "encoding/json"
)

// 指数基本信息 api name
const ApiIndexBasic string = "index_basic"

// 指数日线行情 api name
const ApiIndexDaily string = "index_daily"

// 指数周线行情 api name
const ApiIndexWeekly string = "index_weekly"

// 指数月线行情 api name
const ApiIndexMonthly string = "index_monthly"

// 指数成分和权重 api name
const ApiIndexWeight string = "index_weight"

// 大盘指数每日指标 api name
const ApiIndexDailyBasic string = "index_dailybasic"

// 申万行业分类 api name
const ApiIndexClassify string = "index_classify"

// 申万行业成分 api name
const ApiIndexMember string = "index_member"

// 沪深市场每日交易统计 api name
const ApiDailyInfo string = "daily_info"

// 深圳市场每日交易情况 api name
const ApiSzDailyInfo string = "sz_daily_info"

// 同花顺概念和行业指数 api name
const ApiThsIndex string = "ths_index"

// 同花顺概念和行业指数行情 api name
const ApiThsDaily string = "ths_daily"

// 同花顺概念和行业指数成分 api name
const ApiThsMember string = "ths_member"

// 中信行业指数行情 api name
const ApiCiDaily string = "ci_daily"

// 国际主要指数 api name
const ApiIndexGlobal string = "index_global"

// fields

// 指数基本信息 fields
var FieldsIndexBasic = []string{"ts_code", "name", "fullname", "market", "publisher", "index_type", "category", "base_date", "base_point", "list_date", "weight_rule", "desc", "exp_date"}

// 指数日线行情 fields
var FieldsIndexDaily = []string{"ts_code", "trade_date", "close", "open", "high", "low", "pre_close", "change", "pct_chg", "vol", "amount"}

// 指数周线行情 fields
var FieldsIndexWeekly = []string{"ts_code", "trade_date", "close", "open", "high", "low", "pre_close", "change", "pct_chg", "vol", "amount"}

// 指数月线行情 fields
var FieldsIndexMonthly = []string{"ts_code", "trade_date", "close", "open", "high", "low", "pre_close", "change", "pct_chg", "vol", "amount"}

// 指数成分和权重 fields
var FieldsIndexWeight = []string{"index_code", "con_code", "trade_date", "weight"}

// 大盘指数每日指标 fields
var FieldsIndexDailyBasic = []string{"ts_code", "trade_date", "total_mv", "float_mv", "total_share", "float_share", "free_share", "turnover_rate", "turnover_rate_f", "pe", "pe_ttm", "pb"}

// 申万行业分类 fields
var FieldsIndexClassify = []string{"index_code", "industry_name", "parent_code", "level", "industry_code", "is_pub", "src"}

// 申万行业成分 fields
var FieldsIndexMember = []string{"index_code", "index_name", "con_code", "con_name", "in_date", "out_date", "is_new"}

// 沪深市场每日交易统计 fields
var FieldsDailyInfo = []string{"trade_date", "ts_code", "ts_name", "com_count", "total_share", "float_share", "total_mv", "float_mv", "amount", "vol", "trans_count", "pe", "tr", "exchange"}

// 深圳市场每日交易情况 fields
var FieldsSzDailyInfo = []string{"trade_date", "ts_code", "count", "amount", "vol", "total_share", "total_mv", "float_share", "float_mv"}

// 同花顺概念和行业指数 fields
var FieldsThsIndex = []string{"ts_code", "name", "count", "exchange", "list_date", "type"}

// 同花顺概念和行业指数行情 fields
var FieldsThsDaily = []string{"ts_code", "trade_date", "close", "open", "high", "low", "pre_close", "avg_price", "change", "pct_change", "vol", "turnover_rate", "total_mv", "float_mv"}

// 同花顺概念和行业指数成分 fields
var FieldsThsMember = []string{"ts_code", "code", "name", "weight", "in_date", "out_date", "is_new"}

// 中信行业指数行情 fields
var FieldsCiDaily = []string{"ts_code", "trade_date", "open", "low", "high", "close", "pre_close", "change", "pct_change", "vol", "amount"}

// 国际主要指数 fields
var FieldsIndexGlobal = []string{"ts_code", "trade_date", "open", "close", "high", "low", "pre_close", "change", "pct_chg", "swing", "vol", "amount"}

// struct

// 指数基本信息|index_basic
type IndexBasic struct {
	TsCode     string  `json:"ts_code"`     // TS代码
	Name       string  `json:"name"`        // 简称
	Fullname   string  `json:"fullname"`    // 指数全称
	Market     string  `json:"market"`      // 市场
	Publisher  string  `json:"publisher"`   // 发布方
	IndexType  string  `json:"index_type"`  // 指数风格
	Category   string  `json:"category"`    // 指数类别
	BaseDate   string  `json:"base_date"`   // 基期
	BasePoint  float64 `json:"base_point"`  // 基点
	ListDate   string  `json:"list_date"`   // 发布日期
	WeightRule string  `json:"weight_rule"` // 加权方式
	Desc       string  `json:"desc"`        // 描述
	ExpDate    string  `json:"exp_date"`    // 终止日期
}

type IndexBasicRequest struct {
	Limit     string `json:"limit"`
	Offset    string `json:"offset"`
	TsCode    string `json:"ts_code"`   // 指数代码
	Name      string `json:"name"`      // 指数简称
	Market    string `json:"market"`    // 交易所或服务商(默认SSE)
	Publisher string `json:"publisher"` // 发布商
	Category  string `json:"category"`  // 指数类别
}

type IndexBasicResponse struct {
	List []*IndexBasic `json:"list"`
}

func (x *IndexBasicResponse) String() string {
	bytes, _ := json.Marshal(x)
	return string(bytes)
}

// 指数日线行情|index_daily
type IndexDaily struct {
	TsCode    string  `json:"ts_code"`    // TS指数代码
	TradeDate string  `json:"trade_date"` // 交易日
	Close     float64 `json:"close"`      // 收盘点位
	Open      float64 `json:"open"`       // 开盘点位
	High      float64 `json:"high"`       // 最高点位
	Low       float64 `json:"low"`        // 最低点位
	PreClose  float64 `json:"pre_close"`  // 昨日收盘点
	Change    float64 `json:"change"`     // 涨跌点
	PctChg    float64 `json:"pct_chg"`    // 涨跌幅（%）
	Vol       float64 `json:"vol"`        // 成交量（手）
	Amount    float64 `json:"amount"`     // 成交额（千元）
}

type IndexDailyRequest struct {
	Limit     string `json:"limit"`
	Offset    string `json:"offset"`
	TsCode    string `json:"ts_code"`    // 指数代码
	TradeDate string `json:"trade_date"` // 交易日期 （日期格式：YYYYMMDD，下同）
	StartDate string `json:"start_date"` // 开始日期
	EndDate   string `json:"end_date"`   // 结束日期
}

type IndexDailyResponse struct {
	List []*IndexDaily `json:"list"`
}

func (x *IndexDailyResponse) String() string {
	bytes, _ := json.Marshal(x)
	return string(bytes)
}

// 指数周线行情|index_weekly
type IndexWeekly struct {
	TsCode    string  `json:"ts_code"`    // TS指数代码
	TradeDate string  `json:"trade_date"` // 交易日
	Close     float64 `json:"close"`      // 收盘点位
	Open      float64 `json:"open"`       // 开盘点位
	High      float64 `json:"high"`       // 最高点位
	Low       float64 `json:"low"`        // 最低点位
	PreClose  float64 `json:"pre_close"`  // 昨日收盘点
	Change    float64 `json:"change"`     // 涨跌点位
	PctChg    float64 `json:"pct_chg"`    // 涨跌幅
	Vol       float64 `json:"vol"`        // 成交量（手）
	Amount    float64 `json:"amount"`     // 成交额（千元）
}

type IndexWeeklyRequest struct {
	Limit     string `json:"limit"`
	Offset    string `json:"offset"`
	TsCode    string `json:"ts_code"`    // TS代码
	TradeDate string `json:"trade_date"` // 交易日期
	StartDate string `json:"start_date"` // 开始日期
	EndDate   string `json:"end_date"`   // 结束日期
}

type IndexWeeklyResponse struct {
	List []*IndexWeekly `json:"list"`
}

func (x *IndexWeeklyResponse) String() string {
	bytes, _ := json.Marshal(x)
	return string(bytes)
}

// 指数月线行情|index_monthly
type IndexMonthly struct {
	TsCode    string  `json:"ts_code"`    // TS指数代码
	TradeDate string  `json:"trade_date"` // 交易日
	Close     float64 `json:"close"`      // 收盘点位
	Open      float64 `json:"open"`       // 开盘点位
	High      float64 `json:"high"`       // 最高点位
	Low       float64 `json:"low"`        // 最低点位
	PreClose  float64 `json:"pre_close"`  // 昨日收盘点
	Change    float64 `json:"change"`     // 涨跌点位
	PctChg    float64 `json:"pct_chg"`    // 涨跌幅
	Vol       float64 `json:"vol"`        // 成交量（手）
	Amount    float64 `json:"amount"`     // 成交额（千元）
}

type IndexMonthlyRequest struct {
	Limit     string `json:"limit"`
	Offset    string `json:"offset"`
	TsCode    string `json:"ts_code"`    // TS代码
	TradeDate string `json:"trade_date"` // 交易日期
	StartDate string `json:"start_date"` // 开始日期
	EndDate   string `json:"end_date"`   // 结束日期
}

type IndexMonthlyResponse struct {
	List []*IndexMonthly `json:"list"`
}

func (x *IndexMonthlyResponse) String() string {
	bytes, _ := json.Marshal(x)
	return string(bytes)
}

// 指数成分和权重|index_weight
type IndexWeight struct {
	IndexCode string  `json:"index_code"` // 指数代码
	ConCode   string  `json:"con_code"`   // 成分代码
	TradeDate string  `json:"trade_date"` // 交易日期
	Weight    float64 `json:"weight"`     // 权重
}

type IndexWeightRequest struct {
	Limit     string `json:"limit"`
	Offset    string `json:"offset"`
	IndexCode string `json:"index_code"` // 指数代码 (二选一)
	TradeDate string `json:"trade_date"` // 交易日期 （二选一）
	StartDate string `json:"start_date"` // 开始日期
	EndDate   string `json:"end_date"`   // 结束日期
}

type IndexWeightResponse struct {
	List []*IndexWeight `json:"list"`
}

func (x *IndexWeightResponse) String() string {
	bytes, _ := json.Marshal(x)
	return string(bytes)
}

// 大盘指数每日指标|index_dailybasic
type IndexDailyBasic struct {
	TsCode        string  `json:"ts_code"`         // TS代码
	TradeDate     string  `json:"trade_date"`      // 交易日期
	TotalMv       float64 `json:"total_mv"`        // 当日总市值（元）
	FloatMv       float64 `json:"float_mv"`        // 当日流通市值（元）
	TotalShare    float64 `json:"total_share"`     // 当日总股本（股）
	FloatShare    float64 `json:"float_share"`     // 当日流通股本（股）
	FreeShare     float64 `json:"free_share"`      // 当日自由流通股本（股）
	TurnoverRate  float64 `json:"turnover_rate"`   // 换手率
	TurnoverRateF float64 `json:"turnover_rate_f"` // 换手率(基于自由流通股本)
	Pe            float64 `json:"pe"`              // 市盈率
	PeTtm         float64 `json:"pe_ttm"`          // 市盈率TTM
	Pb            float64 `json:"pb"`              // 市净率
}

type IndexDailyBasicRequest struct {
	Limit     string `json:"limit"`
	Offset    string `json:"offset"`
	TradeDate string `json:"trade_date"` // 交易日期 （格式：YYYYMMDD，比如20181018，下同）
	TsCode    string `json:"ts_code"`    // TS代码
	StartDate string `json:"start_date"` // 开始日期
	EndDate   string `json:"end_date"`   // 结束日期
}

type IndexDailyBasicResponse struct {
	List []*IndexDailyBasic `json:"list"`
}

func (x *IndexDailyBasicResponse) String() string {
	bytes, _ := json.Marshal(x)
	return string(bytes)
}

// 申万行业分类|index_classify
type IndexClassify struct {
	IndexCode    string `json:"index_code"`    // 指数代码
	IndustryName string `json:"industry_name"` // 行业名称
	ParentCode   string `json:"parent_code"`   // 父级代码
	Level        string `json:"level"`         // 行业名称
	IndustryCode string `json:"industry_code"` // 行业代码
	IsPub        string `json:"is_pub"`        // 是否发布了指数
	Src          string `json:"src"`           // 行业分类（SW申万）
}

type IndexClassifyRequest struct {
	Limit      string `json:"limit"`
	Offset     string `json:"offset"`
	IndexCode  string `json:"index_code"`  // 指数代码
	Level      string `json:"level"`       // 行业分级（L1/L2/L3）
	ParentCode string `json:"parent_code"` // 父级代码（一级为0）
	Src        string `json:"src"`         // 指数来源（SW2014：申万2014年版本，SW2021：申万2021年版本）
}

type IndexClassifyResponse struct {
	List []*IndexClassify `json:"list"`
}

func (x *IndexClassifyResponse) String() string {
	bytes, _ := json.Marshal(x)
	return string(bytes)
}

// 申万行业成分|index_member
type IndexMember struct {
	IndexCode string `json:"index_code"` // 指数代码
	IndexName string `json:"index_name"` // 指数名称
	ConCode   string `json:"con_code"`   // 成分股票代码
	ConName   string `json:"con_name"`   // 成分股票名称
	InDate    string `json:"in_date"`    // 纳入日期
	OutDate   string `json:"out_date"`   // 剔除日期
	IsNew     string `json:"is_new"`     // 是否最新Y是N否
}

type IndexMemberRequest struct {
	Limit     string `json:"limit"`
	Offset    string `json:"offset"`
	IndexCode string `json:"index_code"` // 指数代码
	TsCode    string `json:"ts_code"`    // 股票代码
	IsNew     string `json:"is_new"`     // 是否最新（默认为“Y是”）
}

type IndexMemberResponse struct {
	List []*IndexMember `json:"list"`
}

func (x *IndexMemberResponse) String() string {
	bytes, _ := json.Marshal(x)
	return string(bytes)
}

// 沪深市场每日交易统计|daily_info
type DailyInfo struct {
	TradeDate  string  `json:"trade_date"`  // 交易日期
	TsCode     string  `json:"ts_code"`     // 市场代码
	TsName     string  `json:"ts_name"`     // 市场名称
	ComCount   int64   `json:"com_count"`   // 挂牌数
	TotalShare float64 `json:"total_share"` // 总股本（亿股）
	FloatShare float64 `json:"float_share"` // 流通股本（亿股）
	TotalMv    float64 `json:"total_mv"`    // 总市值（亿元）
	FloatMv    float64 `json:"float_mv"`    // 流通市值（亿元）
	Amount     float64 `json:"amount"`      // 交易金额（亿元）
	Vol        float64 `json:"vol"`         // 成交量（亿股）
	TransCount int64   `json:"trans_count"` // 成交笔数（万笔）
	Pe         float64 `json:"pe"`          // 平均市盈率
	Tr         float64 `json:"tr"`          // 换手率（％），注：深交所暂无此列
	Exchange   string  `json:"exchange"`    // 交易所（SH上交所 SZ深交所）
}

type DailyInfoRequest struct {
	Limit     string `json:"limit"`
	Offset    string `json:"offset"`
	TradeDate string `json:"trade_date"` // 交易日期（YYYYMMDD格式，下同）
	TsCode    string `json:"ts_code"`    // 板块代码（请参阅下方列表）
	Exchange  string `json:"exchange"`   // 股票市场（SH上交所 SZ深交所）
	StartDate string `json:"start_date"` // 开始日期
	EndDate   string `json:"end_date"`   // 结束日期
	Fields    string `json:"fields"`     // 指定提取字段
}

type DailyInfoResponse struct {
	List []*DailyInfo `json:"list"`
}

func (x *DailyInfoResponse) String() string {
	bytes, _ := json.Marshal(x)
	return string(bytes)
}

// 深圳市场每日交易情况|sz_daily_info
type SzDailyInfo struct {
	TradeDate  string  `json:"trade_date"`  // 交易日期
	TsCode     string  `json:"ts_code"`     // 市场类型
	Count      int64   `json:"count"`       // 股票个数
	Amount     float64 `json:"amount"`      // 成交金额
	Vol        string  `json:"vol"`         // 成交量
	TotalShare float64 `json:"total_share"` // 总股本
	TotalMv    float64 `json:"total_mv"`    // 总市值
	FloatShare float64 `json:"float_share"` // 流通股票
	FloatMv    float64 `json:"float_mv"`    // 流通市值
}

type SzDailyInfoRequest struct {
	Limit     string `json:"limit"`
	Offset    string `json:"offset"`
	TradeDate string `json:"trade_date"` // 交易日期（YYYYMMDD格式，下同）
	TsCode    string `json:"ts_code"`    // 板块代码
	StartDate string `json:"start_date"` // 开始日期
	EndDate   string `json:"end_date"`   // 结束日期
}

type SzDailyInfoResponse struct {
	List []*SzDailyInfo `json:"list"`
}

func (x *SzDailyInfoResponse) String() string {
	bytes, _ := json.Marshal(x)
	return string(bytes)
}

// 同花顺概念和行业指数|ths_index
type ThsIndex struct {
	TsCode   string `json:"ts_code"`   // 代码
	Name     string `json:"name"`      // 名称
	Count    int64  `json:"count"`     // 成分个数
	Exchange string `json:"exchange"`  // 交易所
	ListDate string `json:"list_date"` // 上市日期
	Type     string `json:"type"`      // N概念指数S特色指数
}

type ThsIndexRequest struct {
	Limit    string `json:"limit"`
	Offset   string `json:"offset"`
	TsCode   string `json:"ts_code"`  // 指数代码
	Exchange string `json:"exchange"` // 市场类型A-a股 HK-港股 US-美股
	Type     string `json:"type"`     // 指数类型 N-板块指数 I-行业指数 R-地域指数 S-同花顺特色指数
}

type ThsIndexResponse struct {
	List []*ThsIndex `json:"list"`
}

func (x *ThsIndexResponse) String() string {
	bytes, _ := json.Marshal(x)
	return string(bytes)
}

// 同花顺概念和行业指数行情|ths_daily
type ThsDaily struct {
	TsCode       string  `json:"ts_code"`       // TS指数代码
	TradeDate    string  `json:"trade_date"`    // 交易日
	Close        float64 `json:"close"`         // 收盘点位
	Open         float64 `json:"open"`          // 开盘点位
	High         float64 `json:"high"`          // 最高点位
	Low          float64 `json:"low"`           // 最低点位
	PreClose     float64 `json:"pre_close"`     // 昨日收盘点
	AvgPrice     float64 `json:"avg_price"`     // 平均价
	Change       float64 `json:"change"`        // 涨跌点位
	PctChange    float64 `json:"pct_change"`    // 涨跌幅
	Vol          float64 `json:"vol"`           // 成交量
	TurnoverRate float64 `json:"turnover_rate"` // 换手率
	TotalMv      float64 `json:"total_mv"`      // 总市值
	FloatMv      float64 `json:"float_mv"`      // 流通市值
}

type ThsDailyRequest struct {
	Limit     string `json:"limit"`
	Offset    string `json:"offset"`
	TsCode    string `json:"ts_code"`    // 指数代码
	TradeDate string `json:"trade_date"` // 交易日期（YYYYMMDD格式，下同）
	StartDate string `json:"start_date"` // 开始日期
	EndDate   string `json:"end_date"`   // 结束日期
}

type ThsDailyResponse struct {
	List []*ThsDaily `json:"list"`
}

func (x *ThsDailyResponse) String() string {
	bytes, _ := json.Marshal(x)
	return string(bytes)
}

// 同花顺概念和行业指数成分|ths_member
type ThsMember struct {
	TsCode  string  `json:"ts_code"`  // 指数代码
	Code    string  `json:"code"`     // 股票代码
	Name    string  `json:"name"`     // 股票名称
	Weight  float64 `json:"weight"`   // 权重(暂无)
	InDate  string  `json:"in_date"`  // 纳入日期(暂无)
	OutDate string  `json:"out_date"` // 剔除日期(暂无)
	IsNew   string  `json:"is_new"`   // 是否最新Y是N否
}

type ThsMemberRequest struct {
	Limit  string `json:"limit"`
	Offset string `json:"offset"`
	TsCode string `json:"ts_code"` // 板块指数代码
	Code   string `json:"code"`    // 股票代码
}

type ThsMemberResponse struct {
	List []*ThsMember `json:"list"`
}

func (x *ThsMemberResponse) String() string {
	bytes, _ := json.Marshal(x)
	return string(bytes)
}

// 中信行业指数行情|ci_daily
type CiDaily struct {
	TsCode    string  `json:"ts_code"`    // 指数代码
	TradeDate string  `json:"trade_date"` // 交易日期
	Open      float64 `json:"open"`       // 开盘点位
	Low       float64 `json:"low"`        // 最低点位
	High      float64 `json:"high"`       // 最高点位
	Close     float64 `json:"close"`      // 收盘点位
	PreClose  float64 `json:"pre_close"`  // 昨日收盘点位
	Change    float64 `json:"change"`     // 涨跌点位
	PctChange float64 `json:"pct_change"` // 涨跌幅
	Vol       float64 `json:"vol"`        // 成交量（万股）
	Amount    float64 `json:"amount"`     // 成交额（万元）
}

type CiDailyRequest struct {
	Limit     string `json:"limit"`
	Offset    string `json:"offset"`
	TsCode    string `json:"ts_code"`    // 行业代码
	TradeDate string `json:"trade_date"` // 交易日期（YYYYMMDD格式，下同）
	StartDate string `json:"start_date"` // 开始日期
	EndDate   string `json:"end_date"`   // 结束日期
}

type CiDailyResponse struct {
	List []*CiDaily `json:"list"`
}

func (x *CiDailyResponse) String() string {
	bytes, _ := json.Marshal(x)
	return string(bytes)
}

// 国际主要指数|index_global
type IndexGlobal struct {
	TsCode    string  `json:"ts_code"`    // TS指数代码
	TradeDate string  `json:"trade_date"` // 交易日
	Open      float64 `json:"open"`       // 开盘点位
	Close     float64 `json:"close"`      // 收盘点位
	High      float64 `json:"high"`       // 最高点位
	Low       float64 `json:"low"`        // 最低点位
	PreClose  float64 `json:"pre_close"`  // 昨日收盘点
	Change    float64 `json:"change"`     // 涨跌点位
	PctChg    float64 `json:"pct_chg"`    // 涨跌幅
	Swing     float64 `json:"swing"`      // 振幅
	Vol       float64 `json:"vol"`        // 成交量 （大部分无此项数据）
	Amount    float64 `json:"amount"`     // 成交额 （大部分无此项数据）
}

type IndexGlobalRequest struct {
	Limit     string `json:"limit"`
	Offset    string `json:"offset"`
	TsCode    string `json:"ts_code"`    // TS指数代码，见下表
	TradeDate string `json:"trade_date"` // 交易日期，YYYYMMDD格式，下同
	StartDate string `json:"start_date"` // 开始日期
	EndDate   string `json:"end_date"`   // 结束日期
}

type IndexGlobalResponse struct {
	List []*IndexGlobal `json:"list"`
}

func (x *IndexGlobalResponse) String() string {
	bytes, _ := json.Marshal(x)
	return string(bytes)
}
