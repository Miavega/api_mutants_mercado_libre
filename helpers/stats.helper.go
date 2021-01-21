package helpers

import (
	"github.com/Miavega/api_mutants/models"
)

//UpdateStats update statistics
func UpdateStats(result bool) {
	defer func() {
		if err := recover(); err != nil {
			panic(err)
		}
	}()

	if stats, err := models.GetStatsById(1); err == nil {
		if result {
			stats.CountMutantDna++
		} else {
			stats.CountHumanDna++
		}

		if stats.CountHumanDna != 0 {
			stats.Ratio = stats.CountMutantDna / stats.CountHumanDna
		} else {
			stats.Ratio = stats.CountMutantDna
		}

		models.UpdateStats(stats)

	}
}
