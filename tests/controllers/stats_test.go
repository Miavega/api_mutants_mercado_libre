package controllers

import (
	"testing"

	_ "github.com/Miavega/api_mutants_mercado_libre/routers"
	//. "github.com/smartystreets/goconvey/convey"
)

func TestGet(t *testing.T) {
	// 	r, _ := http.NewRequest("GET", "/v1/stats", nil)
	// 	w := httptest.NewRecorder()
	// 	beego.BeeApp.Handlers.ServeHTTP(w, r)

	// 	fmt.Println("GOLHAS")

	// 	var response = map[string]interface{}{"Success": true, "Status": "200", "Message": "DNA evaluated is from a Mutant"}
	// 	json.Unmarshal(w.Body.Bytes(), &response)

	// 	Convey("Subject: Test Station Endpoint\n", t, func() {
	// 		Convey("Status Code Should Be 200", func() {
	// 			So(w.Code, ShouldEqual, 200)
	// 		})
	// 		Convey("The Result Should Not Be Empty", func() {
	// 			So(w.Body.Len(), ShouldBeGreaterThan, 0)
	// 		})
	// 		Convey("There Should Be A Result For Station 42002", func() {
	// 			So(response["Success"].(bool), ShouldEqual, true)
	// 		})
	// 	})
}