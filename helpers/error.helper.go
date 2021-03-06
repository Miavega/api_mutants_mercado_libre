package helpers

import (
	"github.com/astaxie/beego"
)

// ErrorController operations for Usuario
type ErrorController struct {
	beego.Controller
}

func (c *ErrorController) Error403() {
	outputError := map[string]interface{}{"Success": false, "Status": "403", "Message": c.Data["mesaage"], "Data": c.Data["data"]}
	c.Data["json"] = outputError
	c.ServeJSON()
}

func (c *ErrorController) Error500() {
	outputError := map[string]interface{}{"Success": false, "Status": "500", "Message": c.Data["mesaage"], "Data": c.Data["data"]}
	c.Data["json"] = outputError
	c.ServeJSON()
}

func (c *ErrorController) Error422() {
	outputError := map[string]interface{}{"Success": false, "Status": "422", "Message": c.Data["mesaage"], "Data": c.Data["data"]}
	c.Data["json"] = outputError
	c.ServeJSON()
}
