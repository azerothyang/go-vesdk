package go_vesdk

import (
	"github.com/azerothyang/go-vesdk/httpclient"
	"github.com/go-dragon/util"
	"reflect"
	"strconv"
)

// IVipService 根据关键词查询商品列表
type IVipService interface {
	// Search 根据关键词查询商品列表：
	// 具体参数可以参考官方文档: http://vop.vip.com/apicenter/struct?serviceName=com.vip.adp.api.open.service.UnionGoodsService-1.0.0&struct=com.vip.adp.api.open.service.QueryGoodsRequest
	Search(*VipSearchReq) (*VipSearchRsp, error)
}

type VipService struct {
	config Config
}

func (v *VipService) Search(vipSearchReq *VipSearchReq) (*VipSearchRsp, error) {
	var vipSearchRsp = new(VipSearchRsp)
	refTy := reflect.TypeOf(*vipSearchReq)
	refV := reflect.ValueOf(*vipSearchReq)
	httpParams := map[string]string{
		"vekey": v.config.VeKey,
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
		case bool:
			v := val.(bool)
			if v {
				httpParams[key] = "true"
			} else {
				httpParams[key] = "false"
			}
		case VipCommonParams:
			v := val.(VipCommonParams)
			if !reflect.DeepEqual(v, VipCommonParams{}) {
				httpParams[key] = v.toJson()
			}
		case Virtual:
			v := val.(Virtual)
			if v != NoCheckVirtual {
				httpParams[key] = strconv.FormatInt(v.ToInt64(), 10)
			}
		}
	}
	rsp := httpclient.DefaultClient.GET(veUrl+"/vip/vip_search", httpParams, nil)
	if rsp.Err != nil {
		return vipSearchRsp, rsp.Err
	}
	err := util.FastJson.Unmarshal([]byte(rsp.Content), vipSearchRsp)
	if err != nil {
		veRspError := new(VeRspError)
		util.FastJson.Unmarshal([]byte(rsp.Content), veRspError)
		return vipSearchRsp, veRspError
	}
	return vipSearchRsp, err
}
