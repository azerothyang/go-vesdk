package go_vesdk

type Virtual uint

// ToInt64 转为int64
func (v Virtual) ToInt64() int64 {
	return int64(v)
}

// ToInt 转为int
func (v Virtual) ToInt() int {
	return int(v)
}

const (
	veUrl = "http://apivip.vephp.com"
)

const (
	NoCheckVirtual Virtual = iota // 默认0表示不检查虚拟类商品
	RefuseVirtual                 // 1表示检测并拒绝虚拟类商品搜索（此时接口error为91）
	CheckVirtual                  // 2表示只检测不过滤，这时返回值中增加一个字段is_virtual（值0表示不是虚拟产品，1表示是虚拟产品）
)
