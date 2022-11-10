package dto

type ShoppingCarGoodsResp struct {
	Current int64 `json:"current"`
	PageSize int64 `json:"pageSize"`
	Price string `json:"price"`
	GoodsTotal int `json:"goodsTotal"`
	Total int `json:"total"`
	Success bool `json:"success"`
	Data interface{} `json:"data"`
}


type ShoppingCarOpParams struct {
	Op string `json:"op"`
	GoodId int `json:"goodId"`
	Uid string `json:"uid"`
	Num int `json:"num"`
}
