package controllers

import (
	"cookie-shop-api/models"
	"cookie-shop-api/models/dto"
	"cookie-shop-api/tools/crawler"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/beego/beego/v2/server/web/context"
	"strconv"
	"strings"

	beego "github.com/beego/beego/v2/server/web"
)

// GoodController operations for Good
type GoodController struct {
	beego.Controller
}

// URLMapping ...
func (c *GoodController) URLMapping() {
	c.Mapping("Post", c.Post)
	c.Mapping("GetOne", c.GetOne)
	c.Mapping("GetAll", c.GetAll)
	c.Mapping("Put", c.Put)
	c.Mapping("Delete", c.Delete)
}

// Post ...
// @Title Post
// @Description create Good
// @Param	body		body 	models.Good	true		"body for Good content"
// @Success 201 {int} models.Good
// @Failure 403 body is empty
// @router / [post]
func (c *GoodController) Post() {
	var v models.Good
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &v); err == nil {
		if _, err := models.AddGood(&v); err == nil {
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
// @Description get Good by id
// @Param	id		path 	string	true		"The key for staticblock"
// @Success 200 {object} models.Good
// @Failure 403 :id is empty
// @router /:id [get]
func (c *GoodController) GetOne() {
	idStr := c.Ctx.Input.Param(":id")
	id, _ := strconv.Atoi(idStr)
	v, err := models.GetGoodById(id)
	if err != nil {
		c.Data["json"] = err.Error()
	} else {
		c.Data["json"] = v
	}
	c.ServeJSON()
}

// GetAll ...
// @Title Get All
// @Description get Good
// @Param	query	query	string	false	"Filter. e.g. col1:v1,col2:v2 ..."
// @Param	fields	query	string	false	"Fields returned. e.g. col1,col2 ..."
// @Param	sortby	query	string	false	"Sorted-by fields. e.g. col1,col2 ..."
// @Param	order	query	string	false	"Order corresponding to each sortby field, if single value, apply to all sortby fields. e.g. desc,asc ..."
// @Param	limit	query	string	false	"Limit the size of result set. Must be an integer"
// @Param	offset	query	string	false	"Start position of result set. Must be an integer"
// @Success 200 {object} models.Good
// @Failure 403
// @router / [get]
func (c *GoodController) GetAll() {
	var fields []string
	var sortby []string
	var order []string
	var query = make(map[string]string)
	var limit int64 = 10
	var offset, current int64

	// fields: col1,col2,entity.col3
	if v := c.GetString("fields"); v != "" {
		fields = strings.Split(v, ",")
	}
	// limit: 10 (default is 10)
	if v, err := c.GetInt64("limit"); err == nil {
		limit = v
	}
	// current, offset: 0 (default is 0)
	if v, err := c.GetInt64("current"); err == nil {
		offset = (v-1) * limit
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

	l, err := models.GetAllGood(query, fields, sortby, order, offset, limit)
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
// @Description update the Good
// @Param	id		path 	string	true		"The id you want to update"
// @Param	body		body 	models.Good	true		"body for Good content"
// @Success 200 {object} models.Good
// @Failure 403 :id is not int
// @router /:id [put]
func (c *GoodController) Put() {
	idStr := c.Ctx.Input.Param(":id")
	id, _ := strconv.Atoi(idStr)
	v := models.Good{Id: id}
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &v); err == nil {
		if err := models.UpdateGoodById(&v); err == nil {
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
// @Description delete the Good
// @Param	id		path 	string	true		"The id you want to delete"
// @Success 200 {string} delete success!
// @Failure 403 id is empty
// @router /:id [delete]
func (c *GoodController) Delete() {
	idStr := c.Ctx.Input.Param(":id")
	id, _ := strconv.Atoi(idStr)
	if err := models.DeleteGood(id); err == nil {
		c.Data["json"] = "OK"
	} else {
		c.Data["json"] = err.Error()
	}
	c.ServeJSON()
}


func (ctrl GoodController) Crawler(ctx *context.Context) {
	goods := crawler.FetchCookies()
	for _, good := range goods {
		fmt.Println(good)
		data, _ := models.GetGoodByName(good.Name)
		if data == nil {
			_, err := models.AddGood(&good)
			if err != nil {
				fmt.Println(err.Error())
			}
		}
	}
}
