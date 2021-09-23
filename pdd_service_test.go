package go_vesdk

import (
	"github.com/azerothyang/go-vesdk/key"
	"log"
	"testing"
)

func TestPddService_Cats(t *testing.T) {
	type fields struct {
		config Config
	}
	type args struct {
		pddReq *PddCatsReq
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *PddCatsRsp
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name:    "TestPddService_PddCats_Case1",
			fields:  fields{config: Config{VeKey: key.TestVeKey}},
			args:    args{pddReq: &PddCatsReq{ParentCatId: 0}},
			want:    nil,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &PddService{
				config: tt.fields.config,
			}
			got, err := p.Cats(tt.args.pddReq)
			if err != nil {
				log.Fatalf("TestPddService_PddCats错误:%s", err.Error())
			}
			log.Printf("TestPddService_PddCats测试成功😁返回：%s", toJson(got))
		})
	}
}

func TestPddService_GoodsOpt(t *testing.T) {
	type fields struct {
		config Config
	}
	type args struct {
		goodsOptReq *PddGoodsOptReq
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *PddGoodsOptRsp
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name:    "TestPddService_GoodsOpt_Case1",
			fields:  fields{config: Config{VeKey: key.TestVeKey}},
			args:    args{goodsOptReq: &PddGoodsOptReq{ParentOptId: 0}},
			want:    nil,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &PddService{
				config: tt.fields.config,
			}
			got, err := p.GoodsOpt(tt.args.goodsOptReq)
			if err != nil {
				log.Fatalf("TestPddService_GoodsOpt错误:%s", err.Error())
			}
			log.Printf("TestPddService_GoodsOpt测试成功😁返回：%s", toJson(got))
		})
	}
}

func TestPddService_GoodsSearch(t *testing.T) {
	type fields struct {
		config Config
	}
	type args struct {
		pddSearchReq *PddGoodsSearchReq
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *PddGoodsSearchRsp
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name:   "TestPddService_GoodsSearch_Case1",
			fields: fields{config: Config{VeKey: key.TestVeKey}},
			args: args{pddSearchReq: &PddGoodsSearchReq{
				AccessToken:      key.TestPddAccessToken,
				ActivityTags:     nil,
				BlockCatPackages: nil,
				BlockCats:        nil,
				CatId:            0,
				CustomParameters: "{\"uid\": \"0\"}",
				GoodsImgType:     0,
				GoodsSignList:    nil,
				IsBrandGoods:     false,
				//Keyword:          "手机", // 这里需要注意，如果传入了Keyword则需要传入CustomParameters，可以默认使用一个uid为0的备案
				Keyword:          "239881990818", // 这里需要注意，如果传入了Keyword则需要传入CustomParameters，可以默认使用一个uid为0的备案。 615845438434这个id的商品会查出total返回1但是没商品信息的情况
				ListId:           "",
				MerchantType:     0,
				MerchantTypeList: nil,
				OptId:            0,
				Page:             0,
				PageSize:         0,
				Pid:              key.TestPddPid,
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
			p := &PddService{
				config: tt.fields.config,
			}
			got, err := p.GoodsSearch(tt.args.pddSearchReq)
			if err != nil {
				log.Fatalf("TestPddService_GoodsSearch错误:%s", err.Error())
			}
			log.Printf("TestPddService_GoodsSearch测试成功😁返回：%s", toJson(got))
		})
	}
}

func TestPddService_PddGenerate(t *testing.T) {
	type fields struct {
		config Config
	}
	type args struct {
		pddGenerateReq *PddGenerateReq
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *PddGenerateRsp
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name:   "TestPddService_PddGenerate_Case1",
			fields: fields{config: Config{VeKey: key.TestVeKey}},
			args: args{pddGenerateReq: &PddGenerateReq{
				AccessToken:      key.TestPddAccessToken,
				PIdList:          key.TestPddPid,
				CustomParameters: "{\"uid\": \"0\"}",
			}},
			want:    nil,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &PddService{
				config: tt.fields.config,
			}
			got, err := p.PddGenerate(tt.args.pddGenerateReq)
			if err != nil {
				log.Fatalf("TestPddService_GoodsSearch错误:%s", err.Error())
			}
			log.Printf("TestPddService_GoodsSearch测试成功😁返回：%s", toJson(got))
		})
	}
}

func TestPddService_PddAuthQuery(t *testing.T) {
	type fields struct {
		config Config
	}
	type args struct {
		pddAuthQueryReq *PddAuthQueryReq
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *PddAuthQueryRsp
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name:   "TestPddService_PddAuthQuery_Case1",
			fields: fields{config: Config{VeKey: key.TestVeKey}},
			args: args{pddAuthQueryReq: &PddAuthQueryReq{
				AccessToken:      key.TestPddAccessToken,
				Pid:              key.TestPddPid,
				CustomParameters: "{\"uid\": \"0\"}",
			}},
			want:    nil,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &PddService{
				config: tt.fields.config,
			}
			got, err := p.PddAuthQuery(tt.args.pddAuthQueryReq)
			if err != nil {
				log.Fatalf("TestPddService_GoodsSearch错误:%s", err.Error())
			}
			log.Printf("TestPddService_GoodsSearch测试成功😁返回：%s", toJson(got))
		})
	}
}
