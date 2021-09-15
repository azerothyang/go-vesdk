package go_vesdk

type QueryType int

const (
	QueryTypeCreateTime QueryType = iota + 1
	QueryTypePaidTime
	QueryTypeTimeSettle
	QueryTypeUpdateTime
)

// OrderDetailsParams 订单查询 start
type OrderDetailsParams struct {
	StartTime     string    `json:"start_time"`
	EndTime       string    `json:"end_time"`
	QueryType     QueryType `json:"query_type"`
	PositionIndex string    `json:"position_index"`
	PageNo        int       `json:"page_no"`
	PageSize      int       `json:"page_size"`
	MemberType    int       `json:"member_type"`
	TKStatus      int       `json:"tk_status"`
	JumpType      int       `json:"jump_type"`
	OrderScene    int       `json:"order_scene"`
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
	RelationId  int64  `json:"relationId"`
	SpecialId   int64  `json:"special_id"`
	ExternalId  string `json:"external_id"`
	Detail      int    `json:"detail"`
	ActivityId  string `json:"activity_id"`
	DeepCoupon  int    `json:"deepcoupon"`
	ShopCoupon  int    `json:"shopcoupon"`
	NoTkl       int    `json:"notkl"`
	NoShortLink int    `json:"noshortlink"`
	CouponId    string `json:"couponId"`
	TklType     int    `json:"tkl_type"`
	Xid         string `json:"xid"`
	GetTopnRate string `json:"get_topn_rate"`
	UcrowdId    string `json:"ucrowd_id"`
}

type HCApiRsp struct {
	Error  string `json:"error"`
	Msg    string `json:"msg"`
	Result int    `json:"result"`
	Data   struct {
		CategoryId           string   `json:"category_id"`
		CouponClickUrl       string   `json:"coupon_click_url"`
		MaxCommissionRate    string   `json:"max_commission_rate"`
		RewardInfo           string   `json:"reward_info"`
		CommissionRate       string   `json:"commission_rate"`
		NumIid               string   `json:"num_iid"`
		OriginalUlandLink    string   `json:"original_uland_link"`
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

type HCApiAllRsp struct {
	Error   string `json:"error"`
	Msg     string `json:"msg"`
	Count   int    `json:"count"`
	Success int    `json:"success"`
	Fail    int    `json:"fail"`
	FailIds string `json:"fail_ids"`
	Data    []struct {
		CategoryId             string   `json:"category_id"`
		CouponClickUrl         string   `json:"coupon_click_url"`
		MaxCommissionRate      string   `json:"max_commission_rate"`
		RewardInfo             string   `json:"reward_info"`
		CommissionRate         string   `json:"commission_rate"`
		NumIid                 string   `json:"num_iid"`
		OriginalUlandLink      string   `json:"original_uland_link,omitempty"`
		SclickUrl              string   `json:"sclick_url"`
		ZkFinalPrice           string   `json:"zk_final_price"`
		WhiteImage             string   `json:"white_image"`
		Volume                 string   `json:"volume"`
		UserType               string   `json:"user_type"`
		Title                  string   `json:"title"`
		SuperiorBrand          string   `json:"superior_brand"`
		SmallImages            []string `json:"small_images"`
		ShopTitle              string   `json:"shop_title"`
		ShopDsr                string   `json:"shop_dsr"`
		SellerId               string   `json:"seller_id"`
		ReservePrice           string   `json:"reserve_price"`
		RealPostFee            string   `json:"real_post_fee"`
		PresaleTailStartTime   int64    `json:"presale_tail_start_time"`
		PresaleTailEndTime     int64    `json:"presale_tail_end_time"`
		PresaleStartTime       int64    `json:"presale_start_time"`
		PresaleEndTime         int64    `json:"presale_end_time"`
		PresaleDiscountFeeText string   `json:"presale_discount_fee_text,omitempty"`
		PresaleDeposit         string   `json:"presale_deposit"`
		PictUrl                string   `json:"pict_url"`
		Nick                   string   `json:"nick"`
		KuadianPromotionInfo   string   `json:"kuadian_promotion_info"`
		ItemUrl                string   `json:"item_url"`
		CatName                string   `json:"cat_name"`
		CatLeafName            string   `json:"cat_leaf_name"`
		TbkPwd                 string   `json:"tbk_pwd"`
		IosTbkPwd              string   `json:"ios_tbk_pwd"`
		GlobalTbkPwd           string   `json:"global_tbk_pwd"`
		CouponShortUrl         string   `json:"coupon_short_url"`
		CouponEndTime          string   `json:"coupon_end_time,omitempty"`
		CouponInfo             string   `json:"coupon_info,omitempty"`
		CouponRemainCount      string   `json:"coupon_remain_count,omitempty"`
		CouponStartTime        string   `json:"coupon_start_time,omitempty"`
		CouponTotalCount       string   `json:"coupon_total_count,omitempty"`
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
	Result int `json:"result"`
	Data   struct {
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

type OrderPunishReq struct {
	StartTime  string `json:"start_time"`
	TbTradeId  string `json:"tb_trade_id"`
	RelationId string `json:"relation_id"`
	Span       int    `json:"span"`
	PageNo     int    `json:"page_no"`
	PageSize   int    `json:"page_size"`
	AdzoneId   string `json:"adzone_id"`
	SiteId     string `json:"site_id"`
}

// OrderPunishRsp
// PunishStatus处罚状态。0 冻结，1 解冻
type OrderPunishRsp struct {
	Error      string `json:"error"`
	Msg        string `json:"msg"`
	PageNo     int    `json:"page_no"`
	PageSize   int    `json:"page_size"`
	TotalCount int    `json:"total_count"`
	Data       struct {
		Result []struct {
			PunishStatus      string `json:"punish_status"`
			RelationId        int64  `json:"relation_id"`
			SettleMonth       int    `json:"settle_month"`
			TbTradeId         int64  `json:"tb_trade_id"`
			TkAdzoneId        int64  `json:"tk_adzone_id"`
			TkPubId           int    `json:"tk_pub_id"`
			TkSiteId          int    `json:"tk_site_id"`
			TkTradeCreateTime string `json:"tk_trade_create_time"`
			ViolationType     string `json:"violation_type"`
		} `json:"result"`
	} `json:"data"`
}

// RefundOrderReq
//page_no 可选，默认值1，表示页码。翻页时使用。
//page_size 可选：默认值30，即每页返回30笔维权订单。
//search_type 可选：默认值1，可选值含义：1-维权发起时间，2-订单结算时间（正向订单），3-维权完成时间，4-订单创建时间，5-订单更新时间，即增量查询订单（出参更新字段modified_time）。
//refund_type 可选：默认值0，可选值含义：1 表示2方，2表示3方，0表示不限
//start_time 必选：表示开始时间，格式如2018-10-10 00:00:00，默认是一天时间。
//biz_type 可选：默认值3，1代表渠道关系id，2代表会员关系id，3表示渠道+会员
type RefundOrderReq struct {
	StartTime  string `json:"start_time"`
	PageNo     int    `json:"page_no"`
	PageSize   int    `json:"page_size"`
	SearchType int    `json:"search_type"`
	RefundType int    `json:"refund_type"`
	BizType    int    `json:"biz_type"`
}

type RefundOrderRsp struct {
	Error string `json:"error"`
	Msg   string `json:"msg"`
	Data  struct {
		BizErrorCode int `json:"biz_error_code"`
		Data         struct {
			PageNo   int `json:"page_no"`
			PageSize int `json:"page_size"`
			Results  struct {
				Result []struct {
					EarningTime              string `json:"earning_time"`
					ModifiedTime             string `json:"modified_time"`
					RefundFee                string `json:"refund_fee"`
					RefundStatus             int    `json:"refund_status"`
					RefundType               int    `json:"refund_type"`
					SpecialId                int64  `json:"special_id"`
					TbAuctionTitle           string `json:"tb_auction_title"`
					TbTradeCreateTime        string `json:"tb_trade_create_time"`
					TbTradeFinishPrice       string `json:"tb_trade_finish_price"`
					TbTradeId                int64  `json:"tb_trade_id"`
					TbTradeParentId          int64  `json:"tb_trade_parent_id"`
					TkCommissionFeeRefundPub string `json:"tk_commission_fee_refund_pub"`
					TkPubId                  int    `json:"tk_pub_id"`
					TkPubShowReturnFee       string `json:"tk_pub_show_return_fee"`
					TkRefundSuitTime         string `json:"tk_refund_suit_time"`
					TkRefundTime             string `json:"tk_refund_time"`
					TkSubsidyFeeRefundPub    string `json:"tk_subsidy_fee_refund_pub"`
				} `json:"result"`
			} `json:"results"`
			TotalCount int `json:"total_count"`
		} `json:"data"`
		ResultCode int `json:"result_code"`
	} `json:"data"`
}

type TbSimpleDetailReq struct {
	Id           int64 `json:"id"`
	OnlyShopInfo int   `json:"onlyshopinfo"`
	GetCatName   int   `json:"getcatname"`
}

type TbSimpleDetailRsp struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
	Data struct {
		Status   int      `json:"status"`
		Url      string   `json:"url"`
		Id       string   `json:"id"`
		Title    string   `json:"title"`
		Images   []string `json:"images"`
		Price    int      `json:"price"`
		VideoUrl string   `json:"videoUrl"`
		SellerID string   `json:"sellerID"`
		Platform string   `json:"platform"`
	} `json:"data"`
}

type TbProductDetailReq struct {
	Id           int64 `json:"id"`
	OnlyShopInfo int   `json:"onlyshopinfo"`
	OnlySimple   int   `json:"onlysimple"`
	GetCatName   int   `json:"getcatname"`
}

type TbProductDetailRsp struct {
	Error         string `json:"error"`
	Msg           string `json:"msg"`
	DebugCode     string `json:"debugCode"`
	RequestGlobal string `json:"requestGlobal"`
	Data          struct {
		ApiStack []struct {
			Name  string `json:"name"`
			Value string `json:"value"`
		} `json:"apiStack"`
		Seller struct {
			UserId           string `json:"userId"`
			ShopId           string `json:"shopId"`
			ShopName         string `json:"shopName"`
			ShopUrl          string `json:"shopUrl"`
			TaoShopUrl       string `json:"taoShopUrl"`
			ShopIcon         string `json:"shopIcon"`
			Fans             string `json:"fans"`
			AllItemCount     string `json:"allItemCount"`
			NewItemCount     string `json:"newItemCount"`
			ShowShopLinkIcon bool   `json:"showShopLinkIcon"`
			ShopCard         string `json:"shopCard"`
			SellerType       string `json:"sellerType"`
			ShopType         string `json:"shopType"`
			Evaluates        []struct {
				Title                     string `json:"title"`
				Score                     string `json:"score"`
				Type                      string `json:"type"`
				Level                     string `json:"level"`
				LevelText                 string `json:"levelText"`
				LevelTextColor            string `json:"levelTextColor"`
				LevelBackgroundColor      string `json:"levelBackgroundColor"`
				TmallLevelTextColor       string `json:"tmallLevelTextColor"`
				TmallLevelBackgroundColor string `json:"tmallLevelBackgroundColor"`
			} `json:"evaluates"`
			Evaluates2 []struct {
				TitleColor     string `json:"titleColor"`
				ScoreTextColor string `json:"scoreTextColor"`
				Title          string `json:"title"`
				Score          string `json:"score"`
				Type           string `json:"type"`
				Level          string `json:"level"`
				LevelText      string `json:"levelText"`
				LevelTextColor string `json:"levelTextColor"`
			} `json:"evaluates2"`
			SellerNick         string `json:"sellerNick"`
			CreditLevel        string `json:"creditLevel"`
			CreditLevelIcon    string `json:"creditLevelIcon"`
			BrandIcon          string `json:"brandIcon"`
			BrandIconRatio     string `json:"brandIconRatio"`
			Starts             string `json:"starts"`
			GoodRatePercentage string `json:"goodRatePercentage"`
			Fbt2User           string `json:"fbt2User"`
			SimpleShopDOStatus string `json:"simpleShopDOStatus"`
			ShopVersion        string `json:"shopVersion"`
			AtmosphereImg      string `json:"atmosphereImg"`
			AtmosphereColor    string `json:"atmosphereColor"`
			ShopTextColor      string `json:"shopTextColor"`
			EntranceList       []struct {
				Text            string `json:"text"`
				TextColor       string `json:"textColor"`
				BorderColor     string `json:"borderColor"`
				BackgroundColor string `json:"backgroundColor"`
				Action          []struct {
					Key    string `json:"key"`
					Params struct {
						Url         string `json:"url,omitempty"`
						TrackParams struct {
							Spm string `json:"spm"`
						} `json:"trackParams,omitempty"`
						TrackName string `json:"trackName,omitempty"`
					} `json:"params"`
				} `json:"action"`
			} `json:"entranceList"`
			AtmophereMask       bool   `json:"atmophereMask"`
			AtmosphereMaskColor string `json:"atmosphereMaskColor"`
		} `json:"seller"`
		PropsCut string `json:"propsCut"`
		Item     struct {
			ItemId           string        `json:"itemId"`
			Title            string        `json:"title"`
			Subtitle         string        `json:"subtitle"`
			Images           []string      `json:"images"`
			CategoryId       string        `json:"categoryId"`
			RootCategoryId   string        `json:"rootCategoryId"`
			BrandValueId     string        `json:"brandValueId"`
			SkuText          string        `json:"skuText"`
			CountMultiple    []interface{} `json:"countMultiple"`
			ExParams         []interface{} `json:"exParams"`
			CommentCount     string        `json:"commentCount"`
			Favcount         string        `json:"favcount"`
			TaobaoDescUrl    string        `json:"taobaoDescUrl"`
			TmallDescUrl     string        `json:"tmallDescUrl"`
			TaobaoPcDescUrl  string        `json:"taobaoPcDescUrl"`
			ModuleDescUrl    string        `json:"moduleDescUrl"`
			OpenDecoration   bool          `json:"openDecoration"`
			ModuleDescParams struct {
				F  string `json:"f"`
				Id string `json:"id"`
			} `json:"moduleDescParams"`
			H5ModuleDescUrl string `json:"h5moduleDescUrl"`
			CartUrl         string `json:"cartUrl"`
		} `json:"item"`
		Resource struct {
			Entrances struct {
				AskAll struct {
					Icon string `json:"icon"`
					Text string `json:"text"`
					Link string `json:"link"`
				} `json:"askAll"`
			} `json:"entrances"`
		} `json:"resource"`
		Vertical struct {
			AskAll struct {
				AskText    string `json:"askText"`
				AskIcon    string `json:"askIcon"`
				AnswerText string `json:"answerText"`
				AnswerIcon string `json:"answerIcon"`
				LinkUrl    string `json:"linkUrl"`
				Title      string `json:"title"`
				QuestNum   string `json:"questNum"`
				ShowNum    string `json:"showNum"`
				ModelList  []struct {
					AskText         string `json:"askText"`
					AnswerCountText string `json:"answerCountText"`
					FirstAnswer     string `json:"firstAnswer"`
				} `json:"modelList"`
				Model4XList []struct {
					AskText         string `json:"askText"`
					AnswerCountText string `json:"answerCountText"`
					AskIcon         string `json:"askIcon"`
					AskTextColor    string `json:"askTextColor"`
				} `json:"model4XList"`
			} `json:"askAll"`
		} `json:"vertical"`
		Params struct {
			TrackParams struct {
				BrandId    string `json:"brandId"`
				BCType     string `json:"BC_type"`
				CategoryId string `json:"categoryId"`
			} `json:"trackParams"`
		} `json:"params"`
		Props struct {
			GroupProps []interface{} `json:"groupProps"`
		} `json:"props"`
		MockData string `json:"mockData"`
		Rate     struct {
			TotalCount string `json:"totalCount"`
			Invite     struct {
				ShowInvite string `json:"showInvite"`
				InviteText string `json:"inviteText"`
			} `json:"invite"`
			Keywords []struct {
				Attribute string `json:"attribute"`
				Word      string `json:"word"`
				Count     string `json:"count"`
				Type      string `json:"type"`
			} `json:"keywords"`
			RateList []struct {
				Content            string `json:"content"`
				UserName           string `json:"userName"`
				HeadPic            string `json:"headPic"`
				MemberLevel        string `json:"memberLevel"`
				DateTime           string `json:"dateTime"`
				CreateTimeInterval string `json:"createTimeInterval"`
				SkuInfo            string `json:"skuInfo"`
				TmallMemberLevel   string `json:"tmallMemberLevel"`
				IsVip              string `json:"isVip"`
				FeedId             string `json:"feedId"`
			} `json:"rateList"`
			UtFeedId string `json:"utFeedId"`
		} `json:"rate"`
		Props2  []interface{} `json:"props2"`
		SkuBase struct {
			Skus []struct {
				SkuId    string `json:"skuId"`
				PropPath string `json:"propPath"`
			} `json:"skus"`
			Props []struct {
				Pid    string `json:"pid"`
				Name   string `json:"name"`
				Values []struct {
					Vid  string `json:"vid"`
					Name string `json:"name"`
				} `json:"values"`
			} `json:"props"`
		} `json:"skuBase"`
		Delivery struct {
			Postage string `json:"postage"`
			From    string `json:"from"`
		} `json:"delivery"`
		Volume struct {
			SellCount      string `json:"sellCount"`
			VagueSellCount string `json:"vagueSellCount"`
		} `json:"volume"`
	} `json:"data"`
}
