package models

import (
	"fmt"

	"github.com/astaxie/beego/orm"
)

// Stats structure
type Stats struct {
	Id             int `orm:"column(id);pk;auto"`
	CountMutantDna int `orm:"column(count_mutant_dna)"`
	CountHumanDna  int `orm:"column(count_human_dna)"`
	Ratio          int `orm:"column(ratio)"`
}

// TableName database entity
func (t *Stats) TableName() string {
	return "stats"
}

func init() {
	orm.RegisterModel(new(Stats))
}

// AddStats insert a new Stats into database and returns
// last inserted Id on success.
func AddStats(m *Stats) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}

// GetStatsById retrieves Stats by Id. Returns error if
// Id doesn't exist
func GetStatsById(id int) (m *Stats, err error) {
	o := orm.NewOrm()
	m = &Stats{Id: id}
	if err = o.Read(m); err == nil {
		return m, nil
	}
	return nil, err
}

// UpdateStats ...
func UpdateStats(m *Stats) (Stats, error) {
	var err error
	o := orm.NewOrm()
	v := Stats{Id: m.Id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		if _, err = o.Update(&v); err == nil {
			fmt.Printf("Number of records updated in database: %v \n", v)
			return v, nil
		}
	}
	return v, err
	/*var result Stats
	o := orm.NewOrm()
	qs := o.QueryTable(new(Stats)).OrderBy("-id").One(&result)
	fmt.Println(qs)
	return nil, nil*/
}
