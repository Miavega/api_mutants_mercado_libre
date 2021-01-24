package controllers

import (
	"encoding/json"

	"github.com/Miavega/api_mutants_mercado_libre/helpers"

	"github.com/Miavega/api_mutants_mercado_libre/models"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/validation"
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
// @Success 200 {object} models.Mutant
// @Failure 403 body is empty
// @router / [post]
func (c *MutantController) Post() {
	var mutant models.MutantStructure
	valid := validation.Validation{}
	json.Unmarshal(c.Ctx.Input.RequestBody, &mutant)
	if mutantValidate, err := valid.Valid(&mutant); err == nil && mutantValidate {
		if res, err := helpers.Validate(mutant.Dna); err == nil {
			if !res {
				c.Data["mesaage"] = "DNA evaluated is from a human"
				c.Abort("403")
			}
			c.Data["json"] = map[string]interface{}{"Success": res, "Status": "200", "Message": "DNA evaluated is from a Mutant"}
			c.ServeJSON()
		}
	}
	if len(valid.Errors) > 0 {
		c.Data["mesaage"] = "Request failure: Input not valid - " + valid.Errors[0].Key + valid.Errors[0].Message
		c.Abort("422")
	}

	c.Data["mesaage"] = "Request failure"
	c.Abort("500")

}
