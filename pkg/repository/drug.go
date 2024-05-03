package repository

import (
	"fmt"
	"med/pkg/model"

	"github.com/jmoiron/sqlx"
)

type DrugRepository struct {
	db *sqlx.DB
}

func NewDrugRepository(db *sqlx.DB) *DrugRepository {
	return &DrugRepository{db: db}
}

// Create drug in database and get him from database
func (r *DrugRepository) CreateDrug(drug model.Drug) (model.Drug, error) {
	var createdDrug model.Drug
	query := fmt.Sprintf("INSERT INTO %s (id, name, dosage_form, active_ingredients, country, manufacturer, prescribing_order, description) VALUES ($1, $2, $3, $4, $5, $6, $7, $8) RETURNING *", drugTable)
	err := r.db.Get(&createdDrug, query,
		drug.Id,
		drug.Name,
		drug.DosageForm,
		drug.ActiveIngredients,
		drug.Country,
		drug.Manufacturer,
		drug.PrescribingOrder,
		drug.Description,
	)
	return createdDrug, err
}

// Get drug list from database
func (r *DrugRepository) GetDrugList() ([]model.Drug, error) {
	var drugList []model.Drug
	query := fmt.Sprintf("SELECT * FROM %s", drugTable)
	err := r.db.Select(&drugList, query)
	return drugList, err
}

// Get drug from database by ID
func (r *DrugRepository) GetDrugById(id string) (model.Drug, error) {
	var drug model.Drug
	query := fmt.Sprintf("SELECT * FROM %s WHERE id=$1", drugTable)
	err := r.db.Get(&drug, query, id)
	return drug, err
}

// Update drug data in database
func (r *DrugRepository) UpdateDrug(drug model.Drug) (model.Drug, error) {
	var updatedDrug model.Drug
	query := fmt.Sprintf("UPDATE %s SET name=$2, dosage_form=$3, active_ingredients=$4, country=$5, manufacturer=$6, prescribing_order=$7, description=$8 WHERE id=$1 RETURNING *", drugTable)
	err := r.db.Get(&updatedDrug, query,
		drug.Id,
		drug.Name,
		drug.DosageForm,
		drug.ActiveIngredients,
		drug.Country,
		drug.Manufacturer,
		drug.PrescribingOrder,
		drug.Description,
	)
	return updatedDrug, err
}

// Delete drug data from database
func (r *DrugRepository) DeleteDrug(id string) error {
	query := fmt.Sprintf("DELETE FROM %s WHERE id=$1", drugTable)
	_, err := r.db.Exec(query, id)
	return err
}
