package go_vesdk

import "github.com/go-dragon/util"

type VipCommonParams struct {
	Plat        int    `json:"plat"`
	DeviceType  string `json:"deviceType"`
	DeviceValue string `json:"deviceValue"`
	Ip          string `json:"ip"`
	Longitude   string `json:"longitude"`
	Latitude    string `json:"latitude"`
	Warehouse   string `json:"warehouse"`
}

func (v VipCommonParams) toJson() string {
	data, _ := util.FastJson.Marshal(v)
	return string(data)
}

type VipSearchReq struct {
	AccessToken                 string          `json:"access_token"`
	Keyword                     string          `json:"keyword"`
	FieldName                   string          `json:"fieldName"`
	Order                       string          `json:"order"`
	Page                        int             `json:"page"`
	PageSize                    int             `json:"pageSize"`
	RequestId                   string          `json:"requestId"`
	PriceStart                  string          `json:"priceStart"`
	PriceEnd                    string          `json:"priceEnd"`
	QueryReputation             bool            `json:"queryReputation"`
	QueryStoreServiceCapability bool            `json:"queryStoreServiceCapability"`
	QueryStock                  bool            `json:"queryStock"`
	QueryActivity               bool            `json:"queryActivity"`
	QueryPrepay                 bool            `json:"queryPrepay"`
	CommonParams                VipCommonParams `json:"commonParams"`
	VendorCode                  string          `json:"vendorCode"`
	ChanTag                     string          `json:"chanTag"`
	QueryExclusiveCoupon        bool            `json:"queryExclusiveCoupon"`
	QueryCpsInfo                int             `json:"queryCpsInfo"`
	Mobile                      string          `json:"mobile"`
	Virtual                     Virtual         `json:"virtual"`
}

type VipSearchRsp struct {
	Error     string `json:"error"`
	Msg       string `json:"msg"`
	IsVirtual int    `json:"is_virtual"`
	Data      struct {
		Total         int `json:"total"`
		GoodsInfoList []struct {
			MarketPrice           string   `json:"marketPrice"`
			CommissionRate        string   `json:"commissionRate"`
			GoodsId               string   `json:"goodsId"`
			Discount              string   `json:"discount"`
			GoodsCarouselPictures []string `json:"goodsCarouselPictures"`
			GoodsDetailPictures   []string `json:"goodsDetailPictures"`
			CategoryName          string   `json:"categoryName"`
			HaiTao                int      `json:"haiTao"`
			PrepayInfo            struct {
				IsPrepay int `json:"isPrepay"`
			} `json:"prepayInfo"`
			Cat2NdName      string `json:"cat2ndName"`
			Cat1StName      string `json:"cat1stName"`
			VipPrice        string `json:"vipPrice"`
			Commission      string `json:"commission"`
			Sn              string `json:"sn"`
			Cat1StId        int    `json:"cat1stId"`
			GoodsName       string `json:"goodsName"`
			BrandName       string `json:"brandName"`
			BrandLogoFull   string `json:"brandLogoFull"`
			BrandStoreSn    string `json:"brandStoreSn"`
			SellTimeFrom    int64  `json:"sellTimeFrom"`
			SchemeStartTime int64  `json:"schemeStartTime"`
			CommentsInfo    struct {
				Comments          int    `json:"comments"`
				GoodCommentsShare string `json:"goodCommentsShare,omitempty"`
			} `json:"commentsInfo"`
			SaleStockStatus int    `json:"saleStockStatus"`
			SchemeEndTime   int64  `json:"schemeEndTime"`
			SourceType      int    `json:"sourceType"`
			SellTimeTo      int64  `json:"sellTimeTo"`
			BrandId         int    `json:"brandId"`
			GoodsThumbUrl   string `json:"goodsThumbUrl"`
			Cat2NdId        int    `json:"cat2ndId"`
			SpuId           string `json:"spuId"`
			StoreInfo       struct {
				StoreName string `json:"storeName"`
				StoreId   string `json:"storeId"`
			} `json:"storeInfo"`
			GoodsMainPicture string `json:"goodsMainPicture"`
			DestUrl          string `json:"destUrl"`
			CategoryId       int    `json:"categoryId"`
			Status           int    `json:"status"`
			CouponInfo       struct {
				UseEndTime        int64  `json:"useEndTime"`
				TotalAmount       int    `json:"totalAmount"`
				CouponName        string `json:"couponName"`
				ActivateEndTime   int64  `json:"activateEndTime"`
				CouponType        int    `json:"couponType"`
				Buy               string `json:"buy"`
				UseBeginTime      int64  `json:"useBeginTime"`
				CouponNo          string `json:"couponNo"`
				Fav               string `json:"fav"`
				ActivateBeginTime int64  `json:"activateBeginTime"`
				ActivedAmount     int    `json:"activedAmount"`
			} `json:"couponInfo,omitempty"`
		} `json:"goodsInfoList"`
		PageSize   int `json:"pageSize"`
		SortFields []struct {
			FieldName string `json:"fieldName"`
			FieldDesc string `json:"fieldDesc"`
		} `json:"sortFields"`
		Page int `json:"page"`
	} `json:"data"`
}
