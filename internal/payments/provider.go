// Package payments defines the payment providers behind the new checkout.
// The providers are in-memory fakes with realistic fee and limit rules.
package payments

import (
	"context"
	"fmt"

	"github.com/stackq-sandbox/shop/internal/money"
)

type Charge struct {
	Order    string
	Amount   money.Cents
	Currency string
}

type Receipt struct {
	ID       string      `json:"id"`
	Provider string      `json:"provider"`
	Amount   money.Cents `json:"amount_cents"`
}

type Provider interface {
	Name() string
	// Fee is what the provider keeps from a captured amount.
	Fee(amount money.Cents) money.Cents
	Authorize(ctx context.Context, c Charge) (Receipt, error)
	Capture(ctx context.Context, id string) error
	Refund(ctx context.Context, id string, amount money.Cents) error
}

// Get returns a fresh provider by name.
func Get(name string) (Provider, error) {
	switch name {
	case "card":
		return NewCard(), nil
	case "paypal":
		return NewPayPal(), nil
	case "applepay":
		return NewApplePay(), nil
	case "klarna":
		return NewKlarna(), nil
	case "sepa":
		return NewSEPA(), nil
	case "ach":
		return NewACH(), nil
	default:
		return nil, fmt.Errorf("payments: unknown provider %q", name)
	}
}

// Names lists every registered provider.
func Names() []string {
	return []string{"card", "paypal", "applepay", "klarna", "sepa", "ach"}
}
