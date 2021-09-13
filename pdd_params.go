package go_vesdk

type PDDCatsReq struct {
	ParentCatId int64 `json:"parent_cat_id"`
}

type PDDCatsRsp struct {
	Error string `json:"error"`
	Msg   string `json:"msg"`
	Data  []struct {
		CatName     string `json:"cat_name"`
		Level       string `json:"level"`
		CatId       string `json:"cat_id"`
		ParentCatId string `json:"parent_cat_id"`
	} `json:"data"`
}

type PDDGoodsOptReq struct {
	ParentOptId int `json:"parent_opt_id"`
}

type PDDGoodsOptRsp struct {
	Error string `json:"error"`
	Msg   string `json:"msg"`
	Data  []struct {
		ParentOptId string `json:"parent_opt_id"`
		Level       string `json:"level"`
		OptName     string `json:"opt_name"`
		OptId       string `json:"opt_id"`
	} `json:"data"`
}

type PDDGoodsSearchRangeList struct {
	RangeId   int64 `json:"range_id"`
	RangeFrom int64 `json:"range_from"`
	RangeTo   int64 `json:"range_to"`
}
type PDDGoodsSearchReq struct {
	AccessToken      string                    `json:"access_token"` // 必填
	ActivityTags     []int                     `json:"activity_tags"`
	BlockCatPackages []int                     `json:"block_cat_packages"`
	BlockCats        []int                     `json:"block_cats"`
	CatId            int                       `json:"cat_id"`
	CustomParameters string                    `json:"custom_parameters"`
	GoodsImgType     int                       `json:"goods_img_type"`
	GoodsSignList    []string                  `json:"goods_sign_list"`
	IsBrandGoods     bool                      `json:"is_brand_goods"`
	Keyword          string                    `json:"keyword"` // keyword和opt_id商品关键词，与opt_id字段选填一个或全部填写。可支持goods_id、拼多多链接（即拼多多app商详的链接）、进宝长链/短链（即为pdd.ddk.goods.promotion.url.generate接口生成的长短链）
	ListId           string                    `json:"list_id"`
	MerchantType     int                       `json:"merchant_type"`
	MerchantTypeList []int                     `json:"merchant_type_list"`
	OptId            int64                     `json:"opt_id"` // 商品标签类目ID，使用pdd.goods.opt.get获取
	Page             int                       `json:"page"`
	PageSize         int                       `json:"page_size"`
	Pid              string                    `json:"pid"` // 必填
	RangeList        []PDDGoodsSearchRangeList `json:"range_list"`
	SortType         int                       `json:"sort_type"`
	UseCustomized    bool                      `json:"use_customized"`
	WithCoupon       bool                      `json:"with_coupon"`
}

type PDDGoodsSearchRsp struct {
	Error     string `json:"error"`
	Msg       string `json:"msg"`
	Total     string `json:"total"`
	ListId    string `json:"list_id"`
	SearchId  string `json:"search_id"`
	RequestId string `json:"request_id"`
	Data      []struct {
		CategoryName                string   `json:"category_name"`
		CouponRemainQuantity        string   `json:"coupon_remain_quantity"`
		PromotionRate               string   `json:"promotion_rate"`
		ServiceTags                 []int    `json:"service_tags"`
		MallId                      string   `json:"mall_id"`
		MallName                    string   `json:"mall_name"`
		MallCouponEndTime           string   `json:"mall_coupon_end_time"`
		LgstTxt                     string   `json:"lgst_txt"`
		GoodsName                   string   `json:"goods_name"`
		HasMaterial                 bool     `json:"has_material"`
		GoodsId                     string   `json:"goods_id"`
		BrandName                   string   `json:"brand_name,omitempty"`
		PredictPromotionRate        string   `json:"predict_promotion_rate"`
		GoodsDesc                   string   `json:"goods_desc"`
		OptName                     string   `json:"opt_name"`
		ShareRate                   string   `json:"share_rate"`
		OptIds                      []int    `json:"opt_ids"`
		GoodsImageUrl               string   `json:"goods_image_url"`
		HasMallCoupon               bool     `json:"has_mall_coupon"`
		UnifiedTags                 []string `json:"unified_tags"`
		CouponStartTime             string   `json:"coupon_start_time"`
		MinGroupPrice               string   `json:"min_group_price"`
		CouponDiscount              string   `json:"coupon_discount"`
		CouponEndTime               string   `json:"coupon_end_time"`
		ZsDuoId                     string   `json:"zs_duo_id"`
		MallCouponRemainQuantity    string   `json:"mall_coupon_remain_quantity"`
		PlanType                    string   `json:"plan_type"`
		ExtraCouponAmount           string   `json:"extra_coupon_amount,omitempty"`
		CatIds                      []int    `json:"cat_ids"`
		CouponMinOrderAmount        string   `json:"coupon_min_order_amount"`
		CategoryId                  string   `json:"category_id"`
		MallCouponDiscountPct       string   `json:"mall_coupon_discount_pct"`
		SubsidyAmount               string   `json:"subsidy_amount,omitempty"`
		CouponTotalQuantity         string   `json:"coupon_total_quantity"`
		MallCouponMinOrderAmount    string   `json:"mall_coupon_min_order_amount"`
		MerchantType                string   `json:"merchant_type"`
		SalesTip                    string   `json:"sales_tip"`
		OnlySceneAuth               bool     `json:"only_scene_auth"`
		DescTxt                     string   `json:"desc_txt"`
		MallCouponId                string   `json:"mall_coupon_id"`
		GoodsThumbnailUrl           string   `json:"goods_thumbnail_url"`
		OptId                       string   `json:"opt_id"`
		SearchId                    string   `json:"search_id"`
		ActivityTags                []int    `json:"activity_tags"`
		HasCoupon                   bool     `json:"has_coupon"`
		MinNormalPrice              string   `json:"min_normal_price"`
		MallCouponStartTime         string   `json:"mall_coupon_start_time"`
		ServTxt                     string   `json:"serv_txt"`
		MallCouponTotalQuantity     string   `json:"mall_coupon_total_quantity"`
		MallCouponMaxDiscountAmount string   `json:"mall_coupon_max_discount_amount"`
		MallCps                     string   `json:"mall_cps"`
		GoodsSign                   string   `json:"goods_sign"`
		SubsidyDuoAmountTenMillion  string   `json:"subsidy_duo_amount_ten_million,omitempty"`
		CltCpnEndTime               string   `json:"clt_cpn_end_time,omitempty"`
		CltCpnMinAmt                string   `json:"clt_cpn_min_amt,omitempty"`
		CltCpnRemainQuantity        string   `json:"clt_cpn_remain_quantity,omitempty"`
		CltCpnBatchSn               string   `json:"clt_cpn_batch_sn,omitempty"`
		CltCpnDiscount              string   `json:"clt_cpn_discount,omitempty"`
		CltCpnQuantity              string   `json:"clt_cpn_quantity,omitempty"`
		CltCpnStartTime             string   `json:"clt_cpn_start_time,omitempty"`
		ActivityPromotionRate       string   `json:"activity_promotion_rate,omitempty"`
		ActivityType                string   `json:"activity_type,omitempty"`
	} `json:"data"`
}
