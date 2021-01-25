package helpers

import (
	"os"
	"testing"

	"github.com/Miavega/api_mutants_mercado_libre/helpers"
	"github.com/astaxie/beego/orm"

	_ "github.com/lib/pq"
)

func init() {
	err := orm.RegisterDriver("postgres", orm.DRPostgres)
	if err != nil {
		panic(err)
	}
	orm.Debug = true
}

func testWithDb(t *testing.T, f func(t *testing.T)) {
	pgUser := os.Getenv("POSTGRES_USER")
	pgPass := os.Getenv("POSTGRES_PASSWORD")
	pgUrls := os.Getenv("MUTANTS_CRUD_PGDB")
	pgDb := os.Getenv("POSTGRES_DB")
	pgSchema := os.Getenv("POSTGRES_SCHEMA")
	orm.RegisterDataBase("default", "postgres", "postgres://"+pgUser+":"+pgPass+"@"+pgUrls+"/"+pgDb+"?sslmode=disable&search_path="+pgSchema+"")

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
