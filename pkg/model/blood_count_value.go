package model

// BloodCountValue represents a blood count value associated with a disease.
type BloodCountValue struct {
	Disease     string  `json:"disease" db:"disease" validate:"required"`         // Name of the disease.
	BloodCount  string  `json:"blood_count" db:"blood_count" validate:"required"` // Type of blood count.
	Coefficient float32 `json:"coefficient" db:"coefficient" validate:"required"` // Coefficient value.
	Description string  `json:"description,omitempty" db:"description"`           // Description of the blood count value. Optional.
}
