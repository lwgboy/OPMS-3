package models

import (
	"errors"
	"fmt"
	"goERP/utils"
	"strings"
	"time"

	"github.com/astaxie/beego/orm"
)

//ProductPricelistItem  产品价格
type ProductPricelistItem struct {
	ID         int64     `orm:"column(id);pk;auto" json:"id"`         //主键
	CreateUser *User     `orm:"rel(fk);null" json:"-"`                //创建者
	UpdateUser *User     `orm:"rel(fk);null" json:"-"`                //最后更新者
	CreateDate time.Time `orm:"auto_now_add;type(datetime)" json:"-"` //创建时间
	UpdateDate time.Time `orm:"auto_now;type(datetime)" json:"-"`     //最后更新时间
	FormAction string    `orm:"-" form:"FormAction"`                  //非数据库字段，用于表示记录的增加，修改
}

func init() {
	orm.RegisterModel(new(ProductPricelistItem))
}

// AddProductPricelistItem insert a new ProductPricelistItem into database and returns
// last inserted ID on success.
func AddProductPricelistItem(obj *ProductPricelistItem) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(obj)
	return id, err
}

// GetProductPricelistItemByID retrieves ProductPricelistItem by ID. Returns error if
// ID doesn't exist
func GetProductPricelistItemByID(id int64) (obj *ProductPricelistItem, err error) {
	o := orm.NewOrm()
	obj = &ProductPricelistItem{ID: id}
	if err = o.Read(obj); err == nil {
		return obj, nil
	}
	return nil, err
}

// GetAllProductPricelistItem retrieves all ProductPricelistItem matches certain condition. Returns empty list if
// no records exist
func GetAllProductPricelistItem(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64) (utils.Paginator, []ProductPricelistItem, error) {
	var (
		objArrs   []ProductPricelistItem
		paginator utils.Paginator
		num       int64
		err       error
	)
	if limit == 0 {
		limit = 20
	}
	o := orm.NewOrm()
	qs := o.QueryTable(new(ProductPricelistItem))
	qs = qs.RelatedSel()
	// query k=v
	for k, v := range query {
		// rewrite dot-notation to Object__Attribute
		k = strings.Replace(k, ".", "__", -1)
		qs = qs.Filter(k, v)
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
					return paginator, nil, errors.New("Error: Invalid order. Must be either [asc|desc]")
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
					return paginator, nil, errors.New("Error: Invalid order. Must be either [asc|desc]")
				}
				sortFields = append(sortFields, orderby)
			}
		} else if len(sortby) != len(order) && len(order) != 1 {
			return paginator, nil, errors.New("Error: 'sortby', 'order' sizes mismatch or 'order' size is not 1")
		}
	} else {
		if len(order) != 0 {
			return paginator, nil, errors.New("Error: unused 'order' fields")
		}
	}

	qs = qs.OrderBy(sortFields...)
	if cnt, err := qs.Count(); err == nil {
		paginator = utils.GenPaginator(limit, offset, cnt)
	}
	if num, err = qs.Limit(limit, offset).All(&objArrs, fields...); err == nil {
		paginator.CurrentPageSize = num
	}
	return paginator, objArrs, err
}

// UpdateProductPricelistItemByID updates ProductPricelistItem by ID and returns error if
// the record to be updated doesn't exist
func UpdateProductPricelistItemByID(m *ProductPricelistItem) (err error) {
	o := orm.NewOrm()
	v := ProductPricelistItem{ID: m.ID}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}

// DeleteProductPricelistItem deletes ProductPricelistItem by ID and returns error if
// the record to be deleted doesn't exist
func DeleteProductPricelistItem(id int64) (err error) {
	o := orm.NewOrm()
	v := ProductPricelistItem{ID: id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&ProductPricelistItem{ID: id}); err == nil {
			fmt.Println("Number of records deleted in database:", num)
		}
	}
	return
}
