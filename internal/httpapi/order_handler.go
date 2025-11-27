package httpapi

import (
	"encoding/json"
	"log"
	"logiTrack/internal/order"
	"logiTrack/internal/validation"
	"net/http"
	"strconv"
	"strings"
)

//type OrderService interface {
//	Save(order *order.Order)
//	GetAll() []order.Order
//	GetByID(id int) (*order.Order, int, error)
//	Update(id int, status string) error
//	Delete(id int) error
//}

type OrderHandler struct {
	order order.StorageIntf
}

func NewOrderHandler(o order.StorageIntf) *OrderHandler {
	return &OrderHandler{
		order: o,
	}
}

func (h *OrderHandler) RegisterRoutes(mux *http.ServeMux) {
	mux.HandleFunc("/orders", h.orders)
	mux.HandleFunc("/orders/", h.ordersByID)
}

func (h *OrderHandler) orders(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		h.getAll(w, r)
	case http.MethodPost:
		AuthMiddleware(http.HandlerFunc(h.create)).ServeHTTP(w, r)
		//h.create(w, r)
	default:
		http.Error(w, "Unsupported method", http.StatusMethodNotAllowed)
	}
}

func (h *OrderHandler) ordersByID(w http.ResponseWriter, r *http.Request) {
	idStr := strings.TrimPrefix(r.URL.Path, "/orders/")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Некорректный ID", http.StatusBadRequest)
		return
	}

	switch r.Method {
	case http.MethodGet:
		getOrder, _, err := h.order.GetByID(id)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		writeJSON(w, getOrder)

	case http.MethodPut:
		AuthMiddleware(
			http.HandlerFunc(
				func(w http.ResponseWriter, r *http.Request) {
					var updateReq order.UpdateReq
					if err := json.NewDecoder(r.Body).Decode(&updateReq); err != nil {
						http.Error(w, "Невалидный JSON", http.StatusBadRequest)
						return
					}
					if err := validation.ValidatingOrder(updateReq); err != nil {

						w.Header().Set("Content-Type", "application/json")
						w.WriteHeader(http.StatusBadRequest)
						json.NewEncoder(w).Encode(
							map[string]interface{}{
								"error":  "Validation Failed",
								"errors": err,
							},
						)
						return
					}

					if err := h.order.Update(id, updateReq.Status); err != nil {
						http.Error(w, err.Error(), http.StatusNotFound)
					}
					w.WriteHeader(http.StatusNoContent)
				},
			),
		).ServeHTTP(w, r)

	case http.MethodDelete:
		AuthMiddleware(
			http.HandlerFunc(
				func(w http.ResponseWriter, r *http.Request) {
					if err := h.order.Delete(id); err != nil {
						http.Error(w, err.Error(), http.StatusInternalServerError)
						return
					}
					w.WriteHeader(http.StatusNoContent)
				},
			),
		).ServeHTTP(w, r)

	default:
		http.Error(w, "Метод не поддерживается", http.StatusMethodNotAllowed)
	}
}

func (h *OrderHandler) getAll(w http.ResponseWriter, r *http.Request) {
	orders := h.order.GetAll()
	writeJSON(w, orders)
}

func (h *OrderHandler) create(w http.ResponseWriter, r *http.Request) {
	var newOrder order.Order
	if err := json.NewDecoder(r.Body).Decode(&newOrder); err != nil {
		http.Error(w, `{"error": "invalid JSON"}`, http.StatusBadRequest)
		return
	}
	if err := validation.ValidatingOrder(newOrder); err != nil {
		//var errors []string
		//for _, e := range err.(validator.ValidationErrors) {
		//	errors = append(
		//		errors, fmt.Sprintf(
		//			"Field '%s' failed validation '%s', value: '%s'", e.Field(),
		//			e.Tag(), e.Value(),
		//		),
		//	)
		//}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(
			map[string]interface{}{
				"error":  "Validation Failed",
				"errors": err,
			},
		)
		return
	}

	allOrders := h.order.GetAll()
	newOrder.ID = len(allOrders) + 1

	order.SafeFuncSaveOrder(&newOrder, h.order)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	if err := json.NewEncoder(w).Encode(newOrder); err != nil {
		log.Printf("failed to encode new order: %s", err)
	}
}

func writeJSON(w http.ResponseWriter, v any) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(v)
}
