package go_vesdk

import (
	"github.com/azerothyang/go-vesdk/key"
	"log"
	"net/url"
	"testing"
)

func TestJDService_JDSearch(t *testing.T) {
	type fields struct {
		config Config
	}
	type args struct {
		jdParams *JDSearchParams
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *JDSearchRsp
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name:   "TestJDService_JDSearch_Case1",
			fields: fields{config: Config{VeKey: key.TestVeKey}},
			args: args{jdParams: &JDSearchParams{
				SkuIds:  "",
				Keyword: "电子产品",
				Cid1:    0,
				Cid2:    0,
				Cid3:    0,
				JxFlags: "",
				Virtual: "",
			}},
			want:    nil,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			j := &JDService{
				config: tt.fields.config,
			}
			got, err := j.JDSearch(tt.args.jdParams)
			if err != nil {
				log.Fatalf("TestJDService_JDSearch错误:\n%s", err.Error())
			}
			log.Printf("TestJDService_JDSearch测试成功😁返回：\n%s\n", toJson(got))
		})
	}
}

func TestJDService_Category(t *testing.T) {
	type fields struct {
		config Config
	}
	type args struct {
		parentId int64
		grade    int64
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *JDCategoryRsp
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name:   "TestJDService_Category_Case1",
			fields: fields{config: Config{VeKey: key.TestVeKey}},
			args: args{
				parentId: 0, // 上一级类目id
				grade:    0, // 当前查询第几级，类目级别(类目级别 0，1，2 代表一、二、三级类目)
			},
			want:    nil,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			j := &JDService{
				config: tt.fields.config,
			}
			got, err := j.Category(tt.args.parentId, tt.args.grade)
			if err != nil {
				log.Fatalf("TestJDService_Category错误:\n%s", err.Error())
			}
			log.Printf("TestJDService_Category测试成功😁返回：\n%s\n", toJson(got))
		})
	}
}

func TestJDService_PromotionGoodsInfo(t *testing.T) {
	type fields struct {
		config Config
	}
	type args struct {
		skuIds []string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *JDPromotionGoodsInfoRsp
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name:    "TestJDService_PromotionGoodsInfo_Case1",
			fields:  fields{config: Config{VeKey: key.TestVeKey}},
			args:    args{skuIds: []string{"100016777664", "4718515"}},
			want:    nil,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			j := &JDService{
				config: tt.fields.config,
			}
			got, err := j.PromotionGoodsInfo(tt.args.skuIds)
			if err != nil {
				log.Fatalf("TestJDService_PromotionGoodsInfo错误:\n%s", err.Error())
			}
			log.Printf("TestJDService_PromotionGoodsInfo测试成功😁返回：\n%s\n", toJson(got))
		})
	}
}

func TestJDService_PromByUid(t *testing.T) {
	type fields struct {
		config Config
	}
	type args struct {
		jdPromotionCodeReq *JDPromotionCodeReq
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
		{
			name:   "TestJDService_PromByUid_case1",
			fields: fields{config: Config{VeKey: key.TestVeKey}},
			args: args{jdPromotionCodeReq: &JDPromotionCodeReq{
				MaterialId:    "http://item.jd.com/67194752420.html",
				UnionId:       123456,
				PositionId:    123456,
				Pid:           "",
				CouponUrl:     url.QueryEscape("https://coupon.m.jd.com/coupons/show.action?linkKey=AAROH_xIpeffAs_-naABEFoexBz59GNdTpNcnZHRrf6jFZPPePXB7CZv6gOi2XjDrsmtRs_nhIbVHftYFqvYfdxcSNGwDQ"),
				SubUnionId:    "",
				ChainType:     0,
				GiftCouponKey: "",
				ChannelId:     0,
			}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			j := &JDService{
				config: tt.fields.config,
			}
			got, err := j.PromByUid(tt.args.jdPromotionCodeReq)
			if err != nil {
				log.Fatalf("TestJDService_PromByUid错误:\n%s", err.Error())
			}
			log.Printf("TestJDService_PromByUid测试成功😁返回：\n%s\n", toJson(got))
		})
	}
}

func TestJDService_OrderRowQuery(t *testing.T) {
	type fields struct {
		config Config
	}
	type args struct {
		orderReq *JDOrderReq
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *JDOrderRsp
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name:   "TestJDService_OrderRowQuery_Case1",
			fields: fields{config: Config{VeKey: key.TestVeKey}},
			args: args{&JDOrderReq{
				PageIndex:    1,
				PageSize:     20,
				Type:         3,
				StartTime:    "2021-06-18 20:23:00",
				EndTime:      "2021-06-18 21:23:00",
				ChildUnionId: 0,
				Key:          key.TestJDKey, // todo 删除掉
				Fields:       "",
			}},
			want:    nil,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			j := &JDService{
				config: tt.fields.config,
			}
			got, err := j.OrderRowQuery(tt.args.orderReq)
			if err != nil {
				log.Fatalf("TestJDService_PromByUid错误:\n%s", err.Error())
			}
			log.Printf("TestJDService_PromByUid测试成功😁返回：\n%s\n", toJson(got))
		})
	}
}

func TestJDService_BigField(t *testing.T) {
	type fields struct {
		config Config
	}
	type args struct {
		skuIds []string
		fields []string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *JDBigFieldRsp
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name:    "TestJDService_BigField_Case1",
			fields:  fields{config: Config{VeKey: key.TestVeKey}},
			args:    args{skuIds: []string{"68516192765", "100011431721"}, fields: []string{"imageInfo"}},
			want:    nil,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			j := &JDService{
				config: tt.fields.config,
			}
			got, err := j.BigField(tt.args.skuIds, tt.args.fields)
			if err != nil {
				log.Fatalf("TestJDService_BigField错误:\n%s", err.Error())
			}
			log.Printf("TestJDService_BigField测试成功😁返回：\n%s\n", toJson(got))
		})
	}
}

func TestJDService_JDMaterialQuery(t *testing.T) {
	type fields struct {
		config Config
	}
	type args struct {
		jdMaterialQueryReq *JDMaterialQueryReq
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *JDMaterialQueryRsp
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name: "TestJDService_JDMaterialQuery_Case1",
			fields: fields{config: Config{
				VeKey: key.TestVeKey,
			}},
			args: args{jdMaterialQueryReq: &JDMaterialQueryReq{
				EliteId:   1,
				PageIndex: 1,
				PageSize:  10,
			}},
			want:    nil,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			j := &JDService{
				config: tt.fields.config,
			}
			got, err := j.JDMaterialQuery(tt.args.jdMaterialQueryReq)
			if err != nil {
				log.Fatalf("TestJDService_BigField错误:\n%s", err.Error())
			}
			log.Printf("TestJDService_BigField测试成功😁返回：\n%s\n", toJson(got))
		})
	}
}
