package helpers

import (
	"fmt"

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
			stats.CountMutantDna = stats.CountMutantDna + 1
		} else {
			stats.CountHumanDna = stats.CountHumanDna + 1
		}
		if stats.CountHumanDna != 0 {
			stats.Ratio = float64(stats.CountMutantDna) / float64(stats.CountHumanDna)
		} else {
			stats.Ratio = float64(stats.CountMutantDna)
		}

		err := models.UpdateStats(stats)
		fmt.Println(err)

	}
}
