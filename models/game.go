package models

type Game struct {
	Id                      string `json:"id"`
	DisplayName             string `json:"display_name"`
	Url                     string `json:"url"`
	LuckyNumberDigitCount   int    `json:"numar_cifre_noroc"`
	LuckyNumberMinMatchLen  int    `json:"noroc_min_match_len"`
	VariantMinNumbersCount  int    `json:"min_numere_per_varianta_jucata"`
	VariantsMaxCount        int    `json:"numar_max_variante"`
	VariantDrawNumbersCount int    `json:"numere_per_varianta_extrasa"`
	VariantMinNumber        int    `json:"min_value_numar_varianta"`
	VariantMaxNumber        int    `json:"max_value_numar_varianta"`
	LuckyNumberName         string `json:"nume_noroc"`
}
