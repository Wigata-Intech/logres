package cart

import (
	"encoding/json"
	"net/http"

	"github.com/wigata-intech/logres/internal/api/service"
)

type CartHandler struct {
	cartService service.ICartService
}

func New(
	cartService service.ICartService,
) *CartHandler {
	return &CartHandler{
		cartService: cartService,
	}
}

func (h *CartHandler) Insert(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusInternalServerError)
	_ = json.NewEncoder(w).Encode(map[string]interface{}{
		"code":    http.StatusInternalServerError,
		"message": "internal server error",
	})
}

func (h *CartHandler) List(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusInternalServerError)
	_ = json.NewEncoder(w).Encode(map[string]interface{}{
		"code":    http.StatusInternalServerError,
		"message": "internal server error",
	})
}

func (h *CartHandler) Patch(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusInternalServerError)
	_ = json.NewEncoder(w).Encode(map[string]interface{}{
		"code":    http.StatusInternalServerError,
		"message": "internal server error",
	})
}

func (h *CartHandler) Delete(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusInternalServerError)
	_ = json.NewEncoder(w).Encode(map[string]interface{}{
		"code":    http.StatusInternalServerError,
		"message": "internal server error",
	})
}
