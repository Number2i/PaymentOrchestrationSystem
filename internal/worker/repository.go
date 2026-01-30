package worker

import "database/sql"

type workerRepository interface {
	ClaimPayment() (paymentId string, err error)
}

type repo struct {
	db *sql.DB
}

// constructure func
func NewWorkerRepository(db *sql.DB) workerRepository {
	return &repo{db: db}
}
func (r *repo) ClaimPayment() (string, error) {
	var paymentId string
	err := r.db.QueryRow(`WITH row AS (
		SELECT payment_id
		FROM payment.payment_intent
		WHERE status = 'CREATED'
		ORDER BY created_at
		FOR UPDATE SKIP LOCKED
		LIMIT 1
	)
	UPDATE payment.payment_intent
	SET status = 'PROCESSING'
	FROM row
	WHERE payment.payment_intent.payment_id = row.payment_id
	RETURNING payment.payment_intent.payment_id;
	`).Scan(&paymentId)
	if err != nil {
		return "", err
	}
	return paymentId, nil
}
