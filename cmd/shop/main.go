package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/stackq-sandbox/shop/internal/cart"
	"github.com/stackq-sandbox/shop/internal/money"
)

func routes() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/healthz", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "ok")
	})
	mux.HandleFunc("/cart/total", func(w http.ResponseWriter, r *http.Request) {
		var c cart.Cart
		c.Add("sku-1", 2, 1999)
		c.Add("sku-2", 1, 499)
		fmt.Fprintln(w, money.Format(c.Total()))
	})
	return mux
}

func main() {
	log.Fatal(http.ListenAndServe(":8080", routes()))
}
