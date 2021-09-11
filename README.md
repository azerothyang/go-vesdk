# go-vesdk
Golang维易相关淘宝，京东，拼多多，唯品会sdk🐲

## 快速开始
```
package main

import (
	"fmt"
	"github.com/azerothyang/go-vesdk"
	"log"
)

func main() {
	jdSrv := go_vesdk.NewJDService(go_vesdk.Config{VeKey: "V0000"})
	cats, err := jdSrv.Category(0, 0)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Printf("类目名称：%v\n", *cats)

	taobaoSrv := go_vesdk.NewTaoBaoService(go_vesdk.Config{VeKey: "V0000"})
	hcApiRsp, err := taobaoSrv.HCApiOne(&go_vesdk.HCApiParams{
		Para:        "527484822054",
	})
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Printf("高佣转链:%v\n", *hcApiRsp)
}
```

## 目录说明
### 淘客相关
- OrderDetails: 订单查询
- HCApiOne: 高佣转链
- HCApiAllByItemIds: 多个商品id批量高佣转链
- SuperSearch: 超级搜索接口
- OrderPunish: 淘宝客处罚订单查询API
- RefundOrder: 淘客维权订单接口-支持渠道订单和会员订单维权

### 京东相关
- JDSearch: 关键词商品查询接口
- Category: 新版分类接口category
- PromotionGoodsInfo: 获取推广商品信息接口
- PromByUid: 通过unionId获取推广链接
- OrderRowQuery: 订单行查询接口
- BigField: 大字段商品详情查询接口
- JDMaterialQuery: 物料商品查询