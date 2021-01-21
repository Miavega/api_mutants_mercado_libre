package controllers

import (
	"encoding/json"
	"fmt"

	"github.com/Miavega/api_mutants/helpers"

	"github.com/Miavega/api_mutants/models"
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
// @Success 201 {object} models.Mutant
// @Failure 403 body is empty
// @router / [post]
func (c *MutantController) Post() {
	var mutant models.MutantStructure
	valid := validation.Validation{}
	json.Unmarshal(c.Ctx.Input.RequestBody, &mutant)
	if mutantValidate, err := valid.Valid(&mutant); err == nil {
		if mutantValidate {
			if res, err := helpers.Validate(mutant.Dna); err == nil {
				fmt.Println("Lleggue controler")
				if !res {
					c.Data["mesaage"] = "DNA evaluated is from a human"
					c.Abort("403")
				}
				c.Data["json"] = map[string]interface{}{"Success": res, "Status": "200", "Message": "DNA evaluated is from a Mutant"}
			} else {
				c.Data["mesaage"] = "Request failure"
				c.Abort("500")
			}
		} else {
			for _, err := range valid.Errors {
				c.Data["mesaage"] = "Request failure: Input not valid - " + err.Key + err.Message
			}
			c.Abort("422")
		}
	} else {
		c.Data["mesaage"] = "Request failure"
		c.Abort("500")
	}

	c.ServeJSON()
}
