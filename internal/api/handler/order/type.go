package order

import (
	"encoding/json"
	"net/http"

	"github.com/wigata-intech/logres/internal/api/service"
)

type OrderHandler struct {
	orderService service.IOrderService
}

func New(
	orderService service.IOrderService,
) *OrderHandler {
	return &OrderHandler{
		orderService: orderService,
	}
}

func (h *OrderHandler) Checkout(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusInternalServerError)
	_ = json.NewEncoder(w).Encode(map[string]interface{}{
		"code":    http.StatusInternalServerError,
		"message": "internal server error",
	})
}

func (h *OrderHandler) Check(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusInternalServerError)
	_ = json.NewEncoder(w).Encode(map[string]interface{}{
		"code":    http.StatusInternalServerError,
		"message": "internal server error",
	})
}
