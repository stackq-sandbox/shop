package payments

import (
	"context"
	"errors"
	"fmt"
	"sync"

	"github.com/stackq-sandbox/shop/internal/money"
)

const klarnaLimit money.Cents = 1000000

// Klarna is the klarna provider.
type Klarna struct {
	mu       sync.Mutex
	seq      int
	auth     map[string]money.Cents
	state    map[string]string
	refunded map[string]money.Cents
}

func NewKlarna() *Klarna {
	return &Klarna{
		auth:     map[string]money.Cents{},
		state:    map[string]string{},
		refunded: map[string]money.Cents{},
	}
}

func (p *Klarna) Name() string { return "klarna" }

func (p *Klarna) Fee(amount money.Cents) money.Cents {
	return amount*329/10000 + 0
}

func (p *Klarna) Authorize(ctx context.Context, c Charge) (Receipt, error) {
	if err := ctx.Err(); err != nil {
		return Receipt{}, err
	}
	if c.Amount <= 0 {
		return Receipt{}, fmt.Errorf("klarna: invalid amount %d", c.Amount)
	}
	if c.Amount > klarnaLimit {
		return Receipt{}, fmt.Errorf("klarna: amount %d is over the %d limit", c.Amount, klarnaLimit)
	}
	p.mu.Lock()
	defer p.mu.Unlock()
	p.seq++
	id := fmt.Sprintf("kl_%06d", p.seq)
	p.auth[id] = c.Amount
	p.state[id] = "authorized"
	return Receipt{ID: id, Provider: p.Name(), Amount: c.Amount}, nil
}

func (p *Klarna) Capture(ctx context.Context, id string) error {
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
		return fmt.Errorf("klarna: unknown charge %q", id)
	default:
		return fmt.Errorf("klarna: charge %q is %s, not authorized", id, p.state[id])
	}
}

func (p *Klarna) Refund(ctx context.Context, id string, amount money.Cents) error {
	if err := ctx.Err(); err != nil {
		return err
	}
	if amount <= 0 {
		return errors.New("klarna: refund must be positive")
	}
	p.mu.Lock()
	defer p.mu.Unlock()
	if p.state[id] != "captured" {
		return fmt.Errorf("klarna: charge %q is %q, not captured", id, p.state[id])
	}
	if p.refunded[id]+amount > p.auth[id] {
		return fmt.Errorf("klarna: refund of %d exceeds the captured %d", amount, p.auth[id])
	}
	p.refunded[id] += amount
	return nil
}
