package models

import "github.com/astaxie/beego/orm"

// Mutant input mutant struct
type Mutant struct {
	Dna []string `json:"dna"`
}

// MutantDb database mutant struct
type MutantDb struct {
	ID       int      `orm:"column(id);pk;auto"`
	Dna      []string `orm:"column(dna)"`
	IsMutant bool     `orm:"column(is_mutant)"`
}

// TableName database entity
func (t *MutantDb) TableName() string {
	return "mutant"
}

func init() {
	orm.RegisterModel(new(MutantDb))
}

// AddMutant insert a new Mutant into database and returns
// last inserted Id on success.
func AddMutant(m *MutantDb) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}
