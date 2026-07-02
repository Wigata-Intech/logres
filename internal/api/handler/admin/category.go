package admin

import (
	"encoding/json"
	"net/http"
)

func (h *AdminHandler) CategoryCreate(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusInternalServerError)
	_ = json.NewEncoder(w).Encode(map[string]interface{}{
		"code":    http.StatusInternalServerError,
		"message": "internal server error",
	})
}

func (h *AdminHandler) CategoryList(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusInternalServerError)
	_ = json.NewEncoder(w).Encode(map[string]interface{}{
		"code":    http.StatusInternalServerError,
		"message": "internal server error",
	})
}

func (h *AdminHandler) CategoryUpdate(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusInternalServerError)
	_ = json.NewEncoder(w).Encode(map[string]interface{}{
		"code":    http.StatusInternalServerError,
		"message": "internal server error",
	})
}

func (h *AdminHandler) CategoryDelete(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusInternalServerError)
	_ = json.NewEncoder(w).Encode(map[string]interface{}{
		"code":    http.StatusInternalServerError,
		"message": "internal server error",
	})
}
