package dbrepo

import (
	"context"
	"creative-service-test/internal/models"
	"time"
)

// InsertOrder inserts new order to the database and returns its id
func (p *postgresDBRepo) InsertOrder(order models.Order) (int, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()


	stmt := `INSERT INTO orders (
			client_name, client_mobile, location_id, area, date, preferred_time, service_id, problem_summary, attachment, created_at, updated_at  
		)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11) returning id
	`

	var id int
	err := p.DB.QueryRowContext(ctx, stmt,
		order.ClientName,
		order.ClientMobile,
		order.LocationID,
		order.Area,
		order.Date,
		order.PreferredTime,
		order.ServiceID,
		order.ProblemSummary,
		order.ImgFile,
		time.Now(),
		time.Now(),
	).Scan(&id)

	if err != nil {
		return 0, err
	}

	return id, nil
}

// GetLocationID returns location id based on division-district-upazila
func (p *postgresDBRepo) GetLocationID(division, district, upazila string) (int, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	var id int
	query :=  `
		SELECT id
		FROM locations 
		WHERE division = $1 AND district = $2 AND upazila = $3
	`
	err := p.DB.QueryRowContext(ctx, query, division, district, upazila).Scan(&id)


	return id, err
}
// GetServiceID returns service id based on service category-name
func (p *postgresDBRepo) GetServiceID(category, subcategory string) (int, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	var id int
	query :=  `
		SELECT id
		FROM services
		WHERE category_name = $1 AND service_name = $2
	`
	err := p.DB.QueryRowContext(ctx, query, category, subcategory).Scan(&id)


	return id, err
}