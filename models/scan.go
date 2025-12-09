package models

type ScanResult struct {
	GameId          string       `json:"game_id"`
	GameDate        string       `json:"game_date"`
	Variants        []Variant    `json:"variante"`
	LuckyNumber     *LuckyNumber `json:"noroc"`
	LuckyNumberName string       `json:"nume_noroc"`
}
