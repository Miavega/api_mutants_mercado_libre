package models

import (
	"github.com/astaxie/beego/orm"
)

// MutantStructure input mutant struct
type MutantStructure struct {
	Dna []string `json:"dna" valid:"Required;Match(/[A-CGT]/)"`
}

// Mutant database mutant struct
type Mutant struct {
	Id       int    `orm:"column(id);pk;auto"`
	Dna      string `orm:"column(dna)"`
	IsMutant bool   `orm:"column(is_mutant)"`
}

// TableName database entity
func (t *Mutant) TableName() string {
	return "mutant"
}

func init() {
	orm.RegisterModel(new(Mutant))
}

// AddMutant insert a new Mutant into database and returns
// last inserted Id on success.
func AddMutant(m *Mutant) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}

// GetMutantByDna retrieves Stats by DNA. Returns error if
// Id doesn't exist
func GetMutantByDna(dna string) (m *Mutant, err error) {
	o := orm.NewOrm()
	m = &Mutant{}
	if err = o.QueryTable("mutant").Filter("dna", dna).One(m); err == nil {
		return m, nil
	}
	return nil, err
}
