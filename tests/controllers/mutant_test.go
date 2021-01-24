package controllers

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/astaxie/beego/logs"

	"github.com/astaxie/beego"
	_ "github.com/lib/pq"
	. "github.com/smartystreets/goconvey/convey"
)

func TestPost(t *testing.T) {
	request := []byte(`{"dna": ["ATGCGA","CAGTGC","TTATGT","AGAAGG","CCCCTA","TCACTG"]}`)
	logs.Info(request)
	r, _ := http.NewRequest("POST", "/v1/mutant/", bytes.NewBuffer(request))
	r.Header.Set("Content-Type", "application/json")
	r.Header.Set("Content-Type", "text/plain")
	w := httptest.NewRecorder()
	beego.BeeApp.Handlers.ServeHTTP(w, r)

	var response = map[string]interface{}{}
	json.Unmarshal(w.Body.Bytes(), &response)

	Convey("Subject: Test mutant Endpoint\n", t, func() {
		Convey("Status Code Should Be 422", func() {
			So(w.Code, ShouldEqual, 422)
		})
		Convey("The Result Should Not Be Empty", func() {
			So(w.Body.Len(), ShouldBeGreaterThan, 0)
		})
	})
}
