package models

import (
	"fmt"

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
	fmt.Println("LLEGUE?EEEEEEEEEEE")
	id, err = o.Insert(m)
	fmt.Println("LLEGUE?")
	return
}
