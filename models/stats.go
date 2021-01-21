package models

import "github.com/astaxie/beego/orm"

// Stats structure
type Stats struct {
	Id int `orm:"column(id);pk;auto"`
}

// TableName database entity
func (t *Stats) TableName() string {
	return "stats"
}

func init() {
	orm.RegisterModel(new(Stats))
}

// GetStats retrieves all Stats matches certain condition. Returns empty list if
// no records exist
func GetStats() (stats []map[string]interface{}, err error) {
	return nil, nil
}
