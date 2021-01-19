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
	var mutant models.Mutant
	json.Unmarshal(c.Ctx.Input.RequestBody, &mutant)
	res := helpers.Validate(mutant.Dna)
	if !res {
		c.Abort("403")
	}
	c.Data["json"] = true
	c.ServeJSON()
}
