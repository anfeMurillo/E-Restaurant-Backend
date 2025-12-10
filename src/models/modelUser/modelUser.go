package modeluser

import "time"

type Role string

const (
	RoleClient  Role = "client"
	RoleAdmin   Role = "admin"
	RoleChef    Role = "chef"
	RoleCashier Role = "cashier"
)

type User struct {
	UserId       int       `db:"user_id" json:"user_id"`
	UserName     string    `db:"user_name" json:"user_name"`
	Email        string    `db:"email" json:"email"`
	FirstName    string    `db:"first_name" json:"first_name"`
	LastName     string    `db:"last_name" json:"last_name"`
	CountryCode  string    `db:"country_code" json:"country_code"`
	UserRole     Role      `db:"user_role" json:"user_role"`
	CreatedAt    time.Time `db:"created_at" json:"created_at"`
	IsActive     bool      `db:"is_active" json:"is_active"`
	PasswordHash string    `db:"password_hash" json:"password_hash"`
}
