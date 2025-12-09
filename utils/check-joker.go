package utils

import (
	"fmt"
	"loto-suite/backend/generics"
	"loto-suite/backend/models"
)

const maxMatchCount = 5

func CheckBiletJoker(checkResult *models.CheckResult) {
	game, _ := GetGameById("joker")
	VerificareNorocJoker(checkResult.LuckyNumber, checkResult.DrawResult.LuckyNumber, game.LuckyNumberDigitCount, game.LuckyNumberMinMatchLen)
	varianteJucateLen := len(checkResult.VarianteJucate)

	if varianteJucateLen == 0 {
		variantaJucata := models.Variant{
			Numbers:     []models.Number{},
			WinsRegular: getDefaultCategoriiCastigVarianteJoker(),
			WinsSpecial: getDefaultCategoriiCastigVarianteJoker(),
		}

		checkResult.VarianteJucate = append(checkResult.VarianteJucate, variantaJucata)
	} else {
		for i := range checkResult.VarianteJucate {
			// Reset Castigator flags once before verification
			for j := range checkResult.VarianteJucate[i].Numbers {
				checkResult.VarianteJucate[i].Numbers[j].IsWinner = false
			}

			VerificareVariantaJoker(&checkResult.VarianteJucate[i], checkResult.DrawResult.VariantRegular, game.VariantMinNumbersCount, game.VariantDrawNumbersCount)
			VerificareVariantaJoker(&checkResult.VarianteJucate[i], checkResult.DrawResult.VariantSpecial, game.VariantMinNumbersCount, game.VariantDrawNumbersCount)
		}
	}
}

func VerificareJoker(variantaJucata *models.Variant, variantaExtrasa *models.Variant) bool {
	if variantaJucata == nil || variantaExtrasa == nil {
		return false
	}

	jokerJucat := variantaJucata.Numbers[len(variantaJucata.Numbers)-1]
	jokerCastigator := variantaExtrasa.Numbers[len(variantaExtrasa.Numbers)-1]

	jokerMatch := jokerJucat.Value == jokerCastigator.Value

	return jokerMatch
}

func VerificareVariantaJoker(variantaJucata *models.Variant, variantaExtrasa *models.Variant, minNumerePerVariantaJucata int, numerePerVariantaExtrasa int) {
	if variantaJucata == nil || variantaExtrasa == nil {
		return
	}

	isValidTicket := len(variantaJucata.Numbers) >= minNumerePerVariantaJucata+1
	isValidDraw := variantaExtrasa.Id != -1 && len(variantaExtrasa.Numbers) == numerePerVariantaExtrasa

	if !isValidTicket || !isValidDraw {
		switch variantaExtrasa.Id {
		case 1:
			variantaJucata.WinsRegular = getDefaultCategoriiCastigVarianteJoker()
		case 2:
			variantaJucata.WinsSpecial = getDefaultCategoriiCastigVarianteJoker()
		}

		return
	}

	matchCount := 0
	jokerMatch := VerificareJoker(variantaJucata, variantaExtrasa)
	if jokerMatch {
		variantaJucata.Numbers[len(variantaJucata.Numbers)-1].IsWinner = true
	}

	for _, n := range variantaExtrasa.Numbers[:numerePerVariantaExtrasa-1] {
		index := generics.IndexOf(variantaJucata.Numbers[:numerePerVariantaExtrasa-1], func(numarVarianta models.Number) bool {
			return numarVarianta.Value == n.Value
		})

		if index != -1 {
			matchCount++
			variantaJucata.Numbers[index].IsWinner = true
		}
	}

	highestWinIndex := 0
	switch {
	case matchCount == 5 && jokerMatch:
		highestWinIndex = 1
	case matchCount == 5:
		highestWinIndex = 2
	case matchCount == 4 && jokerMatch:
		highestWinIndex = 3
	case matchCount == 4:
		highestWinIndex = 4
	case matchCount == 3 && jokerMatch:
		highestWinIndex = 5
	case matchCount == 3:
		highestWinIndex = 6
	case matchCount == 2 && jokerMatch:
		highestWinIndex = 7
	case matchCount == 1 && jokerMatch:
		highestWinIndex = 8
	}

	idx := 1
	for n := maxMatchCount; n >= 3; n-- {
		castig := models.Win{
			Id:          fmt.Sprintf("%v (%v/%v%v)", generics.RomanNumbers[idx], n, maxMatchCount, "+J"),
			Description: getDescriereCategorieNumere(n, maxMatchCount, true),
			IsWinner:    idx == highestWinIndex,
		}

		switch variantaExtrasa.Id {
		case 1:
			variantaJucata.WinsRegular = append(variantaJucata.WinsRegular, castig)
		case 2:
			variantaJucata.WinsSpecial = append(variantaJucata.WinsSpecial, castig)
		}

		idx++

		castig = models.Win{
			Id:          fmt.Sprintf("%v (%v/%v%v)", generics.RomanNumbers[idx], n, maxMatchCount, ""),
			Description: getDescriereCategorieNumere(n, maxMatchCount, false),
			IsWinner:    idx == highestWinIndex,
		}

		switch variantaExtrasa.Id {
		case 1:
			variantaJucata.WinsRegular = append(variantaJucata.WinsRegular, castig)
		case 2:
			variantaJucata.WinsSpecial = append(variantaJucata.WinsSpecial, castig)
		}

		idx++
	}

	castig := models.Win{
		Id:          fmt.Sprintf("%v (%v/%v%v)", generics.RomanNumbers[idx], 2, maxMatchCount, "+J"),
		Description: getDescriereCategorieNumere(2, maxMatchCount, true),
		IsWinner:    idx == highestWinIndex,
	}

	switch variantaExtrasa.Id {
	case 1:
		variantaJucata.WinsRegular = append(variantaJucata.WinsRegular, castig)
	case 2:
		variantaJucata.WinsSpecial = append(variantaJucata.WinsSpecial, castig)
	}

	idx++
	castig = models.Win{
		Id:          fmt.Sprintf("%v (%v/%v%v)", generics.RomanNumbers[idx], 1, maxMatchCount, "+J"),
		Description: getDescriereCategorieNumere(1, maxMatchCount, true),
		IsWinner:    idx == highestWinIndex,
	}

	switch variantaExtrasa.Id {
	case 1:
		variantaJucata.WinsRegular = append(variantaJucata.WinsRegular, castig)
	case 2:
		variantaJucata.WinsSpecial = append(variantaJucata.WinsSpecial, castig)
	}
}

func VerificareNorocJoker(norocJucat *models.LuckyNumber, norocCastigator *models.LuckyNumber, norocLen int, minNorocMatchLen int) {
	foundWinner := false

	numarNorocJucat := norocJucat.Value
	numarNorocCastigator := norocCastigator.Value

	for n := norocLen; n >= minNorocMatchLen; n-- {
		descriere := ""
		if n == norocLen {
			descriere = fmt.Sprintf("Toate cele %v cifre ale numarului (in ordine)", n)
		} else {
			descriere = fmt.Sprintf("Primele sau ultimele %v cifre ale numarului (in ordine)", n)
		}

		isWinner := false

		if !foundWinner {
			if len(numarNorocJucat) == len(numarNorocCastigator) && len(numarNorocJucat) == norocLen {
				isWinner = numarNorocJucat[:n] == numarNorocCastigator[:n] || numarNorocJucat[norocLen-n:] == numarNorocCastigator[norocLen-n:]
			}
		}

		norocJucat.Wins = append(norocJucat.Wins,
			models.Win{
				Id:          fmt.Sprintf("%v", generics.RomanNumbers[norocLen-n+1]),
				Description: descriere,
				IsWinner:    !foundWinner && isWinner,
			})

		if isWinner {
			foundWinner = true
		}
	}

	norocJucat.IsWinner = foundWinner
}

func getDescriereCategorieNumere(numbers int, maxNumbers int, includeJoker bool) string {
	result := ""

	switch numbers {
	case maxNumbers:
		result = fmt.Sprintf("Toate cele %v numere din primul set", numbers)
	case 1:
		result = fmt.Sprintf("Oricare numar din cele %v ale primului set", maxNumbers)
	default:
		result = fmt.Sprintf("Oricare %v numere din cele %v ale primului set", numbers, maxNumbers)
	}

	if includeJoker {
		result += " si JOKER-ul"
	}

	return result
}

func getDefaultCategoriiCastigVarianteJoker() []models.Win {
	castiguri := []models.Win{}

	idx := 1
	for n := maxMatchCount; n >= 3; n-- {
		castiguri = append(castiguri, models.Win{
			Id:          fmt.Sprintf("%v (%v/%v%v)", generics.RomanNumbers[idx], n, maxMatchCount, "+J"),
			Description: getDescriereCategorieNumere(n, maxMatchCount, true),
			IsWinner:    false,
		})
		idx++

		castiguri = append(castiguri, models.Win{
			Id:          fmt.Sprintf("%v (%v/%v%v)", generics.RomanNumbers[idx], n, maxMatchCount, ""),
			Description: getDescriereCategorieNumere(n, maxMatchCount, false),
			IsWinner:    false,
		})
		idx++
	}

	castiguri = append(castiguri, models.Win{
		Id:          fmt.Sprintf("%v (%v/%v%v)", generics.RomanNumbers[idx], 2, maxMatchCount, "+J"),
		Description: getDescriereCategorieNumere(2, maxMatchCount, true),
		IsWinner:    false,
	})

	idx++

	castiguri = append(castiguri, models.Win{
		Id:          fmt.Sprintf("%v (%v/%v%v)", generics.RomanNumbers[idx], 1, maxMatchCount, "+J"),
		Description: getDescriereCategorieNumere(1, maxMatchCount, true),
		IsWinner:    false,
	})

	return castiguri
}
