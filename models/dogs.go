package models

import (
	"errors"
	"fmt"

	"github.com/astaxie/beego/orm"
)

type Dogs struct {
	Id        int
	Name      string
	Race      string
	Age       string
	CreatedAt string
	CreatedBy string
	UpdateAt  string
	UpdateBy  string
}

func init() {
	orm.RegisterModel(new(Dogs))
}

func IsDogExist(dogID int) bool {
	o := orm.NewOrm()
	exist := o.QueryTable("dogs").Filter("id", dogID).Exist()
	return exist
}

func ReadDogs(limit, offset int) ([]Dogs, error) {
	var v []Dogs
	qb, _ := orm.NewQueryBuilder("mysql")
	qb.Select("id", "name", "race", "age", "created_at", "created_by", "update_at", "update_by").From("dogs").Limit(limit).Offset(offset)
	sql := qb.String()
	o := orm.NewOrm()
	fmt.Println(sql)
	if _, err := o.Raw(sql).QueryRows(&v); err != nil {
		fmt.Println(err)
		return nil, err
	}
	if len(v) == 0 {
		return nil, errors.New("empty data")
	}
	return v, nil
}

func ReadDogById(dogID int) (Dogs, error) {
	o := orm.NewOrm()
	var v Dogs
	if err := o.QueryTable("dogs").Filter("id", dogID).One(&v); err != nil {
		fmt.Println()
		return Dogs{}, err
	}
	return v, nil
}

func CreateDog(data Dogs) error {
	o := orm.NewOrm()
	_, err := o.Insert(&data)
	return err
}

func UpdateDog(dog orm.Params) error {
	o := orm.NewOrm()
	if _, err := o.QueryTable("dogs").Filter("id", dog["id"]).Update(dog); err != nil {
		return err
	}
	return nil
}

func DeleteDog(dogID int) error {
	o := orm.NewOrm()
	if _, err := o.QueryTable("dogs").Filter("id", dogID).Delete(); err != nil {
		return err
	}
	return nil
}
