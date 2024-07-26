package repository

import (
	"creative-service-test/internal/models"
)

type DatabaseRepo interface {
	
	InsertOrder(order models.Order) (int, error)

	GetServiceID(category, subcategory string) (int, error)
	GetLocationID(division, district, upazila string) (int, error)	
}
