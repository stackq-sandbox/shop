package checkout

import (
	"encoding/json"
	"net/http"

	"github.com/stackq-sandbox/shop/internal/payments"
)

type placeRequest struct {
	Order Order `json:"order"`
}

type placeResponse struct {
	Quote   Quote             `json:"quote"`
	Receipt *payments.Receipt `json:"receipt,omitempty"`
}

func decode(r *http.Request) (Order, error) {
	var req placeRequest
	if err := json.NewDecoder(http.MaxBytesReader(nil, r.Body, 1<<20)).Decode(&req); err != nil {
		return Order{}, err
	}
	return req.Order, nil
}

func reply(w http.ResponseWriter, status int, v any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	_ = json.NewEncoder(w).Encode(v)
}

func fail(w http.ResponseWriter, status int, err error) {
	reply(w, status, map[string]string{"error": err.Error()})
}

// PreviewHandler prices an order without charging it.
func PreviewHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		fail(w, http.StatusMethodNotAllowed, errMethod)
		return
	}
	o, err := decode(r)
	if err != nil {
		fail(w, http.StatusBadRequest, err)
		return
	}
	q, err := Price(o)
	if err != nil {
		fail(w, http.StatusUnprocessableEntity, err)
		return
	}
	reply(w, http.StatusOK, placeResponse{Quote: q})
}

// Handler prices an order, then authorizes and captures the payment.
func Handler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		fail(w, http.StatusMethodNotAllowed, errMethod)
		return
	}
	o, err := decode(r)
	if err != nil {
		fail(w, http.StatusBadRequest, err)
		return
	}
	q, err := Price(o)
	if err != nil {
		fail(w, http.StatusUnprocessableEntity, err)
		return
	}
	p, err := payments.Get(o.Payment)
	if err != nil {
		fail(w, http.StatusUnprocessableEntity, err)
		return
	}
	rec, err := p.Authorize(r.Context(), payments.Charge{Order: o.ID, Amount: q.Total, Currency: "USD"})
	if err != nil {
		fail(w, http.StatusPaymentRequired, err)
		return
	}
	if err := p.Capture(r.Context(), rec.ID); err != nil {
		fail(w, http.StatusBadGateway, err)
		return
	}
	reply(w, http.StatusOK, placeResponse{Quote: q, Receipt: &rec})
}
