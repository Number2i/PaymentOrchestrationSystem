package http

import (
	"net/http"

	intentService "github.com/susidharan/payment-orchestration-system/internal/payment/intent"
)

func NewRouter(repo intentService.PaymentRepository) http.Handler {
	mux := http.NewServeMux()

	mux.HandleFunc("/payments", func(w http.ResponseWriter, r *http.Request) {
		intentService.CreatePayment(w, r, repo)
	})

	return mux
}
