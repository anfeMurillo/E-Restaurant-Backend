package payment

import (
	"context"
	"database/sql"
	paymentstatus "e-restaurant/models/enums/paymentStatus"
	model "e-restaurant/models/payment"
)

type PaymentRepository struct {
	db *sql.DB
}

func NewPaymentRepository(db *sql.DB) *PaymentRepository {
	return &PaymentRepository{db: db}
}

func (r *PaymentRepository) Create(ctx context.Context, p *model.Payment) (*model.Payment, error) {
	query := `INSERT INTO payments (user_id, order_id, amount, currency, method, status)
			  VALUES ($1, $2, $3, $4, $5, $6)
			  RETURNING payment_id, user_id, order_id, amount, currency, method, status, created_at`
	row := r.db.QueryRowContext(ctx, query, p.UserId, p.OrderId, p.Amount, p.Currency, p.Method, p.Status)
	var res model.Payment
	if err := row.Scan(&res.PaymentId, &res.UserId, &res.OrderId, &res.Amount, &res.Currency, &res.Method, &res.Status, &res.CreatedAt); err != nil {
		return nil, err
	}
	return &res, nil
}

func (r *PaymentRepository) GetById(ctx context.Context, paymentId int) (*model.Payment, error) {
	query := `SELECT payment_id, user_id, order_id, amount, currency, method, status, created_at FROM payments WHERE payment_id = $1`
	row := r.db.QueryRowContext(ctx, query, paymentId)
	var res model.Payment
	if err := row.Scan(&res.PaymentId, &res.UserId, &res.OrderId, &res.Amount, &res.Currency, &res.Method, &res.Status, &res.CreatedAt); err != nil {
		return nil, err
	}
	return &res, nil
}

func (r *PaymentRepository) UpdateStatus(ctx context.Context, paymentId int, status paymentstatus.PaymentStatus) error {
	query := `UPDATE payments SET status = $1 WHERE payment_id = $2`
	_, err := r.db.ExecContext(ctx, query, status, paymentId)
	return err
}

func (r *PaymentRepository) UpdateMethod(ctx context.Context, paymentId int, method paymentstatus.PaymentMethod) error {
	query := `UPDATE payments SET method = $1 WHERE payment_id = $2`
	_, err := r.db.ExecContext(ctx, query, method, paymentId)
	return err
}
