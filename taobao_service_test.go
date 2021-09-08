package go_vesdk

import (
	"log"
	"testing"
)

const (
	TestVeKey = "V" // todo 提交时要修改
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
			fields: fields{config: Config{VeKey: TestVeKey}},
			args: args{hcParams: &HCApiParams{
				Para:        "653217285015,619337473117",
				Pid:         "",
				RelationId:  "",
				SpecialId:   "",
				ExternalId:  "",
				Detail:      "",
				ActivityId:  "",
				DeepCoupon:  "",
				ShopCoupon:  "",
				NoTkl:       "",
				NoShortLink: "",
				CouponId:    "",
				TklType:     "",
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
			fields: fields{config: Config{VeKey: TestVeKey}},
			args: args{hcParams: &HCApiParams{
				Para:        "650096340209",
				Pid:         "",
				RelationId:  "",
				SpecialId:   "",
				ExternalId:  "",
				Detail:      "",
				ActivityId:  "",
				DeepCoupon:  "",
				ShopCoupon:  "",
				NoTkl:       "",
				NoShortLink: "",
				CouponId:    "",
				TklType:     "",
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
			fields: fields{config: Config{VeKey: TestVeKey}},
			args: args{orderDetailParams: &OrderDetailsParams{
				StartTime:     "2021-09-07 18:21:16",
				EndTime:       "2021-09-07 18:30:16",
				QueryType:     0,
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
