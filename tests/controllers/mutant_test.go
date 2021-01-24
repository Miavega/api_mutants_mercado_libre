package controllers

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	_ "github.com/lib/pq"
	. "github.com/smartystreets/goconvey/convey"
)

func TestPost(t *testing.T) {
	//var jsonStr = []byte(`{"title":"Buy cheese and bread for breakfast."}`)
	request := []byte(`{"dna": ["ATGCGA","CAGTGC","TTATGT","AGAAGG","CCCCTA","TCACTG"]}`)
	logs.Info(request)
	r, _ := http.NewRequest("POST", "/v1/mutant/", bytes.NewBuffer(request))
	r.Header.Set("Content-Type", "application/json")
	r.Header.Set("Content-Type", "text/plain")
	w := httptest.NewRecorder()
	beego.BeeApp.Handlers.ServeHTTP(w, r)

	var response = map[string]interface{}{"Success": true, "Status": "200", "Message": "DNA evaluated is from a Mutant"}
	json.Unmarshal(w.Body.Bytes(), &response)

	Convey("Subject: Test mutant Endpoint\n", t, func() {
		Convey("Status Code Should Be 200", func() {
			So(w.Code, ShouldEqual, 200)
		})
		Convey("The Result Should Not Be Empty", func() {
			So(w.Body.Len(), ShouldBeGreaterThan, 0)
		})
		Convey("There Should Be A Result For Station 42002", func() {
			So(response["Success"].(bool), ShouldEqual, true)
		})
	})
}
