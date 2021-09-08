package go_vesdk

// JDSearchParams 关键词商品查询接口：
// 主要参数：
//skuIds 可选，默认为空。可传多个京东SKU（即产品ID）SKU间用逗号相隔。
//keyword 可选，默认为空，关键字，如“手机”
//cid1，cid2，cid3，可选，分类ID
//jxFlags 可选，传参形式可以是 [2,3]或2,3。
//新增参数：
//virtual参数：可选值0，1，2，用于指定是否过滤或检测虚拟类产品，苹果上架APP时需要用到，可临时开启，参数值：默认0表示不检查虚拟类商品，1表示检测并拒绝虚拟类商品搜索（此时接口error为91），2表示只检测不过滤，这时返回值中增加一个字段is_virtual（值0表示不是虚拟产品，1表示是虚拟产品）
type JDSearchParams struct {
	SkuIds  string `json:"skuIds"`
	Keyword string `json:"keyword"`
	Cid1    int    `json:"cid1"`
	Cid2    int    `json:"cid2"`
	Cid3    int    `json:"cid3"`
	JxFlags string `json:"jxFlags"`
	Virtual string `json:"virtual"`
}

// JDSearchRsp 关键词商品查询接口返回
// 含义参考：https://union.jd.com/openplatform/api/10420
type JDSearchRsp struct {
	Error      string `json:"error"`
	Msg        string `json:"msg"`
	IsVirtual  int    `json:"is_virtual"`
	TotalCount string `json:"totalCount"`
	Data       []struct {
		BrandCode    string `json:"brandCode"`
		BrandName    string `json:"brandName"`
		CategoryInfo struct {
			Cid1     string `json:"cid1"`
			Cid1Name string `json:"cid1Name"`
			Cid2     string `json:"cid2"`
			Cid2Name string `json:"cid2Name"`
			Cid3     string `json:"cid3"`
			Cid3Name string `json:"cid3Name"`
		} `json:"categoryInfo"`
		Comments       string `json:"comments"`
		CommissionInfo struct {
			Commission          string `json:"commission"`
			CommissionShare     string `json:"commissionShare"`
			CouponCommission    string `json:"couponCommission"`
			EndTime             string `json:"endTime"`
			IsLock              string `json:"isLock"`
			PlusCommissionShare string `json:"plusCommissionShare"`
			StartTime           string `json:"startTime"`
		} `json:"commissionInfo"`
		CouponInfo struct {
			CouponList []interface{} `json:"couponList"`
		} `json:"couponInfo"`
		DeliveryType      string `json:"deliveryType"`
		ForbidTypes       []int  `json:"forbidTypes"`
		GoodCommentsShare string `json:"goodCommentsShare"`
		ImageInfo         struct {
			ImageList []struct {
				Url string `json:"url"`
			} `json:"imageList"`
			WhiteImage string `json:"whiteImage,omitempty"`
		} `json:"imageInfo"`
		InOrderComm30Days  string `json:"inOrderComm30Days"`
		InOrderCount30Days string `json:"inOrderCount30Days"`
		IsHot              string `json:"isHot"`
		IsJdSale           string `json:"isJdSale"`
		MaterialUrl        string `json:"materialUrl"`
		Owner              string `json:"owner"`
		PinGouInfo         struct {
		} `json:"pinGouInfo"`
		PingGouInfo struct {
		} `json:"pingGouInfo"`
		PriceInfo struct {
			LowestCouponPrice string `json:"lowestCouponPrice"`
			LowestPrice       string `json:"lowestPrice"`
			LowestPriceType   string `json:"lowestPriceType"`
			Price             string `json:"price"`
		} `json:"priceInfo"`
		ShopInfo struct {
			AfsFactorScoreRankGrade     string `json:"afsFactorScoreRankGrade,omitempty"`
			AfterServiceScore           string `json:"afterServiceScore,omitempty"`
			CommentFactorScoreRankGrade string `json:"commentFactorScoreRankGrade,omitempty"`
			ScoreRankRate               string `json:"scoreRankRate,omitempty"`
			ShopId                      string `json:"shopId"`
			ShopLabel                   string `json:"shopLabel"`
			ShopLevel                   string `json:"shopLevel"`
			ShopName                    string `json:"shopName"`
			UserEvaluateScore           string `json:"userEvaluateScore,omitempty"`
		} `json:"shopInfo"`
		SkuId     string        `json:"skuId"`
		SkuName   string        `json:"skuName"`
		Spuid     string        `json:"spuid"`
		VideoInfo []interface{} `json:"videoInfo"`
		EliteType []int         `json:"eliteType,omitempty"`
	} `json:"data"`
}

// JDCategoryRsp 京东分类返回
type JDCategoryRsp struct {
	Error string `json:"error"`
	Msg   string `json:"msg"`
	Data  []struct {
		Grade    string `json:"grade"`
		Name     string `json:"name"`
		Id       string `json:"id"`
		ParentId string `json:"parentId"`
	} `json:"data"`
}

type JDPromotionGoodsInfoRsp struct {
	Error string `json:"error"`
	Msg   string `json:"msg"`
	Data  []struct {
		UnitPrice         string `json:"unitPrice"`
		MaterialUrl       string `json:"materialUrl"`
		EndDate           string `json:"endDate"`
		IsFreeFreightRisk string `json:"isFreeFreightRisk"`
		IsFreeShipping    string `json:"isFreeShipping"`
		CommisionRatioWl  string `json:"commisionRatioWl"` // 返回类似0.1这种表示 0.1%
		CommisionRatioPc  string `json:"commisionRatioPc"`
		ImgUrl            string `json:"imgUrl"`
		Vid               string `json:"vid"`
		CidName           string `json:"cidName"`
		WlUnitPrice       string `json:"wlUnitPrice"`
		Cid2Name          string `json:"cid2Name"`
		IsSeckill         string `json:"isSeckill"`
		Cid2              string `json:"cid2"`
		Cid3Name          string `json:"cid3Name"`
		InOrderCount      string `json:"inOrderCount"`
		Cid3              string `json:"cid3"`
		ShopId            string `json:"shopId"`
		IsJdSale          string `json:"isJdSale"`
		GoodsName         string `json:"goodsName"`
		SkuId             string `json:"skuId"`
		StartDate         string `json:"startDate"`
		Cid               string `json:"cid"`
	} `json:"data"`
}

type JDPromotionCodeReq struct {
	MaterialId    string `json:"materialId"`
	UnionId       int64  `json:"unionId"`
	PositionId    int64  `json:"positionId"`
	Pid           string `json:"pid"`
	CouponUrl     string `json:"couponUrl"`
	SubUnionId    string `json:"subUnionId"`
	ChainType     int    `json:"chainType"`
	GiftCouponKey string `json:"giftCouponKey"`
	ChannelId     int64  `json:"channelId"`
}

type JDPromByUidRsp struct {
	Error string `json:"error"`
	Msg   string `json:"msg"`
	Data  struct {
		ShortURL string `json:"shortURL"`
	} `json:"data"`
}

type JDOrderReq struct {
	PageIndex    int    `json:"pageIndex"`
	PageSize     int    `json:"pageSize"`
	Type         int    `json:"type"`
	StartTime    string `json:"startTime"`
	EndTime      string `json:"endTime"`
	ChildUnionId int64  `json:"childUnionId"`
	Key          string `json:"key"`
	Fields       string `json:"fields"`
}

type JDOrderRsp struct {
	Error   string `json:"error"`
	Msg     string `json:"msg"`
	HasMore bool   `json:"hasMore"`
	Data    []struct {
		ActualCosPrice string `json:"actualCosPrice"`
		ActualFee      string `json:"actualFee"`
		BalanceExt     string `json:"balanceExt"`
		CategoryInfo   struct {
			Cid1 string `json:"cid1"`
			Cid2 string `json:"cid2"`
			Cid3 string `json:"cid3"`
		} `json:"categoryInfo"`
		ChannelId           string `json:"channelId"`
		Cid1                string `json:"cid1"`
		Cid2                string `json:"cid2"`
		Cid3                string `json:"cid3"`
		CommissionRate      string `json:"commissionRate"`
		CpActId             string `json:"cpActId"`
		EstimateCosPrice    string `json:"estimateCosPrice"`
		EstimateFee         string `json:"estimateFee"`
		ExpressStatus       string `json:"expressStatus"`
		Ext1                string `json:"ext1"`
		FinalRate           string `json:"finalRate"`
		FinishTime          string `json:"finishTime"`
		GiftCouponKey       string `json:"giftCouponKey"`
		GiftCouponOcsAmount string `json:"giftCouponOcsAmount"`
		GoodsInfo           struct {
			MainSkuId string `json:"mainSkuId"`
			ProductId string `json:"productId"`
			ShopId    string `json:"shopId"`
		} `json:"goodsInfo"`
		Id             string `json:"id"`
		ModifyTime     string `json:"modifyTime"`
		OrderEmt       string `json:"orderEmt"`
		OrderId        string `json:"orderId"`
		OrderTime      string `json:"orderTime"`
		ParentId       string `json:"parentId"`
		PayMonth       string `json:"payMonth"`
		Pid            string `json:"pid"`
		Plus           string `json:"plus"`
		PopId          string `json:"popId"`
		PositionId     string `json:"positionId"`
		Price          string `json:"price"`
		ProPriceAmount string `json:"proPriceAmount"`
		Rid            string `json:"rid"`
		SiteId         string `json:"siteId"`
		SkuFrozenNum   string `json:"skuFrozenNum"`
		SkuId          string `json:"skuId"`
		SkuName        string `json:"skuName"`
		SkuNum         string `json:"skuNum"`
		SkuReturnNum   string `json:"skuReturnNum"`
		SubSideRate    string `json:"subSideRate"`
		SubUnionId     string `json:"subUnionId"`
		SubsidyRate    string `json:"subsidyRate"`
		TraceType      string `json:"traceType"`
		UnionAlias     string `json:"unionAlias"`
		UnionId        string `json:"unionId"`
		UnionRole      string `json:"unionRole"`
		UnionTag       string `json:"unionTag"`
		ValidCode      string `json:"validCode"`
	} `json:"data"`
}
