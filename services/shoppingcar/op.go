package shoppingcar

import (
	"cookie-shop-api/models"
	"cookie-shop-api/models/dto"
	"errors"
	"github.com/beego/beego/v2/client/orm"
)

func Op(params dto.ShoppingCarOpParams) (bool, error){
	item, err := models.GetShoppingcarByUId(params.Uid)
	if err != nil && err != orm.ErrNoRows {
		return false, err
	}

	good, err := models.GetGoodById(params.GoodId)
	if err != nil {
		return false, errors.New("查询good失败")
	}

	if params.Op == "sub" {
		if item == nil {
			return false, errors.New("购物车数量不支持减少操作")
		}

		if item.Total - params.Num < 0 {
			return false, errors.New("购物车数量不支持减少操作")
		}

		// 再原有数据上修改
		if err := item.SubGood(good, params.Num); err != nil {
			return false, err
		}
	}

	if params.Op == "add" {
		// 新增 一条新数据
		if item == nil {
			_, err = models.AddShoppingcar(&models.Shoppingcar{
				Uid:     params.Uid,
				Content: models.NewShoppingCarContent(good).ToString(),
				Total:   params.Num,
				Price:   good.Price,
			})

			if err != nil {
				return false, err
			}

			return true, nil
		}

		// 再原有数据上修改
		if err := item.AddGood(good, params.Num); err != nil {
			return false, err
		}
	}

	if err := models.UpdateShoppingcarById(item); err != nil {
		return false, err
	}

	return true, nil
}

