package payments

import (
	"context"
	"errors"
	"fmt"
	"sync"

	"github.com/stackq-sandbox/shop/internal/money"
)

const achLimit money.Cents = 10000000

// ACH is the ach provider.
type ACH struct {
	mu       sync.Mutex
	seq      int
	auth     map[string]money.Cents
	state    map[string]string
	refunded map[string]money.Cents
}

func NewACH() *ACH {
	return &ACH{
		auth:     map[string]money.Cents{},
		state:    map[string]string{},
		refunded: map[string]money.Cents{},
	}
}

func (p *ACH) Name() string { return "ach" }

func (p *ACH) Fee(amount money.Cents) money.Cents {
	return amount*80/10000 + 0
}

func (p *ACH) Authorize(ctx context.Context, c Charge) (Receipt, error) {
	if err := ctx.Err(); err != nil {
		return Receipt{}, err
	}
	if c.Amount <= 0 {
		return Receipt{}, fmt.Errorf("ach: invalid amount %d", c.Amount)
	}
	if c.Amount > achLimit {
		return Receipt{}, fmt.Errorf("ach: amount %d is over the %d limit", c.Amount, achLimit)
	}
	p.mu.Lock()
	defer p.mu.Unlock()
	p.seq++
	id := fmt.Sprintf("ac_%06d", p.seq)
	p.auth[id] = c.Amount
	p.state[id] = "authorized"
	return Receipt{ID: id, Provider: p.Name(), Amount: c.Amount}, nil
}

func (p *ACH) Capture(ctx context.Context, id string) error {
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
		return fmt.Errorf("ach: unknown charge %q", id)
	default:
		return fmt.Errorf("ach: charge %q is %s, not authorized", id, p.state[id])
	}
}

func (p *ACH) Refund(ctx context.Context, id string, amount money.Cents) error {
	if err := ctx.Err(); err != nil {
		return err
	}
	if amount <= 0 {
		return errors.New("ach: refund must be positive")
	}
	p.mu.Lock()
	defer p.mu.Unlock()
	if p.state[id] != "captured" {
		return fmt.Errorf("ach: charge %q is %q, not captured", id, p.state[id])
	}
	if p.refunded[id]+amount > p.auth[id] {
		return fmt.Errorf("ach: refund of %d exceeds the captured %d", amount, p.auth[id])
	}
	p.refunded[id] += amount
	return nil
}
