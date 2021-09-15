package go_vesdk

import (
	"github.com/azerothyang/go-vesdk/httpclient"
	"github.com/go-dragon/util"
	"net/url"
	"reflect"
	"strconv"
)

// ITaoBaoService 淘客服务接口
type ITaoBaoService interface {
	// OrderDetails 订单查询
	// http://wsd.591hufu.com/doc/xinbantaokedingdanapi
	// 具体参数可以参考官方文档：https://open.taobao.com/api.htm?docId=43755&docType=2&scopeId=16322
	OrderDetails(*OrderDetailsParams) (*OrderDetailsRsp, error)
	// HCApiOne 高佣转链
	// http://wsd.591hufu.com/doc/gaoyongzhuanlianjiekou
	HCApiOne(*HCApiParams) (*HCApiRsp, error)
	// HCApiAllByItemIds 多个商品id批量高佣转链
	// http://wsd.591hufu.com/doc/gaoyongzhuanlianjiekou
	HCApiAllByItemIds(*HCApiParams) (*HCApiAllRsp, error)
	// SuperSearch 超级搜索接口-淘宝联盟强大搜索功能帮你把握成交概率
	// http://wsd.591hufu.com/doc/chaojisousuojiekou
	// SuperSearchListRsp结果代表泛搜索，搜索返回列表
	// SuperSearchHCApiRsp结果代表定向搜索，搜索返回高佣转链
	SuperSearch(*SuperSearchReq) (*SuperSearchRsp, error)

	// OrderPunish 淘宝客处罚订单查询用法
	// http://wsd.591hufu.com/doc/chufadingdanchaxun
	// 官方文档：https://open.taobao.com/api.htm?docId=41942&docType=2&scopeId=15738
	OrderPunish(*OrderPunishReq) (*OrderPunishRsp, error)

	// RefundOrder 淘客维权订单接口-支持渠道订单和会员订单维权
	// http://wsd.591hufu.com/doc/taokeweiquandingdanjiekou
	// 官方文档：https://open.taobao.com/api.htm?docId=43755&docType=2&scopeId=16322
	RefundOrder(*RefundOrderReq) (*RefundOrderRsp, error)

	// TbSimpleProductDetail 淘宝或天猫的全网产品详情接口，简版产品详情
	// http://wsd.591hufu.com/doc/quanwangshangpinxiangqing
	TbSimpleProductDetail(*TbSimpleDetailReq) (*TbSimpleDetailRsp, error)

	// TbProductDetail 淘宝或天猫的全网产品详情接口
	// http://wsd.591hufu.com/doc/quanwangshangpinxiangqing
	TbProductDetail(*TbProductDetailReq) (*TbProductDetailRsp, error)
}

// TaoBaoService 具体结构体
type TaoBaoService struct {
	config Config
}

// NewTaoBaoService 初始化淘宝服务
func NewTaoBaoService(config Config) ITaoBaoService {
	return &TaoBaoService{config: config}
}

func (t *TaoBaoService) TbSimpleProductDetail(tbSimpleDetailReq *TbSimpleDetailReq) (*TbSimpleDetailRsp, error) {
	tbSimpleDetailRsp := new(TbSimpleDetailRsp)
	refTy := reflect.TypeOf(*tbSimpleDetailReq)
	refV := reflect.ValueOf(*tbSimpleDetailReq)
	httpParams := map[string]string{
		"vekey":      t.config.VeKey,
		"onlysimple": "1",
	}
	for i := 0; i < refTy.NumField(); i++ {
		val := refV.Field(i).Interface()
		key := refTy.Field(i).Tag.Get("json")
		switch val.(type) {
		case int:
			v := val.(int)
			if v != 0 {
				httpParams[key] = strconv.FormatInt(int64(v), 10)
			}
		case int64:
			v := val.(int64)
			if v != 0 {
				httpParams[key] = strconv.FormatInt(v, 10)
			}
		}
	}
	rsp := httpclient.DefaultClient.GET(veUrl+"/tbprodetail", httpParams, nil)
	if rsp.Err != nil {
		return tbSimpleDetailRsp, rsp.Err
	}
	err := util.FastJson.Unmarshal([]byte(rsp.Content), tbSimpleDetailRsp)
	if err != nil {
		veRspError := new(VeRspError)
		util.FastJson.Unmarshal([]byte(rsp.Content), veRspError)
		return tbSimpleDetailRsp, veRspError
	}
	return tbSimpleDetailRsp, err
}

func (t *TaoBaoService) TbProductDetail(tbProductDetailReq *TbProductDetailReq) (*TbProductDetailRsp, error) {
	tbProductDetailRsp := new(TbProductDetailRsp)
	refTy := reflect.TypeOf(*tbProductDetailReq)
	refV := reflect.ValueOf(*tbProductDetailReq)
	httpParams := map[string]string{
		"vekey": t.config.VeKey,
	}
	for i := 0; i < refTy.NumField(); i++ {
		val := refV.Field(i).Interface()
		key := refTy.Field(i).Tag.Get("json")
		switch val.(type) {
		case int:
			v := val.(int)
			if v != 0 {
				httpParams[key] = strconv.FormatInt(int64(v), 10)
			}
		case int64:
			v := val.(int64)
			if v != 0 {
				httpParams[key] = strconv.FormatInt(v, 10)
			}
		}
	}
	rsp := httpclient.DefaultClient.GET(veUrl+"/tbprodetail", httpParams, nil)
	if rsp.Err != nil {
		return tbProductDetailRsp, rsp.Err
	}
	err := util.FastJson.Unmarshal([]byte(rsp.Content), tbProductDetailRsp)
	if err != nil {
		veRspError := new(VeRspError)
		util.FastJson.Unmarshal([]byte(rsp.Content), veRspError)
		return tbProductDetailRsp, veRspError
	}
	return tbProductDetailRsp, err
}

func (t *TaoBaoService) RefundOrder(refundOrderReq *RefundOrderReq) (*RefundOrderRsp, error) {
	refundOrderRsp := new(RefundOrderRsp)
	refTy := reflect.TypeOf(*refundOrderReq)
	refV := reflect.ValueOf(*refundOrderReq)
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
	rsp := httpclient.DefaultClient.GET(veUrl+"/refundorder", httpParams, nil)
	if rsp.Err != nil {
		return refundOrderRsp, rsp.Err
	}
	err := util.FastJson.Unmarshal([]byte(rsp.Content), refundOrderRsp)
	if err != nil {
		veRspError := new(VeRspError)
		util.FastJson.Unmarshal([]byte(rsp.Content), veRspError)
		return refundOrderRsp, veRspError
	}
	return refundOrderRsp, err
}

func (t *TaoBaoService) OrderPunish(orderPunishReq *OrderPunishReq) (*OrderPunishRsp, error) {
	var orderPunRsp = new(OrderPunishRsp)
	refTy := reflect.TypeOf(*orderPunishReq)
	refV := reflect.ValueOf(*orderPunishReq)
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
	rsp := httpclient.DefaultClient.GET(veUrl+"/orderpunish", httpParams, nil)
	if rsp.Err != nil {
		return orderPunRsp, rsp.Err
	}
	err := util.FastJson.Unmarshal([]byte(rsp.Content), orderPunRsp)
	if err != nil {
		veRspError := new(VeRspError)
		util.FastJson.Unmarshal([]byte(rsp.Content), veRspError)
		return orderPunRsp, veRspError
	}
	return orderPunRsp, nil
}

func (t *TaoBaoService) SuperSearch(superSearchReq *SuperSearchReq) (*SuperSearchRsp, error) {
	var superRsp = new(SuperSearchRsp)
	refTy := reflect.TypeOf(*superSearchReq)
	refV := reflect.ValueOf(*superSearchReq)
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
				if key == "para" {
					// 需要url编码
					httpParams[key] = url.QueryEscape(v)
				} else {
					httpParams[key] = v
				}
			}
		case int:
			v := val.(int)
			if v != 0 {
				httpParams[key] = strconv.FormatInt(int64(v), 10)
			}
		case int64:
			v := val.(int64)
			if v != 0 {
				httpParams[key] = strconv.FormatInt(v, 10)
			}
		case float64:
			v := val.(float64)
			if v != 0 {
				httpParams[key] = strconv.FormatFloat(v, 'f', -1, 64)
			}
		}
	}
	rsp := httpclient.DefaultClient.GET(veUrl+"/super", httpParams, nil)
	if rsp.Err != nil {
		return superRsp, rsp.Err
	}
	err := util.FastJson.Unmarshal([]byte(rsp.Content), superRsp)
	if err != nil {
		veRspError := new(VeRspError)
		util.FastJson.Unmarshal([]byte(rsp.Content), veRspError)
		return superRsp, veRspError
	}
	return superRsp, nil
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
		case QueryType:
			v := val.(QueryType)
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
				if key == "para" {
					// 需要url编码
					httpParams[key] = url.QueryEscape(v)
				} else {
					httpParams[key] = v
				}
			}
		case int:
			v := val.(int)
			if v != 0 {
				httpParams[key] = strconv.FormatInt(int64(v), 10)
			}
		case int64:
			v := val.(int64)
			if v != 0 {
				httpParams[key] = strconv.FormatInt(v, 10)
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
				if key == "para" {
					// 需要url编码
					httpParams[key] = url.QueryEscape(v)
				} else {
					httpParams[key] = v
				}
			}
		case int:
			v := val.(int)
			if v != 0 {
				httpParams[key] = strconv.FormatInt(int64(v), 10)
			}
		case int64:
			v := val.(int64)
			if v != 0 {
				httpParams[key] = strconv.FormatInt(v, 10)
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
