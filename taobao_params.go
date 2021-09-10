package go_vesdk

var (
	SearchType = map[string]string{
		"normalSearch":  "1", // 泛搜索
		"specialSearch": "2", // 定向搜索
	}


)

// OrderDetailsParams 订单查询 start
type OrderDetailsParams struct {
	StartTime     string `json:"start_time"`
	EndTime       string `json:"end_time"`
	QueryType     int    `json:"query_type"`
	PositionIndex string `json:"position_index"`
	PageNo        int    `json:"page_no"`
	PageSize      int    `json:"page_size"`
	MemberType    int    `json:"member_type"`
	TKStatus      int    `json:"tk_status"`
	JumpType      int    `json:"jump_type"`
	OrderScene    int    `json:"order_scene"`
}

type OrderDetailsRsp struct {
	Error string `json:"error"`
	Msg   string `json:"msg"`
	Data  struct {
		HasNext       bool   `json:"has_next"`
		HasPre        bool   `json:"has_pre"`
		PageNo        int    `json:"page_no"`
		PageSize      int    `json:"page_size"`
		PositionIndex string `json:"position_index"`
		Results       struct {
			PublisherOrderDto []struct {
				AdzoneId                           int64  `json:"adzone_id"`
				AdzoneName                         string `json:"adzone_name"`
				AlimamaRate                        string `json:"alimama_rate"`
				AlimamaShareFee                    string `json:"alimama_share_fee"`
				AlipayTotalPrice                   string `json:"alipay_total_price"`
				ClickTime                          string `json:"click_time"`
				DepositPrice                       string `json:"deposit_price"`
				FlowSource                         string `json:"flow_source"`
				IncomeRate                         string `json:"income_rate"`
				IsLx                               string `json:"is_lx"`
				ItemCategoryName                   string `json:"item_category_name"`
				ItemId                             int64  `json:"item_id"`
				ItemImg                            string `json:"item_img"`
				ItemLink                           string `json:"item_link"`
				ItemNum                            int    `json:"item_num"`
				ItemPrice                          string `json:"item_price"`
				ItemTitle                          string `json:"item_title"`
				MarketingType                      string `json:"marketing_type"`
				ModifiedTime                       string `json:"modified_time"`
				OrderType                          string `json:"order_type"`
				PayPrice                           string `json:"pay_price,omitempty"`
				PubId                              int    `json:"pub_id"`
				PubShareFee                        string `json:"pub_share_fee"`
				PubSharePreFee                     string `json:"pub_share_pre_fee"`
				PubShareRate                       string `json:"pub_share_rate"`
				RefundTag                          int    `json:"refund_tag"`
				SellerNick                         string `json:"seller_nick"`
				SellerShopTitle                    string `json:"seller_shop_title"`
				SiteId                             int    `json:"site_id"`
				SiteName                           string `json:"site_name"`
				SubsidyFee                         string `json:"subsidy_fee"`
				SubsidyRate                        string `json:"subsidy_rate"`
				SubsidyType                        string `json:"subsidy_type"`
				TbDepositTime                      string `json:"tb_deposit_time"`
				TbPaidTime                         string `json:"tb_paid_time"`
				TerminalType                       string `json:"terminal_type"`
				TkCommissionFeeForMediaPlatform    string `json:"tk_commission_fee_for_media_platform"`
				TkCommissionPreFeeForMediaPlatform string `json:"tk_commission_pre_fee_for_media_platform"`
				TkCommissionRateForMediaPlatform   string `json:"tk_commission_rate_for_media_platform"`
				TkCreateTime                       string `json:"tk_create_time"`
				TkDepositTime                      string `json:"tk_deposit_time"`
				TkEarningTime                      string `json:"tk_earning_time,omitempty"`
				TkOrderRole                        int    `json:"tk_order_role"`
				TkPaidTime                         string `json:"tk_paid_time"`
				TkStatus                           int    `json:"tk_status"`
				TkTotalRate                        string `json:"tk_total_rate"`
				TotalCommissionFee                 string `json:"total_commission_fee"`
				TotalCommissionRate                string `json:"total_commission_rate"`
				TradeId                            string `json:"trade_id"`
				TradeParentId                      string `json:"trade_parent_id"`
				AppKey                             string `json:"app_key,omitempty"`
			} `json:"publisher_order_dto"`
		} `json:"results"`
	} `json:"data"`
}

// 订单查询 end

type HCApiParams struct {
	Para        string `json:"para"`
	Pid         string `json:"pid"`
	RelationId  string `json:"relationId"`
	SpecialId   string `json:"special_id"`
	ExternalId  string `json:"external_id"`
	Detail      string `json:"detail"`
	ActivityId  string `json:"activity_id"`
	DeepCoupon  string `json:"deepcoupon"`
	ShopCoupon  string `json:"shopcoupon"`
	NoTkl       string `json:"notkl"`
	NoShortLink string `json:"noshortlink"`
	CouponId    string `json:"couponId"`
	TklType     string `json:"tkl_type"`
	Xid         string `json:"xid"`
	GetTopnRate string `json:"get_topn_rate"`
	UcrowdId    string `json:"ucrowd_id"`
}

type HCApiRsp struct {
	Error  string `json:"error"`
	Msg    string `json:"msg"`
	Result int    `json:"result"`
	Data   struct {
		CategoryId        string `json:"category_id"`
		CouponClickUrl    string `json:"coupon_click_url"`
		CouponEndTime     string `json:"coupon_end_time"`
		CouponInfo        string `json:"coupon_info"`
		CouponRemainCount string `json:"coupon_remain_count"`
		CouponStartTime   string `json:"coupon_start_time"`
		CouponTotalCount  string `json:"coupon_total_count"`
		MaxCommissionRate string `json:"max_commission_rate"`
		RewardInfo        string `json:"reward_info"`
		CommissionRate    string `json:"commission_rate"`
		NumIid            string `json:"num_iid"`
		SclickUrl         string `json:"sclick_url"`
		TbkPwd            string `json:"tbk_pwd"`
		IosTbkPwd         string `json:"ios_tbk_pwd"`
		GlobalTbkPwd      string `json:"global_tbk_pwd"`
		CouponShortUrl    string `json:"coupon_short_url"`
	} `json:"data"`
}

type HCApiAllRsp struct {
	Error   string `json:"error"`
	Msg     string `json:"msg"`
	Count   int    `json:"count"`
	Success int    `json:"success"`
	Fail    int    `json:"fail"`
	FailIds string `json:"fail_ids"`
	Data    []struct {
		CategoryId        string `json:"category_id"`
		CouponClickUrl    string `json:"coupon_click_url"`
		CouponEndTime     string `json:"coupon_end_time"`
		CouponInfo        string `json:"coupon_info"`
		CouponRemainCount string `json:"coupon_remain_count"`
		CouponStartTime   string `json:"coupon_start_time"`
		CouponTotalCount  string `json:"coupon_total_count"`
		MaxCommissionRate string `json:"max_commission_rate"`
		RewardInfo        string `json:"reward_info"`
		CommissionRate    string `json:"commission_rate"`
		NumIid            string `json:"num_iid"`
		SclickUrl         string `json:"sclick_url"`
		TbkPwd            string `json:"tbk_pwd"`
		IosTbkPwd         string `json:"ios_tbk_pwd"`
		GlobalTbkPwd      string `json:"global_tbk_pwd"`
		CouponShortUrl    string `json:"coupon_short_url"`
	} `json:"data"`
}

type SuperSearchReq struct {
	Para              string  `json:"para"`
	Cat               string  `json:"cat"`
	Pid               string  `json:"pid"`
	RelationId        string  `json:"relationId"`
	Page              int     `json:"page"`
	PageSize          int     `json:"pagesize"`
	ForceIndex        string  `json:"force_index"`
	Coupon            string  `json:"coupon"`
	StartPrice        string  `json:"start_price"`
	EndPrice          string  `json:"end_price"`
	StartTkRate       string  `json:"start_tk_rate"`
	EndTkRate         string  `json:"end_tk_rate"`
	IsOverseas        int     `json:"is_overseas"`
	IsTmall           int     `json:"is_tmall"`
	Sort              string  `json:"sort"`
	OnlySearch        int     `json:"onlysearch"`
	Similar           int     `json:"similar"`
	SearchItem        int     `json:"search_item"`
	Ip                string  `json:"ip"`
	FreeShip          int     `json:"freeship"`
	Npx               int     `json:"npx"`
	ItemId            int64   `json:"item_id"`
	Virtual           int     `json:"virtual"`
	Presale           int     `json:"presale"`
	IncludePayRate30  int     `json:"include_pay_rate_30"`
	IncludeGoodRate   int     `json:"include_good_rate"`
	EndKaTkRate       int     `json:"end_ka_tk_rate"`
	StartKaTkRate     int     `json:"start_ka_tk_rate"`
	DeviceValue       string  `json:"device_value"`
	DeviceEncrypt     string  `json:"device_encrypt"`
	DeviceType        string  `json:"device_type"`
	LockRateEndTime   int64   `json:"lock_rate_end_time"`
	LockRateStartTime int64   `json:"lock_rate_start_time"`
	SellerIds         string  `json:"seller_ids"`
	CityCode          string  `json:"city_code"`
	Latitude          float64 `json:"latitude"`
	Longitude         float64 `json:"longitude"`
	SpecialId         string  `json:"special_id"`
	NeedPrepay        int     `json:"need_prepay"`
	TkLink            int     `json:"tklink"`
}

type SuperSearchRsp struct {
	Error        string `json:"error"`
	Msg          string `json:"msg"`
	SearchType   string `json:"search_type"`
	IsVirtual    int    `json:"is_virtual"`
	IsSimilar    string `json:"is_similar"`
	IsSplitWord  int    `json:"is_splitWord"`
	ForceIndex   string `json:"force_index"`
	TotalResults int    `json:"total_results"`
	ResultList   []struct {
		CategoryId           int      `json:"category_id"`
		CategoryName         string   `json:"category_name"`
		CommissionRate       string   `json:"commission_rate"`
		CouponId             string   `json:"coupon_id"`
		CouponInfo           string   `json:"coupon_info"`
		CouponRemainCount    int      `json:"coupon_remain_count"`
		CouponTotalCount     int      `json:"coupon_total_count"`
		ItemDescription      string   `json:"item_description"`
		ItemId               int64    `json:"item_id"`
		ItemUrl              string   `json:"item_url"`
		LevelOneCategoryId   int      `json:"level_one_category_id"`
		LevelOneCategoryName string   `json:"level_one_category_name"`
		Nick                 string   `json:"nick"`
		NumIid               int64    `json:"num_iid"`
		PictUrl              string   `json:"pict_url"`
		PresaleDeposit       string   `json:"presale_deposit"`
		PresaleEndTime       int      `json:"presale_end_time"`
		PresaleStartTime     int      `json:"presale_start_time"`
		PresaleTailEndTime   int      `json:"presale_tail_end_time"`
		PresaleTailStartTime int      `json:"presale_tail_start_time"`
		Provcity             string   `json:"provcity"`
		RealPostFee          string   `json:"real_post_fee"`
		ReservePrice         string   `json:"reserve_price"`
		SellerId             int64    `json:"seller_id"`
		ShopDsr              int      `json:"shop_dsr"`
		ShopTitle            string   `json:"shop_title"`
		ShortTitle           string   `json:"short_title"`
		SmallImages          []string `json:"small_images"`
		SuperiorBrand        string   `json:"superior_brand"`
		Title                string   `json:"title"`
		TkTotalCommi         string   `json:"tk_total_commi"`
		TkTotalSales         string   `json:"tk_total_sales"`
		Url                  string   `json:"url"`
		UserType             int      `json:"user_type"`
		Volume               int      `json:"volume"`
		WhiteImage           string   `json:"white_image"`
		XId                  string   `json:"x_id"`
		ZkFinalPrice         string   `json:"zk_final_price"`
		KuadianPromotionInfo string   `json:"kuadian_promotion_info,omitempty"`
	} `json:"result_list"`
	Result      int    `json:"result"`
	Data        struct {
		CategoryId           string   `json:"category_id"`
		CouponClickUrl       string   `json:"coupon_click_url"`
		CouponEndTime        string   `json:"coupon_end_time"`
		CouponInfo           string   `json:"coupon_info"`
		CouponRemainCount    string   `json:"coupon_remain_count"`
		CouponStartTime      string   `json:"coupon_start_time"`
		CouponTotalCount     string   `json:"coupon_total_count"`
		MaxCommissionRate    string   `json:"max_commission_rate"`
		RewardInfo           string   `json:"reward_info"`
		CommissionRate       string   `json:"commission_rate"`
		NumIid               string   `json:"num_iid"`
		SclickUrl            string   `json:"sclick_url"`
		ZkFinalPrice         string   `json:"zk_final_price"`
		WhiteImage           string   `json:"white_image"`
		Volume               string   `json:"volume"`
		UserType             string   `json:"user_type"`
		Title                string   `json:"title"`
		SuperiorBrand        string   `json:"superior_brand"`
		SmallImages          []string `json:"small_images"`
		ShopTitle            string   `json:"shop_title"`
		ShopDsr              string   `json:"shop_dsr"`
		SellerId             string   `json:"seller_id"`
		ReservePrice         string   `json:"reserve_price"`
		RealPostFee          string   `json:"real_post_fee"`
		PresaleTailStartTime int      `json:"presale_tail_start_time"`
		PresaleTailEndTime   int      `json:"presale_tail_end_time"`
		PresaleStartTime     int      `json:"presale_start_time"`
		PresaleEndTime       int      `json:"presale_end_time"`
		PresaleDeposit       string   `json:"presale_deposit"`
		PictUrl              string   `json:"pict_url"`
		Nick                 string   `json:"nick"`
		KuadianPromotionInfo string   `json:"kuadian_promotion_info"`
		ItemUrl              string   `json:"item_url"`
		CatName              string   `json:"cat_name"`
		CatLeafName          string   `json:"cat_leaf_name"`
		TbkPwd               string   `json:"tbk_pwd"`
		IosTbkPwd            string   `json:"ios_tbk_pwd"`
		GlobalTbkPwd         string   `json:"global_tbk_pwd"`
		CouponShortUrl       string   `json:"coupon_short_url"`
	} `json:"data"`
}