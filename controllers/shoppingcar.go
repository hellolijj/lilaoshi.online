package controllers

import (
	"cookie-shop-api/models"
	"cookie-shop-api/models/dto"
	"cookie-shop-api/services/shoppingcar"
	"cookie-shop-api/services/user"
	"encoding/json"
	"errors"
	"github.com/beego/beego/v2/server/web/context"
	"strconv"
	"strings"

	beego "github.com/beego/beego/v2/server/web"
)

// ShoppingcarController operations for Shoppingcar
type ShoppingcarController struct {
	beego.Controller
}

// URLMapping ...
func (c *ShoppingcarController) URLMapping() {
	c.Mapping("Post", c.Post)
	c.Mapping("GetOne", c.GetOne)
	c.Mapping("GetAll", c.GetAll)
	c.Mapping("Put", c.Put)
	c.Mapping("Delete", c.Delete)
}

// Post ...
// @Title Post
// @Description create Shoppingcar
// @Param	body		body 	models.Shoppingcar	true		"body for Shoppingcar content"
// @Success 201 {int} models.Shoppingcar
// @Failure 403 body is empty
// @router / [post]
func (c *ShoppingcarController) Post() {
	var v models.Shoppingcar
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &v); err == nil {
		//fmt.Println(v)
		//if _, err := models.AddShoppingcar(&v); err == nil {
		//	c.Ctx.Output.SetStatus(201)
		//	c.Data["json"] = v
		//} else {
		//	c.Data["json"] = err.Error()
		//}
	} else {
		c.Data["json"] = err.Error()
	}
	c.ServeJSON()
}

// GetOne ...
// @Title Get One
// @Description get Shoppingcar by id
// @Param	id		path 	string	true		"The key for staticblock"
// @Success 200 {object} models.Shoppingcar
// @Failure 403 :id is empty
// @router /:id [get]
func (c *ShoppingcarController) GetOne() {
	idStr := c.Ctx.Input.Param(":id")
	id, _ := strconv.Atoi(idStr)
	v, err := models.GetShoppingcarById(id)
	if err != nil {
		c.Data["json"] = err.Error()
	} else {
		c.Data["json"] = v
	}
	c.ServeJSON()
}

// GetAll ...
// @Title Get All
// @Description get Shoppingcar
// @Param	query	query	string	false	"Filter. e.g. col1:v1,col2:v2 ..."
// @Param	fields	query	string	false	"Fields returned. e.g. col1,col2 ..."
// @Param	sortby	query	string	false	"Sorted-by fields. e.g. col1,col2 ..."
// @Param	order	query	string	false	"Order corresponding to each sortby field, if single value, apply to all sortby fields. e.g. desc,asc ..."
// @Param	limit	query	string	false	"Limit the size of result set. Must be an integer"
// @Param	offset	query	string	false	"Start position of result set. Must be an integer"
// @Success 200 {object} models.Shoppingcar
// @Failure 403
// @router / [get]
func (c *ShoppingcarController) GetAll() {
	var fields []string
	var sortby []string
	var order []string
	var query = make(map[string]string)
	var limit int64 = 10
	var offset, current int64
	_, uid := user.Online(c.Ctx)
	query["uid"] = uid

	// fields: col1,col2,entity.col3
	if v := c.GetString("fields"); v != "" {
		fields = strings.Split(v, ",")
	}
	// limit: 10 (default is 10)
	if v, err := c.GetInt64("limit"); err == nil {
		limit = v
	}
	// offset: 0 (default is 0)
	if v, err := c.GetInt64("offset"); err == nil {
		offset = v
	}
	// sortby: col1,col2
	if v := c.GetString("sortby"); v != "" {
		sortby = strings.Split(v, ",")
	}
	// order: desc,asc
	if v := c.GetString("order"); v != "" {
		order = strings.Split(v, ",")
	}
	// query: k:v,k:v
	if v := c.GetString("query"); v != "" {
		for _, cond := range strings.Split(v, ",") {
			kv := strings.SplitN(cond, ":", 2)
			if len(kv) != 2 {
				c.Data["json"] = errors.New("Error: invalid query key/value pair")
				c.ServeJSON()
				return
			}
			k, v := kv[0], kv[1]
			query[k] = v
		}
	}

	l, err := models.GetAllShoppingcar(query, fields, sortby, order, offset, limit)
	if err != nil {
		c.Data["json"] = err.Error()
	} else {
		c.Data["json"] = dto.ListResp{
			Success: true,
			Data:    l,
			Total: len(l),
			Current: current,
			PageSize: limit,
		}
	}
	c.ServeJSON()
}

// Put ...
// @Title Put
// @Description update the Shoppingcar
// @Param	id		path 	string	true		"The id you want to update"
// @Param	body		body 	models.Shoppingcar	true		"body for Shoppingcar content"
// @Success 200 {object} models.Shoppingcar
// @Failure 403 :id is not int
// @router /:id [put]
func (c *ShoppingcarController) Put() {
	idStr := c.Ctx.Input.Param(":id")
	id, _ := strconv.Atoi(idStr)
	v := models.Shoppingcar{Id: id}
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &v); err == nil {
		if err := models.UpdateShoppingcarById(&v); err == nil {
			c.Data["json"] = "OK"
		} else {
			c.Data["json"] = err.Error()
		}
	} else {
		c.Data["json"] = err.Error()
	}
	c.ServeJSON()
}

// Delete ...
// @Title Delete
// @Description delete the Shoppingcar
// @Param	id		path 	string	true		"The id you want to delete"
// @Success 200 {string} delete success!
// @Failure 403 id is empty
// @router /:id [delete]
func (c *ShoppingcarController) Delete() {
	idStr := c.Ctx.Input.Param(":id")
	id, _ := strconv.Atoi(idStr)
	if err := models.DeleteShoppingcar(id); err == nil {
		c.Data["json"] = "OK"
	} else {
		c.Data["json"] = err.Error()
	}
	c.ServeJSON()
}


// 操作购物车
func (ctrl ShoppingcarController) Op(ctx *context.Context) {
	var input dto.ShoppingCarOpParams
	err := json.Unmarshal(ctx.Input.RequestBody, &input)
	if err != nil {
		ctx.WriteString(err.Error())
		return
	}
	_, input.Uid = user.Online(ctx)

	if success, err := shoppingcar.Op(input); !success || err != nil {
		ctx.WriteString(err.Error())
		return
	}

	data, _ := json.Marshal(input)
	ctx.ResponseWriter.Write(data)
}

//Goods
func (ctrl ShoppingcarController) Goods(ctx *context.Context) {
	_, uid := user.Online(ctx)
	car, err := models.GetShoppingcarByUId(uid)
	if err != nil {
		ctx.WriteString(err.Error())
		return
	}

	if car == nil {
		ctx.WriteString("null")
		return
	}

	goods, err := car.ToGoods()
	if err != nil {
		ctx.WriteString(err.Error())
		return
	}
	data, _ := json.Marshal(dto.ShoppingCarGoodsResp{
		Success: true,
		Data:    goods.Data,
		Total: car.Total,
		GoodsTotal:  goods.Total,
		Price: goods.Price,
	})
	ctx.ResponseWriter.Write(data)
}
