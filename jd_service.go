package go_vesdk

import (
	"github.com/azerothyang/go-vesdk/httpclient"
	"github.com/go-dragon/util"
	"reflect"
	"strconv"
	"strings"
)

// IJDService 京东服务接口
type IJDService interface {
	// JDSearch 关键词商品查询接口
	// https://union.jd.com/openplatform/api/10420
	JDSearch(jdParams *JDSearchParams) (*JDSearchRsp, error)

	// Category 新版分类接口category
	// https://union.jd.com/openplatform/api/10434
	Category(parentId int64, grade int64) (*JDCategoryRsp, error)

	// PromotionGoodsInfo 获取推广商品信息接口
	// https://union.jd.com/openplatform/api/10422
	// skuIds 必选，京东skuID串，逗号分割，最多100个（非常重要 请大家关注：如果输入的sk串中某个skuID的商品不在推广中[就是没有佣金]，返回结果中不会包含这个商品的信息
	PromotionGoodsInfo(skuIds []string) (*JDPromotionGoodsInfoRsp, error)

	// PromByUid 通过unionId获取推广链接
	// https://union.jd.com/openplatform/api/10425
	PromByUid(jdPromotionCodeReq *JDPromotionCodeReq) (*JDPromByUidRsp, error)

	// OrderRowQuery 订单行查询接口
	// https://union.jd.com/openplatform/api/12707
	OrderRowQuery(jdOrderReq *JDOrderReq) (*JDOrderRsp, error)

	// BigField 大字段商品详情查询接口
	// https://union.jd.com/openplatform/api/11248
	BigField(skuId []string, fields []string) (*JDBigFieldRsp, error)

	// JDMaterialQuery 物料商品查询
	// https://union.jd.com/openplatform/api/13625
	JDMaterialQuery(jdMaterialQueryReq *JDMaterialQueryReq) (*JDMaterialQueryRsp, error)
}

// NewJDService 初始化京东服务
func NewJDService(config Config) IJDService {
	return &JDService{config: config}
}

// JDService 具体结构体
type JDService struct {
	config Config
}

func (j *JDService) JDMaterialQuery(jdMaterialQueryReq *JDMaterialQueryReq) (*JDMaterialQueryRsp, error) {
	queryRsp := new(JDMaterialQueryRsp)
	refTy := reflect.TypeOf(*jdMaterialQueryReq)
	refV := reflect.ValueOf(*jdMaterialQueryReq)
	httpParams := map[string]string{
		"vekey": j.config.VeKey,
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
	rsp := httpclient.DefaultClient.GET("http://apivip.vephp.com/jd/jd_materialquery", httpParams, nil)
	if rsp.Err != nil {
		return queryRsp, rsp.Err
	}
	err := util.FastJson.Unmarshal([]byte(rsp.Content), queryRsp)
	if err != nil {
		veRspError := new(VeRspError)
		util.FastJson.Unmarshal([]byte(rsp.Content), veRspError)
		return queryRsp, veRspError
	}
	return queryRsp, err
}

func (j *JDService) BigField(skuIds []string, fields []string) (*JDBigFieldRsp, error) {
	jdBigFieldRsp := new(JDBigFieldRsp)
	httpParams := map[string]string{
		"vekey":  j.config.VeKey,
		"skuIds": strings.Join(skuIds, ","),
	}
	if len(fields) > 1 {
		httpParams["fields"] = strings.Join(fields, ",")
	}
	rsp := httpclient.DefaultClient.GET("http://apivip.vephp.com/jd/bigfield", httpParams, nil)
	if rsp.Err != nil {
		return nil, rsp.Err
	}
	err := util.FastJson.Unmarshal([]byte(rsp.Content), jdBigFieldRsp)
	if err != nil {
		veRspError := new(VeRspError)
		util.FastJson.Unmarshal([]byte(rsp.Content), veRspError)
		return jdBigFieldRsp, veRspError
	}
	return jdBigFieldRsp, nil
}

func (j *JDService) OrderRowQuery(orderReq *JDOrderReq) (*JDOrderRsp, error) {
	jdOrderRsp := new(JDOrderRsp)
	refTy := reflect.TypeOf(*orderReq)
	refV := reflect.ValueOf(*orderReq)
	httpParams := map[string]string{
		"vekey": j.config.VeKey,
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
		case int64:
			v := val.(int64)
			if v != 0 {
				httpParams[key] = strconv.FormatInt(v, 10)
			}
		}
	}
	rsp := httpclient.DefaultClient.GET("http://apivip.vephp.com/jd/orderrowquery", httpParams, nil)
	if rsp.Err != nil {
		return jdOrderRsp, rsp.Err
	}
	err := util.FastJson.Unmarshal([]byte(rsp.Content), jdOrderRsp)
	if err != nil {
		veRspError := new(VeRspError)
		util.FastJson.Unmarshal([]byte(rsp.Content), veRspError)
		return jdOrderRsp, veRspError
	}
	return jdOrderRsp, err
}

func (j *JDService) PromByUid(jdPromotionCodeReq *JDPromotionCodeReq) (*JDPromByUidRsp, error) {
	jdPromByUidRsp := new(JDPromByUidRsp)
	refTy := reflect.TypeOf(*jdPromotionCodeReq)
	refV := reflect.ValueOf(*jdPromotionCodeReq)
	httpParams := map[string]string{
		"vekey": j.config.VeKey,
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
		case int64:
			v := val.(int64)
			if v != 0 {
				httpParams[key] = strconv.FormatInt(v, 10)
			}
		}
	}
	rsp := httpclient.DefaultClient.GET("http://apivip.vephp.com/jd/prombyuid", httpParams, nil)
	if rsp.Err != nil {
		return jdPromByUidRsp, rsp.Err
	}
	err := util.FastJson.Unmarshal([]byte(rsp.Content), jdPromByUidRsp)
	if err != nil {
		veRspError := new(VeRspError)
		util.FastJson.Unmarshal([]byte(rsp.Content), veRspError)
		return jdPromByUidRsp, veRspError
	}
	return jdPromByUidRsp, err
}

func (j *JDService) PromotionGoodsInfo(skuIds []string) (*JDPromotionGoodsInfoRsp, error) {
	jdPromotionGoodsInfoRsp := new(JDPromotionGoodsInfoRsp)
	skuIdsStr := strings.Join(skuIds, ",")
	httpParams := map[string]string{
		"vekey":  j.config.VeKey,
		"skuIds": skuIdsStr,
	}
	rsp := httpclient.DefaultClient.GET("http://apivip.vephp.com/jd/promotiongoodsinfo", httpParams, nil)
	if rsp.Err != nil {
		return jdPromotionGoodsInfoRsp, rsp.Err
	}

	err := util.FastJson.Unmarshal([]byte(rsp.Content), jdPromotionGoodsInfoRsp)
	if err != nil {
		veRspError := new(VeRspError)
		util.FastJson.Unmarshal([]byte(rsp.Content), veRspError)
		return jdPromotionGoodsInfoRsp, veRspError
	}
	return jdPromotionGoodsInfoRsp, err
}

func (j *JDService) JDSearch(jdParams *JDSearchParams) (*JDSearchRsp, error) {
	jdSearchRsp := new(JDSearchRsp)
	refTy := reflect.TypeOf(*jdParams)
	refV := reflect.ValueOf(*jdParams)
	httpParams := map[string]string{
		"vekey": j.config.VeKey,
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
	rsp := httpclient.DefaultClient.GET("http://apivip.vephp.com/jd/jd_search", httpParams, nil)
	if rsp.Err != nil {
		return jdSearchRsp, rsp.Err
	}

	err := util.FastJson.Unmarshal([]byte(rsp.Content), jdSearchRsp)
	if err != nil {
		veRspError := new(VeRspError)
		util.FastJson.Unmarshal([]byte(rsp.Content), veRspError)
		return jdSearchRsp, veRspError
	}
	return jdSearchRsp, err
}

func (j *JDService) Category(parentId int64, grade int64) (*JDCategoryRsp, error) {
	jdCatRsp := new(JDCategoryRsp)
	httpParams := map[string]string{
		"vekey":    j.config.VeKey,
		"parentId": strconv.FormatInt(parentId, 10),
		"grade":    strconv.FormatInt(grade, 10),
	}
	rsp := httpclient.DefaultClient.GET("http://apivip.vephp.com/jd/category", httpParams, nil)
	if rsp.Err != nil {
		return jdCatRsp, rsp.Err
	}
	err := util.FastJson.Unmarshal([]byte(rsp.Content), jdCatRsp)
	if err != nil {
		veRspError := new(VeRspError)
		util.FastJson.Unmarshal([]byte(rsp.Content), veRspError)
		return jdCatRsp, veRspError
	}
	return jdCatRsp, err
}
