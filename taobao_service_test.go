package go_vesdk

import (
	"github.com/azerothyang/go-vesdk/key"
	"go/types"
	"log"
	"testing"
)

func TestTaoBaoService_HCApiAllByItemIds(t1 *testing.T) {
	type fields struct {
		config Config
	}
	type args struct {
		hcParams *HCApiParams
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *HCApiAllRsp
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name:   "TestTaoBaoService_HCApiAllByItemIds",
			fields: fields{config: Config{VeKey: key.TestVeKey}},
			args: args{hcParams: &HCApiParams{
				Para:        "653217285015,619337473117",
				Pid:         "",
				RelationId:  0,
				SpecialId:   0,
				ExternalId:  "",
				Detail:      1,
				ActivityId:  "",
				DeepCoupon:  1,
				ShopCoupon:  0,
				NoTkl:       0,
				NoShortLink: 0,
				CouponId:    "",
				TklType:     0,
				Xid:         "",
				GetTopnRate: "",
				UcrowdId:    "",
			}},
			want:    nil,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t1.Run(tt.name, func(t1 *testing.T) {
			t := &TaoBaoService{
				config: tt.fields.config,
			}
			got, err := t.HCApiAllByItemIds(tt.args.hcParams)
			if err != nil {
				log.Fatalf("TestTaoBaoService_HCApiAllByItemIds错误:%s", err.Error())
			}
			log.Printf("TestTaoBaoService_HCApiAllByItemIds测试成功😁返回：%s", toJson(got))
		})
	}
}

func TestTaoBaoService_HCApiOne(t1 *testing.T) {
	type fields struct {
		config Config
	}
	type args struct {
		hcParams *HCApiParams
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *HCApiRsp
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name:   "TestTaoBaoService_HCApiOne_case_1",
			fields: fields{config: Config{VeKey: key.TestVeKey}},
			args: args{hcParams: &HCApiParams{
				Para:        "2👈，7lE9Xnhpvlv信 https://m.tb.cn/h.fYwVbcV?sm=c90a70  太空熊猫联名nasa旗舰店官网秋季外套小熊男情侣休闲夹克nasa外套",
				Pid:         "",
				RelationId:  0,
				SpecialId:   0,
				ExternalId:  "",
				Detail:      1,
				ActivityId:  "",
				DeepCoupon:  1,
				ShopCoupon:  0,
				NoTkl:       0,
				NoShortLink: 0,
				CouponId:    "",
				TklType:     0,
				Xid:         "",
				GetTopnRate: "",
				UcrowdId:    "",
			}},
			want:    nil,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t1.Run(tt.name, func(t1 *testing.T) {
			t := &TaoBaoService{
				config: tt.fields.config,
			}
			got, err := t.HCApiOne(tt.args.hcParams)
			if err != nil {
				log.Fatalf("TestTaoBaoService_HCApiOne错误:%s", err.Error())
			}
			log.Printf("TestTaoBaoService_HCApiOne测试成功😁返回：%s", toJson(got))
		})
	}
}

func TestTaoBaoService_OrderDetails(t1 *testing.T) {
	type fields struct {
		config Config
	}
	type args struct {
		orderDetailParams *OrderDetailsParams
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *OrderDetailsRsp
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name:   "TestTaoBaoService_OrderDetails_case_1",
			fields: fields{config: Config{VeKey: key.TestVeKey}},
			args: args{orderDetailParams: &OrderDetailsParams{
				StartTime:     "2021-09-07 18:21:16",
				EndTime:       "2021-09-07 18:30:16",
				QueryType:     QueryTypeUpdateTime,
				PositionIndex: "",
				PageNo:        0,
				PageSize:      0,
				MemberType:    0,
				TKStatus:      0,
				JumpType:      0,
				OrderScene:    1,
			}},
		},
	}
	for _, tt := range tests {
		t1.Run(tt.name, func(t1 *testing.T) {
			t := &TaoBaoService{
				config: tt.fields.config,
			}
			got, err := t.OrderDetails(tt.args.orderDetailParams)
			if err != nil {
				log.Fatalf("TestTaoBaoService_OrderDetails错误:%s", err.Error())
			}
			log.Printf("TestTaoBaoService_OrderDetails测试成功😁返回：%s", toJson(got))
		})
	}
}

func TestTaoBaoService_SuperSearch(t1 *testing.T) {
	type fields struct {
		config Config
	}
	type args struct {
		superSearchReq *SuperSearchReq
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    types.Nil
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name:   "TestTaoBaoService_SuperSearch_case1",
			fields: fields{config: Config{VeKey: key.TestVeKey}},
			args: args{superSearchReq: &SuperSearchReq{
				Para:              "2👈，7lE9Xnhpvlv信 https://m.tb.cn/h.fYwVbcV?sm=c90a70  太空熊猫联名nasa旗舰店官网秋季外套小熊男情侣休闲夹克nasa外套",
				Cat:               "",
				Pid:               "mm_708300086_1360150325_110124300439",
				RelationId:        "",
				Page:              0,
				PageSize:          0,
				ForceIndex:        "",
				Coupon:            "",
				StartPrice:        "",
				EndPrice:          "",
				StartTkRate:       "",
				EndTkRate:         "",
				IsOverseas:        0,
				IsTmall:           0,
				Sort:              "",
				OnlySearch:        0,
				Similar:           0,
				SearchItem:        0,
				Ip:                "",
				FreeShip:          0,
				Npx:               2,
				ItemId:            0,
				Virtual:           2,
				Presale:           0,
				IncludePayRate30:  0,
				IncludeGoodRate:   0,
				EndKaTkRate:       0,
				StartKaTkRate:     0,
				DeviceValue:       "",
				DeviceEncrypt:     "",
				DeviceType:        "",
				LockRateEndTime:   0,
				LockRateStartTime: 0,
				SellerIds:         "",
				CityCode:          "",
				Latitude:          0,
				Longitude:         0,
				SpecialId:         "",
				NeedPrepay:        0,
				TkLink:            1,
			}},
			want:    types.Nil{},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t1.Run(tt.name, func(t1 *testing.T) {
			t := &TaoBaoService{
				config: tt.fields.config,
			}
			got, err := t.SuperSearch(tt.args.superSearchReq)
			if err != nil {
				log.Fatalf("TestTaoBaoService_SuperSearch错误:%s", err.Error())
			}
			log.Printf("TestTaoBaoService_SuperSearch测试成功😁-搜索返回：%s", toJson(got))

		})
	}
}

func TestTaoBaoService_OrderPunish(t1 *testing.T) {
	type fields struct {
		config Config
	}
	type args struct {
		orderPunishReq *OrderPunishReq
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *OrderPunishRsp
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name:   "TestTaoBaoService_OrderPunish_case1",
			fields: fields{config: Config{VeKey: key.TestVeKey}},
			args: args{orderPunishReq: &OrderPunishReq{
				StartTime:  "2021-09-01",
				TbTradeId:  "",
				RelationId: "",
				Span:       30,
				PageNo:     0,
				PageSize:   0,
				AdzoneId:   "",
				SiteId:     "",
			}},
			want:    nil,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t1.Run(tt.name, func(t1 *testing.T) {
			t := &TaoBaoService{
				config: tt.fields.config,
			}
			got, err := t.OrderPunish(tt.args.orderPunishReq)
			if err != nil {
				log.Fatalf("TestTaoBaoService_OrderDetails错误:%s", err.Error())
			}
			log.Printf("TestTaoBaoService_OrderDetails测试成功😁返回：%s", toJson(got))
		})
	}
}

func TestTaoBaoService_RefundOrder(t1 *testing.T) {
	type fields struct {
		config Config
	}
	type args struct {
		refundOrderReq *RefundOrderReq
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *RefundOrderRsp
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name:   "TestTaoBaoService_RefundOrder_case1",
			fields: fields{config: Config{VeKey: key.TestVeKey}},
			args: args{refundOrderReq: &RefundOrderReq{
				StartTime:  "2021-09-11 00:00:00",
				PageNo:     1,
				PageSize:   30,
				SearchType: 5,
				RefundType: 1,
				BizType:    1,
			}},
			want:    nil,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t1.Run(tt.name, func(t1 *testing.T) {
			t := &TaoBaoService{
				config: tt.fields.config,
			}
			got, err := t.RefundOrder(tt.args.refundOrderReq)
			if err != nil {
				log.Fatalf("TestTaoBaoService_RefundOrder错误:%s", err.Error())
			}
			log.Printf("TestTaoBaoService_RefundOrder测试成功😁返回：%s", toJson(got))
		})
	}
}

func TestTaoBaoService_TbProductDetail(t1 *testing.T) {
	type fields struct {
		config Config
	}
	type args struct {
		tbProductDetailReq *TbProductDetailReq
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *TbProductDetailRsp
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name:   "TestTaoBaoService_TbProductDetail_case1",
			fields: fields{config: Config{VeKey: key.TestVeKey}},
			args: args{tbProductDetailReq: &TbProductDetailReq{
				Id:           646011446519,
				OnlyShopInfo: 0,
				OnlySimple:   0,
				GetCatName:   0,
			}},
			want:    nil,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t1.Run(tt.name, func(t1 *testing.T) {
			t := &TaoBaoService{
				config: tt.fields.config,
			}
			got, err := t.TbProductDetail(tt.args.tbProductDetailReq)
			if err != nil {
				log.Fatalf("TestTaoBaoService_RefundOrder错误:%s", err.Error())
			}
			log.Printf("TestTaoBaoService_RefundOrder测试成功😁返回：%s", toJson(got))
		})
	}
}

func TestTaoBaoService_TbSimpleProductDetail(t1 *testing.T) {
	type fields struct {
		config Config
	}
	type args struct {
		tbSimpleDetailReq *TbSimpleDetailReq
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *TbSimpleDetailRsp
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name:   "TestTaoBaoService_TbSimpleProductDetail_case1",
			fields: fields{config: Config{VeKey: key.TestVeKey}},
			args: args{tbSimpleDetailReq: &TbSimpleDetailReq{
				Id:           646011446519,
				OnlyShopInfo: 0,
				GetCatName:   0,
			}},
			want:    nil,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t1.Run(tt.name, func(t1 *testing.T) {
			t := &TaoBaoService{
				config: tt.fields.config,
			}
			got, err := t.TbSimpleProductDetail(tt.args.tbSimpleDetailReq)
			if err != nil {
				log.Fatalf("TestTaoBaoService_TbSimpleProductDetail错误:%s", err.Error())
			}
			log.Printf("TestTaoBaoService_TbSimpleProductDetail测试成功😁返回：%s", toJson(got))
		})
	}
}
