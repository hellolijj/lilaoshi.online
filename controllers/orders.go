package controllers

import (
	"cookie-shop-api/models"
	"cookie-shop-api/models/dto"
	orderservice "cookie-shop-api/services/order"
	"cookie-shop-api/services/shoppingcar"
	"cookie-shop-api/services/user"
	"encoding/json"
	"errors"
	"github.com/beego/beego/v2/server/web/context"
	"strconv"
	"strings"

	beego "github.com/beego/beego/v2/server/web"
)

// OrdersController operations for Orders
type OrdersController struct {
	beego.Controller
}

// URLMapping ...
func (c *OrdersController) URLMapping() {
	c.Mapping("Post", c.Post)
	c.Mapping("GetOne", c.GetOne)
	c.Mapping("GetAll", c.GetAll)
	c.Mapping("Put", c.Put)
	c.Mapping("Delete", c.Delete)
}

// Post ...
// @Title Post
// @Description create Orders
// @Param	body		body 	models.Orders	true		"body for Orders content"
// @Success 201 {int} models.Orders
// @Failure 403 body is empty
// @router / [post]
func (c *OrdersController) Post() {
	var v models.Orders
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &v); err == nil {
		if _, err := models.AddOrders(&v); err == nil {
			c.Ctx.Output.SetStatus(201)
			c.Data["json"] = v
		} else {
			c.Data["json"] = err.Error()
		}
	} else {
		c.Data["json"] = err.Error()
	}
	c.ServeJSON()
}

// GetOne ...
// @Title Get One
// @Description get Orders by id
// @Param	id		path 	string	true		"The key for staticblock"
// @Success 200 {object} models.Orders
// @Failure 403 :id is empty
// @router /:id [get]
func (c *OrdersController) GetOne() {
	idStr := c.Ctx.Input.Param(":id")
	id, _ := strconv.Atoi(idStr)
	v, err := models.GetOrdersById(id)
	if err != nil {
		c.Data["json"] = err.Error()
	} else {
		c.Data["json"] = v
	}
	c.ServeJSON()
}

// GetAll ...
// @Title Get All
// @Description get Orders
// @Param	query	query	string	false	"Filter. e.g. col1:v1,col2:v2 ..."
// @Param	fields	query	string	false	"Fields returned. e.g. col1,col2 ..."
// @Param	sortby	query	string	false	"Sorted-by fields. e.g. col1,col2 ..."
// @Param	order	query	string	false	"Order corresponding to each sortby field, if single value, apply to all sortby fields. e.g. desc,asc ..."
// @Param	limit	query	string	false	"Limit the size of result set. Must be an integer"
// @Param	offset	query	string	false	"Start position of result set. Must be an integer"
// @Success 200 {object} models.Orders
// @Failure 403
// @router / [get]
func (c *OrdersController) GetAll() {
	var fields []string
	var sortby []string
	var order []string
	var query = make(map[string]string)
	var limit int64 = 10
	var offset int64
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

	l, err := models.GetAllOrders(query, fields, sortby, order, offset, limit)
	if err != nil {
		c.Data["json"] = err.Error()
	} else {
		c.Data["json"] = dto.ListResp{
			Current:  1,
			PageSize: 20,
			Total:    len(l),
			Success:  true,
			Data:     orderservice.ConvertOrder(l),
		}
	}
	c.ServeJSON()
}

// Put ...
// @Title Put
// @Description update the Orders
// @Param	id		path 	string	true		"The id you want to update"
// @Param	body		body 	models.Orders	true		"body for Orders content"
// @Success 200 {object} models.Orders
// @Failure 403 :id is not int
// @router /:id [put]
func (c *OrdersController) Put() {
	idStr := c.Ctx.Input.Param(":id")
	id, _ := strconv.Atoi(idStr)
	v := models.Orders{Id: id}
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &v); err == nil {
		if err := models.UpdateOrdersById(&v); err == nil {
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
// @Description delete the Orders
// @Param	id		path 	string	true		"The id you want to delete"
// @Success 200 {string} delete success!
// @Failure 403 id is empty
// @router /:id [delete]
func (c *OrdersController) Delete() {
	idStr := c.Ctx.Input.Param(":id")
	id, _ := strconv.Atoi(idStr)
	if err := models.DeleteOrders(id); err == nil {
		c.Data["json"] = "OK"
	} else {
		c.Data["json"] = err.Error()
	}
	c.ServeJSON()
}

// 创建订单
func (ctrl OrdersController) Create(ctx *context.Context) {
	var input dto.OrderCreateParams
	err := json.Unmarshal(ctx.Input.RequestBody, &input)
	if err != nil {
		ctx.WriteString(err.Error())
		return
	}
	_, input.Uid = user.Online(ctx)
	err = orderservice.Add(input)
	if err != nil {
		ctx.WriteString(err.Error())
		return
	}

	for _, d := range input.Data {
		shoppingcar.Op(dto.ShoppingCarOpParams{
			Op:     "sub",
			GoodId: d.Id,
			Uid:    input.Uid,
			Num:    d.Count,
		})
	}

	ctx.WriteString("success")
}

// 创建订单
func (ctrl OrdersController) Pay(ctx *context.Context) {
	var input dto.OrderCreateParams
	err := json.Unmarshal(ctx.Input.RequestBody, &input)
	if err != nil {
		ctx.WriteString(err.Error())
		return
	}
	_, input.Uid = user.Online(ctx)
	err = orderservice.Pay(input)
	if err != nil {
		ctx.WriteString(err.Error())
		return
	}

	ctx.WriteString("success")
}

