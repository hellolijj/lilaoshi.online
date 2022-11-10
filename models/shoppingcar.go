package models

import (
	"encoding/json"
	"errors"
	"fmt"
	"reflect"
	"sort"
	"strconv"
	"strings"

	"github.com/beego/beego/v2/client/orm"
)

type Shoppingcar struct {
	Id      int    `orm:"column(id);auto" description:"购物车id" json:"id"`
	Uid     string `orm:"column(uid);size(45)" description:"用户id" json:"uid"`
	Content string `orm:"column(content);size(64);null" description:"购物车内容json结构" json:"content"`
	Total   int `orm:"column(total);size(8);null" description:"购物车总数量" json:"total"`
	Price   string `orm:"column(price);size(8);null" description:"购物车总价格" json:"price"`
}

func (t *Shoppingcar) TableName() string {
	return "shoppingcar"
}

func init() {
	orm.RegisterModel(new(Shoppingcar))
}

// AddShoppingcar insert a new Shoppingcar into database and returns
// last inserted Id on success.
func AddShoppingcar(m *Shoppingcar) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}

// GetShoppingcarById retrieves Shoppingcar by Id. Returns error if
// Id doesn't exist
func GetShoppingcarById(id int) (v *Shoppingcar, err error) {
	o := orm.NewOrm()
	v = &Shoppingcar{Id: id}
	if err = o.Read(v); err == nil {
		return v, nil
	}
	return nil, err
}

// GetAllShoppingcar retrieves all Shoppingcar matches certain condition. Returns empty list if
// no records exist
func GetAllShoppingcar(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64) (ml []interface{}, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(Shoppingcar))
	// query k=v
	for k, v := range query {
		// rewrite dot-notation to Object__Attribute
		k = strings.Replace(k, ".", "__", -1)
		if strings.Contains(k, "isnull") {
			qs = qs.Filter(k, (v == "true" || v == "1"))
		} else {
			qs = qs.Filter(k, v)
		}
	}
	// order by:
	var sortFields []string
	if len(sortby) != 0 {
		if len(sortby) == len(order) {
			// 1) for each sort field, there is an associated order
			for i, v := range sortby {
				orderby := ""
				if order[i] == "desc" {
					orderby = "-" + v
				} else if order[i] == "asc" {
					orderby = v
				} else {
					return nil, errors.New("Error: Invalid order. Must be either [asc|desc]")
				}
				sortFields = append(sortFields, orderby)
			}
			qs = qs.OrderBy(sortFields...)
		} else if len(sortby) != len(order) && len(order) == 1 {
			// 2) there is exactly one order, all the sorted fields will be sorted by this order
			for _, v := range sortby {
				orderby := ""
				if order[0] == "desc" {
					orderby = "-" + v
				} else if order[0] == "asc" {
					orderby = v
				} else {
					return nil, errors.New("Error: Invalid order. Must be either [asc|desc]")
				}
				sortFields = append(sortFields, orderby)
			}
		} else if len(sortby) != len(order) && len(order) != 1 {
			return nil, errors.New("Error: 'sortby', 'order' sizes mismatch or 'order' size is not 1")
		}
	} else {
		if len(order) != 0 {
			return nil, errors.New("Error: unused 'order' fields")
		}
	}

	var l []Shoppingcar
	qs = qs.OrderBy(sortFields...)
	if _, err = qs.Limit(limit, offset).All(&l, fields...); err == nil {
		if len(fields) == 0 {
			for _, v := range l {
				ml = append(ml, v)
			}
		} else {
			// trim unused fields
			for _, v := range l {
				m := make(map[string]interface{})
				val := reflect.ValueOf(v)
				for _, fname := range fields {
					m[fname] = val.FieldByName(fname).Interface()
				}
				ml = append(ml, m)
			}
		}
		return ml, nil
	}
	return nil, err
}

// UpdateShoppingcar updates Shoppingcar by Id and returns error if
// the record to be updated doesn't exist
func UpdateShoppingcarById(m *Shoppingcar) (err error) {
	o := orm.NewOrm()
	v := Shoppingcar{Id: m.Id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m, "content", "total", "price"); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}

// DeleteShoppingcar deletes Shoppingcar by Id and returns error if
// the record to be deleted doesn't exist
func DeleteShoppingcar(id int) (err error) {
	o := orm.NewOrm()
	v := Shoppingcar{Id: id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&Shoppingcar{Id: id}); err == nil {
			fmt.Println("Number of records deleted in database:", num)
		}
	}
	return
}

func GetShoppingcarByUId(uid string) (v *Shoppingcar, err error) {
	o := orm.NewOrm()
	v = &Shoppingcar{Uid: uid}
	if err = o.Read(v, "uid"); err == nil {
		return v, nil
	}
	return nil, err
}


func (t *Shoppingcar) ToGoods() (*ShoppingCarGoods, error) {
	if t == nil {
		return nil, nil
	}

	var res ShoppingCarGoods
	res.Shoppingcar = *t

	data, err := t.toShoppingCarContent()
	if err != nil {
		return nil, err
	}

	for goodId, goodCount := range data.Content {
		var shoppingCarGood ShoppingCarGood
		if good, err := GetGoodById(goodId); err == nil {
			shoppingCarGood.Good = *good
			shoppingCarGood.Count = goodCount
		}
		res.Data = append(res.Data, shoppingCarGood)
	}

	// 按照id排序
	sort.Slice(res.Data, func(i, j int) bool {
		return res.Data[i].Id < res.Data[j].Id
	})

	return &res, nil
}


func (t *Shoppingcar) toShoppingCarContent() (*ShoppingCarContent, error)   {
	var data ShoppingCarContent
	err := json.Unmarshal([]byte(t.Content), &data)
	if err != nil {
		return nil, err
	}
	return &data, nil
}

func (t *Shoppingcar) AddGood(good *Good, count int) error {
	price, err := strconv.ParseFloat(good.Price, 32)
	if err != nil {
		return err
	}

	var data ShoppingCarContent
	err = json.Unmarshal([]byte(t.Content), &data)
	if err != nil {
		return err
	}

	data.Count = data.Count + count
	data.Content[good.Id] = data.Content[good.Id] + count
	data.Price = data.Price + float32(price)*float32(count)

	t.Price = fmt.Sprintf("%.2f", data.Price)
	t.Total = data.Count
	t.Content = data.ToString()
	return nil
}

func (t *Shoppingcar) SubGood(good *Good, count int) error {
	price, err := strconv.ParseFloat(good.Price, 32)
	if err != nil {
		return err
	}

	var data ShoppingCarContent
	err = json.Unmarshal([]byte(t.Content), &data)
	if err != nil {
		return err
	}

	data.Count = data.Count - count
	data.Content[good.Id] = data.Content[good.Id] - count
	data.Price = data.Price - float32(price)*float32(count)

	if data.Count < 0 || data.Content[good.Id] < 0 {
		return errors.New("count can not lt 0")
	}

	t.Price = fmt.Sprintf("%.2f", data.Price)
	t.Total = data.Count
	t.Content = data.ToString()
	return nil
}

type ShoppingCarContent struct {
	Count int `json:"count"`
	Price float32 `json:"price"`
	Content map[int]int `json:"content"`
}

func NewShoppingCarContent(good *Good) *ShoppingCarContent {
	price, _ := strconv.ParseFloat(good.Price, 32)
	return &ShoppingCarContent{
		Count:   1,
		Price:   float32(price),
		Content: map[int]int{good.Id: 1},
	}
}

func (car *ShoppingCarContent) Add(good *Good) {
	car.Count ++
	price, _ := strconv.ParseFloat(good.Price, 32)
	car.Price += float32(price)
	car.Content[good.Id] ++
}

func (car *ShoppingCarContent) ToString() string {
	str, _ := json.Marshal(car)
	return string(str)
}

type ShoppingCarGoods struct {
	Shoppingcar
	Data []ShoppingCarGood `json:"data"`
}

type ShoppingCarGood struct {
	Good
	Count int `json:"count"`
}

