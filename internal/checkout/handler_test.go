package checkout

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

const body = `{"order":{"id":"o1","lines":[{"sku":"a","qty":1,"unit_cents":10000}],
"ship":{"line1":"1 Main St","city":"SF","region":"CA","postal":"94105","country":"US"},
"payment":"%s"}}`

func post(h http.HandlerFunc, payload string) *httptest.ResponseRecorder {
	rec := httptest.NewRecorder()
	h(rec, httptest.NewRequest(http.MethodPost, "/", strings.NewReader(payload)))
	return rec
}

func TestHandlerPlacesOrder(t *testing.T) {
	rec := post(Handler, strings.Replace(body, "%s", "card", 1))
	if rec.Code != http.StatusOK {
		t.Fatalf("status = %d, body %s", rec.Code, rec.Body)
	}
	var resp placeResponse
	if err := json.Unmarshal(rec.Body.Bytes(), &resp); err != nil {
		t.Fatal(err)
	}
	if resp.Quote.Total != 11284 {
		t.Errorf("total = %d, want 11284", resp.Quote.Total)
	}
	if resp.Receipt == nil || resp.Receipt.Provider != "card" {
		t.Errorf("receipt = %+v", resp.Receipt)
	}
}

func TestHandlerUnknownProvider(t *testing.T) {
	rec := post(Handler, strings.Replace(body, "%s", "barter", 1))
	if rec.Code != http.StatusUnprocessableEntity {
		t.Fatalf("status = %d, want 422", rec.Code)
	}
}

func TestHandlerRejectsBadJSONAndMethod(t *testing.T) {
	if rec := post(Handler, "{"); rec.Code != http.StatusBadRequest {
		t.Errorf("bad JSON status = %d, want 400", rec.Code)
	}
	rec := httptest.NewRecorder()
	Handler(rec, httptest.NewRequest(http.MethodGet, "/", nil))
	if rec.Code != http.StatusMethodNotAllowed {
		t.Errorf("GET status = %d, want 405", rec.Code)
	}
}

func TestPreviewDoesNotCharge(t *testing.T) {
	rec := post(PreviewHandler, strings.Replace(body, "%s", "card", 1))
	if rec.Code != http.StatusOK {
		t.Fatalf("status = %d, body %s", rec.Code, rec.Body)
	}
	var resp placeResponse
	if err := json.Unmarshal(rec.Body.Bytes(), &resp); err != nil {
		t.Fatal(err)
	}
	if resp.Receipt != nil {
		t.Errorf("preview must not return a receipt: %+v", resp.Receipt)
	}
}
