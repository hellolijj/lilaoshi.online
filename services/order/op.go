package order

import (
	"cookie-shop-api/models"
	"cookie-shop-api/models/dto"
	"time"
)

func Add(params dto.OrderCreateParams) error {
	for _, item := range params.Data {
		good, err := models.GetGoodById(item.Id)
		if err != nil || good == nil {
			continue
		}

		_, err = models.AddOrders(&models.Orders{
			Uid:         params.Uid,
			GoodId:      item.Id,
			Price:       item.Price,
			Count:       item.Count,
			Status:      models.OrderCreated,
		})

		if err != nil {
			return err
		}
	}
	return nil
}

func Pay(params dto.OrderCreateParams) error {
	for _, item := range params.Data {
		good, err := models.GetGoodById(item.Id)
		if err != nil || good == nil {
			continue
		}

		err = models.UpdateOrdersById(&models.Orders{
			Id: item.Id,
			Uid:         params.Uid,
			GmtModified: time.Now().Local(),
			GoodId:      item.Id,
			Price:       item.Price,
			Count:       item.Count,
			Status:      models.OrderPayed,
		})

		if err != nil {
			return err
		}
	}
	return nil
}



// interface格式转为order
func ConvertOrder(m []interface{}) []models.OrderItem {
	var res []models.OrderItem

	for _, i := range m {
		o, ok := i.(models.Orders)
		if ok {
			good, _ := models.GetGoodById(o.GoodId)
			if good == nil {
				continue
			}

			res = append(res, models.OrderItem{
				Orders: o,
				Name: good.Name,
				Cover: good.Cover,
				Intro: good.Intro,
			})
		}
	}
	return res
}