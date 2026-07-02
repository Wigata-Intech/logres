package product

import (
	"encoding/json"
	"net/http"

	"github.com/wigata-intech/logres/internal/api/service"
)

type ProductHandler struct {
	productService service.IProductService
}

func New(
	productService service.IProductService,
) *ProductHandler {
	return &ProductHandler{
		productService: productService,
	}
}

func (h *ProductHandler) List(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusInternalServerError)
	_ = json.NewEncoder(w).Encode(map[string]interface{}{
		"code":    http.StatusInternalServerError,
		"message": "internal server error",
	})
}

func (h *ProductHandler) Detail(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusInternalServerError)
	_ = json.NewEncoder(w).Encode(map[string]interface{}{
		"code":    http.StatusInternalServerError,
		"message": "internal server error",
	})
}
