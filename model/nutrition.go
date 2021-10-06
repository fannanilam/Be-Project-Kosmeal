package model

type Nutrition struct {
	//	ID                  float32 `json:"id"`
	Name                string  `json:"name"`
	Servingsizeg        float32 `json:"serving_size_g"`
	Fattotalg           float32 `json:"fat_total_g"`
	Fatsaturatedg       float32 `json:"fat_saturated_g"`
	Proteing            float32 `json:"protein_g"`
	Sodiummg            float32 `json:"sodium_mg"`
	Potassiummg         float32 `json:"potassium_mg"`
	Cholesterolmg       float32 `json:"cholesterol_mg"`
	Carbohydratestotalg float32 `json:"carbohydrates_total_g"`
	Fiberg              float32 `json:"fiber_g"`
	Sugarg              float32 `json:"sugar_g"`
}
