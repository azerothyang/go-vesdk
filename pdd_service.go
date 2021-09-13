package go_vesdk

import (
	"fmt"
	"github.com/azerothyang/go-vesdk/httpclient"
	"github.com/go-dragon/util"
	"reflect"
	"strconv"
	"strings"
)

// IPddService 拼多多服务接口
type IPddService interface {
	// Cats 获取标准分类：
	// 具体参数可以参考官方文档: http://open.pinduoduo.com/application/document/api?id=pdd.goods.cats.get
	Cats(*PddCatsReq) (*PddCatsRsp, error)
	// GoodsOpt 获取标准分类(和商品搜索关联)：
	// 具体参数可以参考官方文档: https://open.pinduoduo.com/application/document/api?id=pdd.goods.opt.get
	GoodsOpt(*PddGoodsOptReq) (*PddGoodsOptRsp, error)
	// GoodsSearch 多多进宝商品查询
	// 具体参数可以参考官方文档: https://open.pinduoduo.com/application/document/api?id=pdd.ddk.goods.search
	GoodsSearch(*PddGoodsSearchReq) (*PddGoodsSearchRsp, error)
	// PddAuthQuery 查询是否绑定备案
	// https://open.pinduoduo.com/application/document/api?id=pdd.ddk.oauth.member.authority.query
	PddAuthQuery(*PddAuthQueryReq) (*PddAuthQueryRsp, error)
	// PddGenerate PID和备案关联-生成授权备案链接
	// https://open.pinduoduo.com/application/document/api?id=pdd.ddk.goods.promotion.url.generate
	PddGenerate(*PddGenerateReq) (*PddGenerateRsp, error)
}

// PddService 具体结构体
type PddService struct {
	config Config
}

// NewPddService 初始化拼多多服务
func NewPddService(config Config) IPddService {
	return &PddService{config: config}
}

func (p *PddService) PddAuthQuery(pddAuthQueryReq *PddAuthQueryReq) (*PddAuthQueryRsp, error) {
	var pddAuthQueryRsp = new(PddAuthQueryRsp)
	refTy := reflect.TypeOf(*pddAuthQueryReq)
	refV := reflect.ValueOf(*pddAuthQueryReq)
	httpParams := map[string]string{
		"vekey": p.config.VeKey,
	}
	for i := 0; i < refTy.NumField(); i++ {
		val := refV.Field(i).Interface()
		key := refTy.Field(i).Tag.Get("json")
		switch val.(type) {
		case string:
			v := val.(string)
			if v != "" {
				httpParams[key] = v
			}
		}
	}
	rsp := httpclient.DefaultClient.GET(veUrl+"/pdd/pdd_authquery", httpParams, nil)
	if rsp.Err != nil {
		return pddAuthQueryRsp, rsp.Err
	}
	err := util.FastJson.Unmarshal([]byte(rsp.Content), pddAuthQueryRsp)
	if err != nil {
		veRspError := new(VeRspError)
		util.FastJson.Unmarshal([]byte(rsp.Content), veRspError)
		return pddAuthQueryRsp, veRspError
	}
	return pddAuthQueryRsp, err
}

func (p *PddService) PddGenerate(pddGenerateReq *PddGenerateReq) (*PddGenerateRsp, error) {
	var pddGenerateRsp = new(PddGenerateRsp)
	refTy := reflect.TypeOf(*pddGenerateReq)
	refV := reflect.ValueOf(*pddGenerateReq)
	httpParams := map[string]string{
		"vekey": p.config.VeKey,
	}
	for i := 0; i < refTy.NumField(); i++ {
		val := refV.Field(i).Interface()
		key := refTy.Field(i).Tag.Get("json")
		switch val.(type) {
		case string:
			v := val.(string)
			if v != "" {
				httpParams[key] = v
			}
		}
	}
	rsp := httpclient.DefaultClient.GET(veUrl+"/pdd/pdd_generate", httpParams, nil)
	if rsp.Err != nil {
		return pddGenerateRsp, rsp.Err
	}
	err := util.FastJson.Unmarshal([]byte(rsp.Content), pddGenerateRsp)
	if err != nil {
		veRspError := new(VeRspError)
		util.FastJson.Unmarshal([]byte(rsp.Content), veRspError)
		return pddGenerateRsp, veRspError
	}
	return pddGenerateRsp, err
}

func (p *PddService) Cats(pddCatsReq *PddCatsReq) (*PddCatsRsp, error) {
	var pddCatsRsp = new(PddCatsRsp)
	refTy := reflect.TypeOf(*pddCatsReq)
	refV := reflect.ValueOf(*pddCatsReq)
	httpParams := map[string]string{
		"vekey": p.config.VeKey,
	}
	for i := 0; i < refTy.NumField(); i++ {
		val := refV.Field(i).Interface()
		key := refTy.Field(i).Tag.Get("json")
		switch val.(type) {
		case int64:
			v := val.(int64)
			if v != 0 {
				httpParams[key] = strconv.FormatInt(v, 10)
			}
		}
	}
	rsp := httpclient.DefaultClient.GET(veUrl+"/pdd/pdd_cats", httpParams, nil)
	if rsp.Err != nil {
		return pddCatsRsp, rsp.Err
	}
	err := util.FastJson.Unmarshal([]byte(rsp.Content), pddCatsRsp)
	if err != nil {
		veRspError := new(VeRspError)
		util.FastJson.Unmarshal([]byte(rsp.Content), veRspError)
		return pddCatsRsp, veRspError
	}
	return pddCatsRsp, err
}

func (p *PddService) GoodsOpt(goodsOptReq *PddGoodsOptReq) (*PddGoodsOptRsp, error) {
	var goodsOptRsp = new(PddGoodsOptRsp)
	refTy := reflect.TypeOf(*goodsOptReq)
	refV := reflect.ValueOf(*goodsOptReq)
	httpParams := map[string]string{
		"vekey": p.config.VeKey,
	}
	for i := 0; i < refTy.NumField(); i++ {
		val := refV.Field(i).Interface()
		key := refTy.Field(i).Tag.Get("json")
		switch val.(type) {
		case int:
			v := val.(int)
			httpParams[key] = strconv.FormatInt(int64(v), 10)
		}
	}
	rsp := httpclient.DefaultClient.GET(veUrl+"/pdd/pdd_goodsopt", httpParams, nil)
	if rsp.Err != nil {
		return goodsOptRsp, rsp.Err
	}
	err := util.FastJson.Unmarshal([]byte(rsp.Content), goodsOptRsp)
	if err != nil {
		veRspError := new(VeRspError)
		util.FastJson.Unmarshal([]byte(rsp.Content), veRspError)
		return goodsOptRsp, veRspError
	}
	return goodsOptRsp, err
}

func (p *PddService) GoodsSearch(pddSearchReq *PddGoodsSearchReq) (*PddGoodsSearchRsp, error) {
	var pddSearchRsp = new(PddGoodsSearchRsp)
	refTy := reflect.TypeOf(*pddSearchReq)
	refV := reflect.ValueOf(*pddSearchReq)
	httpParams := map[string]string{
		"vekey": p.config.VeKey,
	}
	for i := 0; i < refTy.NumField(); i++ {
		val := refV.Field(i).Interface()
		key := refTy.Field(i).Tag.Get("json")
		switch val.(type) {
		case string:
			v := val.(string)
			if v != "" {
				httpParams[key] = v
			}
		case []int:
			v := val.([]int)
			if len(v) > 0 {
				params := fmt.Sprintf("[%s]", strings.Join(intSliceToStringSlice(v), ","))
				httpParams[key] = params
			}
		case []string:
			v := val.([]string)
			if len(v) > 0 {
				params := fmt.Sprintf("[%s]", strings.Join(v, ","))
				httpParams[key] = params
			}
		case int64:
			v := val.(int64)
			if v != 0 {
				httpParams[key] = strconv.FormatInt(v, 10)
			}
		case int:
			v := val.(int)
			if v != 0 {
				httpParams[key] = strconv.FormatInt(int64(v), 10)
			}
		case bool:
			v := val.(bool)
			if v {
				httpParams[key] = "true"
			}
		case []PddGoodsSearchRangeList:
			v := val.([]PddGoodsSearchRangeList)
			if len(v) > 0 {
				params, err := util.FastJson.Marshal(v)
				if err != nil {
					return pddSearchRsp, err
				}
				httpParams[key] = string(params)
			}
		}
	}
	rsp := httpclient.DefaultClient.GET(veUrl+"/pdd/pdd_search", httpParams, nil)
	if rsp.Err != nil {
		return pddSearchRsp, rsp.Err
	}
	err := util.FastJson.Unmarshal([]byte(rsp.Content), pddSearchRsp)
	if err != nil {
		veRspError := new(VeRspError)
		util.FastJson.Unmarshal([]byte(rsp.Content), veRspError)
		return pddSearchRsp, veRspError
	}
	return pddSearchRsp, err
}
