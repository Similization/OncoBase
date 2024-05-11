package model

import "github.com/guregu/null/v5"

type Drug struct {
	Id                null.String `json:"id" db:"id" binding:"required"`
	Name              null.String `json:"name" db:"name" binding:"required"`
	DosageForm        null.String `json:"dosage-form" db:"dosage_form" binding:"required"`
	ActiveIngredients null.String `json:"active-ingredients" db:"active_ingredients" binding:"required"`
	Country           null.String `json:"country" db:"country"`
	Manufacturer      null.String `json:"manufacturer" db:"manufacturer"`
	PrescribingOrder  null.String `json:"prescribing-order" db:"prescribing_order"`
	Description       null.String `json:"description" db:"description"`
}
