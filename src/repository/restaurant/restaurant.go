package restaurant

import (
	"context"
	"database/sql"
	model "e-restaurant/models/restaurant"
)

type RestaurantRepository struct {
	db *sql.DB
}

func NewRestaurantRepository(db *sql.DB) *RestaurantRepository {
	return &RestaurantRepository{db: db}
}

func (r *RestaurantRepository) Create(ctx context.Context, rest *model.Restaurant) (*model.Restaurant, error) {
	query := `INSERT INTO restaurants (restaurant_name, address, is_active, is_open)
			  VALUES ($1, $2, $3, $4)
			  RETURNING restaurant_id, restaurant_name, address, is_active, is_open`
	row := r.db.QueryRowContext(ctx, query, rest.RestaurantName, rest.Address, rest.IsActive, rest.IsOpen)
	var res model.Restaurant
	if err := row.Scan(&res.RestaurantId, &res.RestaurantName, &res.Address, &res.IsActive, &res.IsOpen); err != nil {
		return nil, err
	}
	return &res, nil
}

func (r *RestaurantRepository) UpdateName(ctx context.Context, restaurantId int, new string) error {
	query := `UPDATE restaurants SET restaurant_name = $1 WHERE restaurant_id = $2`
	_, err := r.db.ExecContext(ctx, query, new, restaurantId)
	return err
}

func (r *RestaurantRepository) UpdateAddress(ctx context.Context, restaurantId int, new string) error {
	query := `UPDATE restaurants SET address = $1 WHERE restaurant_id = $2`
	_, err := r.db.ExecContext(ctx, query, new, restaurantId)
	return err
}

func (r *RestaurantRepository) UpdateState(ctx context.Context, restaurantId int, new bool) error {
	query := `UPDATE restaurants SET is_active = $1 WHERE restaurant_id = $2`
	_, err := r.db.ExecContext(ctx, query, new, restaurantId)
	return err
}

func (r *RestaurantRepository) ToggleOpen(ctx context.Context, restaurantId int) error {
	// Toggle boolean value
	query := `UPDATE restaurants SET is_open = NOT is_open WHERE restaurant_id = $1`
	_, err := r.db.ExecContext(ctx, query, restaurantId)
	return err
}

func (r *RestaurantRepository) GetById(ctx context.Context, restaurantId int) (*model.Restaurant, error) {
	query := `SELECT restaurant_id, restaurant_name, address, is_active, is_open FROM restaurants WHERE restaurant_id = $1`
	row := r.db.QueryRowContext(ctx, query, restaurantId)
	var res model.Restaurant
	if err := row.Scan(&res.RestaurantId, &res.RestaurantName, &res.Address, &res.IsActive, &res.IsOpen); err != nil {
		return nil, err
	}
	return &res, nil
}

func (r *RestaurantRepository) Delete(ctx context.Context, restaurantId int) error {
	query := `DELETE FROM restaurants WHERE restaurant_id = $1`
	_, err := r.db.ExecContext(ctx, query, restaurantId)
	return err
}

func (r *RestaurantRepository) GetAll(ctx context.Context) ([]*model.Restaurant, error) {
	query := `SELECT restaurant_id, restaurant_name, address, is_active, is_open FROM restaurants`
	rows, err := r.db.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var result []*model.Restaurant
	for rows.Next() {
		var rs model.Restaurant
		if err := rows.Scan(&rs.RestaurantId, &rs.RestaurantName, &rs.Address, &rs.IsActive, &rs.IsOpen); err != nil {
			return nil, err
		}
		result = append(result, &rs)
	}
	return result, nil
}
