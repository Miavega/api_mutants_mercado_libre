package controllers

import (
	"testing"
	//. "github.com/smartystreets/goconvey/convey"
)

func TestPost(t *testing.T) {
	// r, _ := http.NewRequest("POST", "/v1/mutant/", nil)
	// w := httptest.NewRecorder()
	// beego.BeeApp.Handlers.ServeHTTP(w, r)

	// fmt.Println("GOLHAS")

	// var response = map[string]interface{}{"Success": true, "Status": "404", "Message": "DNA evaluated is from a Mutant"}
	// json.Unmarshal(w.Body.Bytes(), &response)

	// Convey("Subject: Test Station Endpoint\n", t, func() {
	// 	Convey("Status Code Should Be 200", func() {
	// 		So(w.Code, ShouldEqual, 200)
	// 	})
	// 	Convey("The Result Should Not Be Empty", func() {
	// 		So(w.Body.Len(), ShouldBeGreaterThan, 0)
	// 	})
	// 	Convey("There Should Be A Result For Station 42002", func() {
	// 		So(response["Success"].(bool), ShouldEqual, true)
	// 	})
	// })
}
