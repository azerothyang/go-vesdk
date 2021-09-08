package go_vesdk

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
