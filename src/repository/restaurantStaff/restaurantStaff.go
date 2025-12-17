package restaurantstaff

import (
	"context"
	"database/sql"
	model "e-restaurant/models/restaurantStaff"
	userapp "e-restaurant/models/userApp"
)

type RestaurantStaffRepository struct {
	db *sql.DB
}

func NewRestaurantStaffRepository(db *sql.DB) *RestaurantStaffRepository {
	return &RestaurantStaffRepository{db: db}
}

func (r *RestaurantStaffRepository) Create(ctx context.Context, staff *model.RestaurantStaff) (*model.RestaurantStaff, error) {
	query := `INSERT INTO restaurant_staff (restaurant_id, user_id) VALUES ($1, $2) RETURNING restaurant_id, user_id`
	row := r.db.QueryRowContext(ctx, query, staff.RestaurantId, staff.UserId)
	var res model.RestaurantStaff
	if err := row.Scan(&res.RestaurantId, &res.UserId); err != nil {
		return nil, err
	}
	return &res, nil
}

func (r *RestaurantStaffRepository) Delete(ctx context.Context, restaurantId int, userId int) error {
	query := `DELETE FROM restaurant_staff WHERE restaurant_id = $1 AND user_id = $2`
	_, err := r.db.ExecContext(ctx, query, restaurantId, userId)
	return err
}

func (r *RestaurantStaffRepository) GetByRestaurant(ctx context.Context, restaurantId int, onlyActive bool) ([]*userapp.User, error) {
	query := `SELECT u.user_id, u.user_name, u.email, u.first_name, u.last_name, u.country_code, u.user_role, u.created_at, u.is_active, u.password_hash
			  FROM restaurant_staff rs
			  JOIN users u ON rs.user_id = u.user_id
			  WHERE rs.restaurant_id = $1`
	if onlyActive {
		query += ` AND u.is_active = true`
	}
	rows, err := r.db.QueryContext(ctx, query, restaurantId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var users []*userapp.User
	for rows.Next() {
		var u userapp.User
		if err := rows.Scan(&u.UserId, &u.UserName, &u.Email, &u.FirstName, &u.LastName, &u.CountryCode, &u.UserRole, &u.CreatedAt, &u.IsActive, &u.PasswordHash); err != nil {
			return nil, err
		}
		users = append(users, &u)
	}
	return users, nil
}
