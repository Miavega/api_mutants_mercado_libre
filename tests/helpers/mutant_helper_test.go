package helpers

import (
	"testing"

	"github.com/Miavega/api_mutants_mercado_libre/helpers"
	"github.com/astaxie/beego/logs"
	"github.com/astaxie/beego/orm"

	//_ "github.com/go-sql-driver/mysql"
	_ "github.com/lib/pq"
)

// var o orm.Ormer
// var connection = "postgres://postgres:mutants_db@us-east-2.console.aws.amazon.com/mutants_db?sslmode=disable&search_path=schema_xmen"

// func init() {
// 	err := orm.RegisterDriver("postgres", orm.DRPostgres)
// 	if err != nil {
// 		logs.Error(err)
// 		panic(err)
// 	}
// 	orm.Debug = true
// }

// func testWithDb(t *testing.T, f func(t *testing.T, mock sqlmock.Sqlmock)) {
// 	db, mock, err := sqlmock.New()
// 	if err != nil {
// 		t.Errorf("failed to open stub database connection, error: %v", err)
// 	}

// 	orm.AddAliasWthDB("default", "postgres", db)
// 	defer db.Close()

// 	f(t, mock)
// }

// func TestDBWithMockedSqlDriver(t *testing.T) {
// 	testWithDb(t, func(t *testing.T, mock sqlmock.Sqlmock) {
// 		// setup mock
// 		//columns := []string{"id", "dna", "is_mutant"}
// 		rows := sqlmock.NewRows([]string{"id", "dna", "is_mutant"}).
// 			AddRow(1, sqlmock.AnyArg(), sqlmock.AnyArg())
// 		mock.ExpectQuery("SELECT (.+) FROM mutant WHERE (.+)").WillReturnRows(rows)
// 		// call function to test
// 		request := []string{"ATGCGA", "CAGTGC", "TTATGT", "AGAAGG", "CCCCTA", "TCACTG"}
// 		t.Log("TestValidateHuman Finalizado Correctamente (OK CARAJO)")
// 		if res, err := helpers.Validate(request); err == nil {
// 			if !res {
// 				t.Log("TestValidateHuman Finalizado Correctamente (OK)")
// 			} else {
// 				t.Error("Se espera true y se obtuvo", res)
// 				t.Fail()
// 			}
// 		}
// 		if err := mock.ExpectationsWereMet(); err != nil {
// 			t.Errorf("there were unfulfilled expectations: %s", err)
// 		} else {

// 		}

// 		// we make sure that all expectations were met
// 		/*if err := mock.ExpectationsWereMet(); err != nil {
// 			t.Errorf("there were unfulfilled expectations: %s", err)
// 		}*/
// 	})
// }

func init() {
	err := orm.RegisterDriver("postgres", orm.DRPostgres)
	if err != nil {
		logs.Error(err)
		panic(err)
	}
	orm.Debug = true
}

func testWithDb(t *testing.T, f func(t *testing.T)) {
	// db, mock, err := sqlmock.New()
	// if err != nil {
	// 	t.Errorf("failed to open stub database connection, error: %v", err)
	// }

	orm.RegisterDataBase("default", "postgres", "postgres://postgres:mutants_db@ec2-3-139-102-224.us-east-2.compute.amazonaws.com/mutants_db?sslmode=disable&search_path=schema_xmen")

	//orm.AddAliasWthDB("default", "postgres", db)
	// defer db.Close()

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
