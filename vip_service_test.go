package go_vesdk

import (
	"github.com/azerothyang/go-vesdk/key"
	"log"
	"testing"
)

func TestVipService_Search(t *testing.T) {
	type fields struct {
		config Config
	}
	type args struct {
		vipSearchReq *VipSearchReq
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *VipSearchRsp
		wantErr bool
	}{
		// TODO: Add test cases.
		{name: "TestVipService_Search_Case1",
			fields: fields{config: Config{VeKey: key.TestVeKey}},
			args: args{vipSearchReq: &VipSearchReq{
				AccessToken:                 key.TestVipAccessToken,
				Keyword:                     "衣服",
				FieldName:                   "",
				Order:                       "",
				Page:                        0,
				PageSize:                    0,
				RequestId:                   "",
				PriceStart:                  "",
				PriceEnd:                    "",
				QueryReputation:             true,
				QueryStoreServiceCapability: true,
				QueryStock:                  true,
				QueryActivity:               true,
				QueryPrepay:                 true,
				CommonParams:                VipCommonParams{},
				VendorCode:                  "",
				ChanTag:                     "",
				QueryExclusiveCoupon:        true,
				QueryCpsInfo:                3,
				Mobile:                      "",
				Virtual:                     NoCheckVirtual,
			}},
			want:    nil,
			wantErr: false,
		},

		{name: "TestVipService_Search_Case2",
			fields: fields{config: Config{VeKey: key.TestVeKey}},
			args: args{vipSearchReq: &VipSearchReq{
				AccessToken:                 key.TestVipAccessToken,
				Keyword:                     "衣服",
				FieldName:                   "",
				Order:                       "",
				Page:                        2,
				PageSize:                    5,
				RequestId:                   "",
				PriceStart:                  "",
				PriceEnd:                    "",
				QueryReputation:             false,
				QueryStoreServiceCapability: false,
				QueryStock:                  false,
				QueryActivity:               false,
				QueryPrepay:                 false,
				CommonParams:                VipCommonParams{},
				VendorCode:                  "",
				ChanTag:                     "",
				QueryExclusiveCoupon:        false,
				QueryCpsInfo:                0,
				Mobile:                      "18829301023",
				Virtual:                     CheckVirtual,
			}},
			want:    nil,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			v := &VipService{
				config: tt.fields.config,
			}
			got, err := v.Search(tt.args.vipSearchReq)
			if err != nil {
				log.Fatalf("TestVipService_Search错误:%s", err.Error())
			}
			log.Printf("TestVipService_Search测试成功😁返回：%s", toJson(got))
		})
	}
}
