package models

import (
	"time"

	"github.com/astaxie/beego/orm"
)

//ProductImage 产品图片
type ProductImage struct {
	ID              int64            `orm:"column(id);pk;auto" json:"id"`         //主键
	CreateUser      *User            `orm:"rel(fk);null" json:"-"`                //创建者
	UpdateUser      *User            `orm:"rel(fk);null" json:"-"`                //最后更新者
	CreateDate      time.Time        `orm:"auto_now_add;type(datetime)" json:"-"` //创建时间
	UpdateDate      time.Time        `orm:"auto_now;type(datetime)" json:"-"`     //最后更新时间
	FormAction      string           `orm:"-" form:"FormAction"`                  //非数据库字段，用于表示记录的增加，修改
	Name            string           `orm:"unique" form:"name"`                   //图片名称
	ProductTemplate *ProductTemplate `orm:"rel(fk);null"`                         //款式图片
}

func init() {
	orm.RegisterModel(new(ProductImage))
}
