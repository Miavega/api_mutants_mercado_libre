package controllers

import (
	"github.com/Miavega/api_mutants_mercado_libre/models"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
)

// StatsController operations for Stats
type StatsController struct {
	beego.Controller
}

// URLMapping ...
func (c *StatsController) URLMapping() {
	c.Mapping("GetStats", c.GetStats)
}

// GetStats ...
// @Title Get Stats
// @Description get Mutant Stats
// @Success 200 {object} models.Stats
// @Failure 404 not found resource
// @router / [get]
func (c *StatsController) GetStats() {
	v, err := models.GetStatsById(1)
	if err != nil {
		logs.Error(err)
		c.Data["mesaage"] = "Error service GetOne: The request contains an incorrect parameter or no record exists"
		c.Abort("404")
	} else {
		c.Data["json"] = map[string]interface{}{"Success": true, "Status": "200", "Message": "Request successful", "Data": v}
	}
	c.ServeJSON()
}
