package payments

import (
	"context"
	"errors"
	"fmt"
	"sync"

	"github.com/stackq-sandbox/shop/internal/money"
)

const sepaLimit money.Cents = 10000000

// SEPA is the sepa provider.
type SEPA struct {
	mu       sync.Mutex
	seq      int
	auth     map[string]money.Cents
	state    map[string]string
	refunded map[string]money.Cents
}

func NewSEPA() *SEPA {
	return &SEPA{
		auth:     map[string]money.Cents{},
		state:    map[string]string{},
		refunded: map[string]money.Cents{},
	}
}

func (p *SEPA) Name() string { return "sepa" }

func (p *SEPA) Fee(amount money.Cents) money.Cents {
	return amount*80/10000 + 35
}

func (p *SEPA) Authorize(ctx context.Context, c Charge) (Receipt, error) {
	if err := ctx.Err(); err != nil {
		return Receipt{}, err
	}
	if c.Amount <= 0 {
		return Receipt{}, fmt.Errorf("sepa: invalid amount %d", c.Amount)
	}
	if c.Amount > sepaLimit {
		return Receipt{}, fmt.Errorf("sepa: amount %d is over the %d limit", c.Amount, sepaLimit)
	}
	p.mu.Lock()
	defer p.mu.Unlock()
	p.seq++
	id := fmt.Sprintf("sd_%06d", p.seq)
	p.auth[id] = c.Amount
	p.state[id] = "authorized"
	return Receipt{ID: id, Provider: p.Name(), Amount: c.Amount}, nil
}

func (p *SEPA) Capture(ctx context.Context, id string) error {
	if err := ctx.Err(); err != nil {
		return err
	}
	p.mu.Lock()
	defer p.mu.Unlock()
	switch p.state[id] {
	case "authorized":
		p.state[id] = "captured"
		return nil
	case "":
		return fmt.Errorf("sepa: unknown charge %q", id)
	default:
		return fmt.Errorf("sepa: charge %q is %s, not authorized", id, p.state[id])
	}
}

func (p *SEPA) Refund(ctx context.Context, id string, amount money.Cents) error {
	if err := ctx.Err(); err != nil {
		return err
	}
	if amount <= 0 {
		return errors.New("sepa: refund must be positive")
	}
	p.mu.Lock()
	defer p.mu.Unlock()
	if p.state[id] != "captured" {
		return fmt.Errorf("sepa: charge %q is %q, not captured", id, p.state[id])
	}
	if p.refunded[id]+amount > p.auth[id] {
		return fmt.Errorf("sepa: refund of %d exceeds the captured %d", amount, p.auth[id])
	}
	p.refunded[id] += amount
	return nil
}
