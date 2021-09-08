package go_vesdk

import (
	"github.com/azerothyang/go-vesdk/httpclient"
	"github.com/go-dragon/util"
	"reflect"
	"strconv"
)

// ITaoBaoService 淘客服务接口
type ITaoBaoService interface {
	// OrderDetails 订单查询
	OrderDetails(orderDetailParams *OrderDetailsParams) (*OrderDetailsRsp, error)
	// HCApiOne 高佣转链
	HCApiOne(hcParams *HCApiParams) (*HCApiRsp, error)
	// HCApiAllByItemIds 多个商品id批量高佣转链
	HCApiAllByItemIds(hcParams *HCApiParams) (*HCApiAllRsp, error)
}

// TaoBaoService 具体结构体
type TaoBaoService struct {
	config Config
}

// NewTaoBaoService 初始化淘宝服务
func NewTaoBaoService(config Config) ITaoBaoService {
	return &TaoBaoService{config: config}
}

// OrderDetails 拉取订单列表
// http://wsd.591hufu.com/doc/xinbantaokedingdanapi
func (t *TaoBaoService) OrderDetails(orderDetailParams *OrderDetailsParams) (*OrderDetailsRsp, error) {
	var orderDetailsRsp = &OrderDetailsRsp{}
	refTy := reflect.TypeOf(*orderDetailParams)
	refV := reflect.ValueOf(*orderDetailParams)
	httpParams := map[string]string{
		"vekey": t.config.VeKey,
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
		case int:
			v := val.(int)
			if v != 0 {
				httpParams[key] = strconv.FormatInt(int64(v), 10)
			}
		}
	}
	rsp := httpclient.DefaultClient.GET(veUrl+"/orderdetails", httpParams, nil)
	if rsp.Err != nil {
		return orderDetailsRsp, rsp.Err
	}
	err := util.FastJson.Unmarshal([]byte(rsp.Content), orderDetailsRsp)
	if err != nil {
		veRspError := new(VeRspError)
		util.FastJson.Unmarshal([]byte(rsp.Content), veRspError)
		return orderDetailsRsp, veRspError
	}
	return orderDetailsRsp, err
}

// HCApiOne 高佣转链单个
// http://wsd.591hufu.com/doc/gaoyongzhuanlianjiekou
func (t *TaoBaoService) HCApiOne(hcParams *HCApiParams) (*HCApiRsp, error) {
	var hcApiRsp = &HCApiRsp{}
	refTy := reflect.TypeOf(*hcParams)
	refV := reflect.ValueOf(*hcParams)
	httpParams := map[string]string{
		"vekey": t.config.VeKey,
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
		case int:
			v := val.(int)
			if v != 0 {
				httpParams[key] = strconv.FormatInt(int64(v), 10)
			}
		}
	}
	rsp := httpclient.DefaultClient.GET(veUrl+"/hcapi", httpParams, nil)
	if rsp.Err != nil {
		return hcApiRsp, rsp.Err
	}
	err := util.FastJson.Unmarshal([]byte(rsp.Content), hcApiRsp)
	if err != nil {
		veRspError := new(VeRspError)
		util.FastJson.Unmarshal([]byte(rsp.Content), veRspError)
		return hcApiRsp, veRspError
	}
	return hcApiRsp, err
}

// HCApiAllByItemIds 高佣转链
// para传: 23883488910,598450446711,984504  多个商品id
// http://wsd.591hufu.com/doc/gaoyongzhuanlianjiekou
func (t *TaoBaoService) HCApiAllByItemIds(hcParams *HCApiParams) (*HCApiAllRsp, error) {
	var hcApiAllRsp = &HCApiAllRsp{}
	refTy := reflect.TypeOf(*hcParams)
	refV := reflect.ValueOf(*hcParams)
	httpParams := map[string]string{
		"vekey": t.config.VeKey,
		"plhc":  "1",
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
		case int:
			v := val.(int)
			if v != 0 {
				httpParams[key] = strconv.FormatInt(int64(v), 10)
			}
		}
	}
	rsp := httpclient.DefaultClient.GET(veUrl+"/hcapi", httpParams, nil)
	if rsp.Err != nil {
		return hcApiAllRsp, rsp.Err
	}
	err := util.FastJson.Unmarshal([]byte(rsp.Content), hcApiAllRsp)
	if err != nil {
		veRspError := new(VeRspError)
		util.FastJson.Unmarshal([]byte(rsp.Content), veRspError)
		return hcApiAllRsp, veRspError
	}
	return hcApiAllRsp, err
}
