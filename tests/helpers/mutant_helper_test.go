package helpers

import (
	"testing"

	"github.com/Miavega/api_mutants_mercado_libre/helpers"
	"github.com/astaxie/beego/logs"
	"github.com/astaxie/beego/orm"

	_ "github.com/lib/pq"
)

func init() {
	err := orm.RegisterDriver("postgres", orm.DRPostgres)
	if err != nil {
		logs.Error(err)
		panic(err)
	}
	orm.Debug = true
}

func testWithDb(t *testing.T, f func(t *testing.T)) {
	orm.RegisterDataBase("default", "postgres", "postgres://postgres:mutants_db@ec2-3-139-102-224.us-east-2.compute.amazonaws.com/mutants_db?sslmode=disable&search_path=schema_xmen")

	f(t)
}

func TestValidateHuman(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			t.Error("Error", r)
			t.Fail()
		} else {
			t.Log("OK")
		}
	}()

	testWithDb(t, func(t *testing.T) {
		request := []string{"ATGCGA", "CAGTGC", "TTATGT", "AGAAGG", "CCCCTA", "TCACTG"}
		if _, err := helpers.Validate(request); err != nil {
			panic(err.Error())
		}
	})
}

func TestValidateMutant(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			t.Error("Error", r)
			t.Fail()
		} else {
			t.Log("OK")
		}
	}()

	testWithDb(t, func(t *testing.T) {
		request := []string{"ATGCGA", "CAGTGC", "TTATTT", "AGACGG", "GCGCTA", "TCACTG"}
		if _, err := helpers.Validate(request); err != nil {
			panic(err.Error())
		}
	})
}
