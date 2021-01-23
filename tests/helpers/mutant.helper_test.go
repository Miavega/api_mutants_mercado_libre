package helpers

import (
	"database/sql/driver"
	"testing"
	"time"

	"github.com/Miavega/api_mutants/helpers"
	"github.com/astaxie/beego/logs"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
	"gopkg.in/DATA-DOG/go-sqlmock.v1"
)

var o orm.Ormer

func init() {
	err := orm.RegisterDriver("postgres", orm.DRPostgres)
	if err != nil {
		logs.Error(err)
		panic(err)
	}
	orm.Debug = true
}

func testWithDb(t *testing.T, f func(t *testing.T, mock sqlmock.Sqlmock)) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Errorf("failed to open stub database connection, error: %v", err)
	}

	orm.AddAliasWthDB("default", "postgres", db)
	orm.SetDataBaseTZ("default", time.Now().Location())
	defer db.Close()

	f(t, mock)
}

func TestDBWithMockedSqlDriver(t *testing.T) {
	testWithDb(t, func(t *testing.T, mock sqlmock.Sqlmock) {
		// setup mock
		//columns := []string{"id", "dna", "is_mutant"}
		/*rows := sqlmock.NewRows([]string{"id"}).
		AddRow(1)*/
		mock.ExpectExec("INSERT INTO `mutant` (`dna`, `is_mutant`) VALUES ($1, $2)").WillReturnResult(driver.RowsAffected(0))
		if err := mock.ExpectationsWereMet(); err != nil {
			t.Errorf("there were unfulfilled expectations: %s", err)
		} else {

			// call function to test
			request := []string{"ATGCGA", "CAGTGC", "TTATGT", "AGAAGG", "CCCCTA", "TCACTG"}
			t.Log("TestValidateHuman Finalizado Correctamente (OK CARAJO)")
			if res, err := helpers.Validate(request); err == nil {
				if !res {
					t.Log("TestValidateHuman Finalizado Correctamente (OK)")
				} else {
					t.Error("Se espera true y se obtuvo", res)
					t.Fail()
				}
			}
		}

		// we make sure that all expectations were met
		/*if err := mock.ExpectationsWereMet(); err != nil {
			t.Errorf("there were unfulfilled expectations: %s", err)
		}*/
	})
}

/*
func Test_ValidateHuman(t *testing.T) {
	request := []string{"ATGCGA", "CAGTGC", "TTATGT", "AGAAGG", "CCCCTA", "TCACTG"}
	if res, err := helpers.Validate(request); err == nil {
		if !res {
			t.Log("TestValidateHuman Finalizado Correctamente (OK)")
		} else {
			t.Error("Se espera true y se obtuvo", res)
			t.Fail()
		}
	}
}*/
