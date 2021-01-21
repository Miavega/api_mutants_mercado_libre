package helpers

import (
	"strings"

	"github.com/Miavega/api_mutants/models"
)

// Validate validate if human is mutant
func Validate(dna []string) (bool, error) {
	var matriz [][]string
	var alteracionADN int = 0

	for _, n := range dna {
		matriz = append(matriz, strings.Split(n, ""))
	}

	for i := 0; i < len(matriz); i++ {
		for j := 0; j < len(matriz[0]); j++ {
			if j+3 < len(matriz[0]) {
				if isMutant(matriz[i][j], matriz[i][j+1], matriz[i][j+2], matriz[i][j+3]) {
					alteracionADN++
				}
			}
			if i+3 < len(matriz) {
				if isMutant(matriz[i][j], matriz[i+1][j], matriz[i+2][j], matriz[i+3][j]) {
					alteracionADN++
				}
			}
			if i+3 < len(matriz) && j+3 < len(matriz[0]) {
				if isMutant(matriz[i][j], matriz[i+1][j+1], matriz[i+2][j+2], matriz[i+3][j+3]) {
					alteracionADN++
				}
			}
			if j-3 >= 0 && i+3 < len(matriz) {
				if isMutant(matriz[i][j], matriz[i+1][j-1], matriz[i+2][j-2], matriz[i+3][j-3]) {
					alteracionADN++
				}
			}
			if alteracionADN > 1 {
				if _, err := saveMutantResult(dna, true); err == nil {
					go UpdateStats(true)
					return true, nil
				}
			}
		}
	}
	if _, err := saveMutantResult(dna, true); err == nil {
		go UpdateStats(false)
		return false, nil
	}

	return false, nil
}

func isMutant(a, b, c, d string) bool {
	if a == b && b == c && c == d {
		return true
	}
	return false
}

func saveMutantResult(dna []string, result bool) (int64, error) {
	var mutantDb models.Mutant
	mutantDb.Dna = strings.Join(dna, ",")
	mutantDb.IsMutant = result
	return models.AddMutant(&mutantDb)
}
