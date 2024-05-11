package model

import "github.com/guregu/null/v5"

// BloodCountValue represents a blood count value associated with a disease.
type BloodCountValue struct {
	Disease     null.String `json:"disease" db:"disease" binding:"required"`         // Name of the disease.
	BloodCount  null.String `json:"blood_count" db:"blood_count" binding:"required"` // Type of blood count.
	Coefficient null.Float  `json:"coefficient" db:"coefficient" binding:"required"` // Coefficient value.
	Description null.String `json:"description,omitempty" db:"description"`          // Description of the blood count value. Optional.
}
