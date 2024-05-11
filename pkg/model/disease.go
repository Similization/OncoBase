package model

import "github.com/guregu/null/v5"

type Disease struct {
	Id          null.String `json:"id" db:"id" binding:"required"`
	Description null.String `json:"description" db:"description"`
}
