package model

type Drug struct {
	Id                string `json:"id" db:"id"`
	Name              string `json:"name" db:"id"`
	DosageForm        string `json:"dosage-form" db:"dosage_form"`
	ActiveIngridients string `json:"active-ingridients" db:"active_ingridients"`
	Country           string `json:"country" db:"country"`
	Manufacturer      string `json:"manufacturer" db:"manufacturer"`
	PrescribingOrder  string `json:"prescribing-order" db:"prescribing_order"`
	Description       string `json:"description" db:"description"`
}
