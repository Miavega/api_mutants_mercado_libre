package controllers

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	_ "github.com/Miavega/api_mutants_mercado_libre/routers"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"github.com/astaxie/beego/orm"
	. "github.com/smartystreets/goconvey/convey"
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

func TestGet(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			t.Error("Error", r)
			t.Fail()
		} else {
			t.Log("OK")
		}
	}()

	testWithDb(t, func(t *testing.T) {
		r, _ := http.NewRequest("GET", "/v1/stats", nil)
		w := httptest.NewRecorder()
		beego.BeeApp.Handlers.ServeHTTP(w, r)

		var response = map[string]interface{}{}
		json.Unmarshal(w.Body.Bytes(), &response)

		Convey("Subject: Test Stats Endpoint\n", t, func() {
			Convey("Status Code Should Be 200", func() {
				So(w.Code, ShouldEqual, 200)
			})
			Convey("The Result Should Not Be Empty", func() {
				So(w.Body.Len(), ShouldBeGreaterThan, 0)
			})
			Convey("There Should Be A Result For Stats", func() {
				So(response["count_human_dna"], ShouldBeGreaterThanOrEqualTo, 0)
			})
		})
	})

}
