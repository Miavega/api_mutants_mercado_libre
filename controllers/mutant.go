package controllers

import (
	"encoding/json"

	"github.com/Miavega/api_mutants/helpers"

	"github.com/Miavega/api_mutants/models"
	"github.com/astaxie/beego"
)

// MutantController operations for Mutant
type MutantController struct {
	beego.Controller
}

// URLMapping ...
func (c *MutantController) URLMapping() {
	c.Mapping("Post", c.Post)
}

// Post ...
// @Title Create
// @Description create Mutant
// @Param	body		body 	models.Mutant	true		"body for Mutant content"
// @Success 201 {object} models.Mutant
// @Failure 403 body is empty
// @router / [post]
func (c *MutantController) Post() {
	var mutant models.MutantStructure
	json.Unmarshal(c.Ctx.Input.RequestBody, &mutant)
	if res, err := helpers.Validate(mutant.Dna); err == nil {
		if !res {
			c.Data["json"] = map[string]interface{}{"Success": res, "Status": "403", "Message": "DNA not Mutant"}
			c.Abort("403")
		}
		c.Data["json"] = map[string]interface{}{"Success": res, "Status": "200", "Message": "DNA Mutant"}
	} else {
		c.Data["json"] = map[string]interface{}{"Error": res, "Status": "500", "Message": "Request failure"}
		c.Abort("500")
	}
	c.ServeJSON()
}
