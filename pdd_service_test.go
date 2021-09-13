package go_vesdk

import (
	"github.com/azerothyang/go-vesdk/key"
	"log"
	"testing"
)

func TestPDDService_Cats(t *testing.T) {
	type fields struct {
		config Config
	}
	type args struct {
		pddReq *PDDCatsReq
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *PDDCatsRsp
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name:    "TestPDDService_PDDCats",
			fields:  fields{config: Config{VeKey: key.TestVeKey}},
			args:    args{pddReq: &PDDCatsReq{ParentCatId: 0}},
			want:    nil,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &PDDService{
				config: tt.fields.config,
			}
			got, err := p.Cats(tt.args.pddReq)
			if err != nil {
				log.Fatalf("TestPDDService_PDDCats错误:%s", err.Error())
			}
			log.Printf("TestPDDService_PDDCats测试成功😁返回：%s", toJson(got))
		})
	}
}

func TestPDDService_GoodsOpt(t *testing.T) {
	type fields struct {
		config Config
	}
	type args struct {
		goodsOptReq *PDDGoodsOptReq
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *PDDGoodsOptRsp
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name:    "TestPDDService_GoodsOpt_case1",
			fields:  fields{config: Config{VeKey: key.TestVeKey}},
			args:    args{goodsOptReq: &PDDGoodsOptReq{ParentOptId: 0}},
			want:    nil,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &PDDService{
				config: tt.fields.config,
			}
			got, err := p.GoodsOpt(tt.args.goodsOptReq)
			if err != nil {
				log.Fatalf("TestPDDService_GoodsOpt错误:%s", err.Error())
			}
			log.Printf("TestPDDService_GoodsOpt测试成功😁返回：%s", toJson(got))
		})
	}
}

func TestPDDService_GoodsSearch(t *testing.T) {
	type fields struct {
		config Config
	}
	type args struct {
		pddSearchReq *PDDGoodsSearchReq
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *PDDGoodsSearchRsp
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name:    "TestPDDService_GoodsSearch_case1",
			fields:  fields{config: Config{VeKey: key.TestVeKey}},
			args:    args{pddSearchReq: &PDDGoodsSearchReq{
				AccessToken:      key.TestPDDAccessToken,
				ActivityTags:     nil,
				BlockCatPackages: nil,
				BlockCats:        nil,
				CatId:            0,
				CustomParameters: "",
				GoodsImgType:     0,
				GoodsSignList:    nil,
				IsBrandGoods:     false,
				Keyword:          "",
				ListId:           "",
				MerchantType:     0,
				MerchantTypeList: nil,
				OptId:            0,
				Page:             0,
				PageSize:         0,
				Pid:              key.TestPDDPid,
				RangeList:        nil,
				SortType:         0,
				UseCustomized:    false,
				WithCoupon:       false,
			}},
			want:    nil,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &PDDService{
				config: tt.fields.config,
			}
			got, err := p.GoodsSearch(tt.args.pddSearchReq)
			if err != nil {
				log.Fatalf("TestPDDService_GoodsOpt错误:%s", err.Error())
			}
			log.Printf("TestPDDService_GoodsOpt测试成功😁返回：%s", toJson(got))
		})
	}
}