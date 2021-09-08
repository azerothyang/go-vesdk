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
	fmt.Printf("类目名称1：%s\n", cats.Data[0].Name)
}
```

## 目录说明
### 淘客相关
- OrderDetails: 订单查询
- HCApiOne: 高佣转链
- HCApiAllByItemIds: 多个商品id批量高佣转链

### 京东相关
- JDSearch: 关键词商品查询接口
- Category: 新版分类接口category
- PromotionGoodsInfo: 获取推广商品信息接口
- PromByUid: 通过unionId获取推广链接
- OrderRowQuery: 订单行查询接口