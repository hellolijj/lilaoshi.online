package dto

type OrderCreateParams struct {
	Data []ShoppingCarGoodItem `json:"data"`
	Uid string `json:"uid"`
} 

type ShoppingCarGoodItem struct {
	Count int `json:"count"`
	Id int `json:"id"`
	Price string `json:"price"`
}
