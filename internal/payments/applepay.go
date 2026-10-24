package payments

import (
	"context"
	"errors"
	"fmt"
	"sync"

	"github.com/stackq-sandbox/shop/internal/money"
)

const applepayLimit money.Cents = 1000000

// ApplePay is the applepay provider.
type ApplePay struct {
	mu       sync.Mutex
	seq      int
	auth     map[string]money.Cents
	state    map[string]string
	refunded map[string]money.Cents
}

func NewApplePay() *ApplePay {
	return &ApplePay{
		auth:     map[string]money.Cents{},
		state:    map[string]string{},
		refunded: map[string]money.Cents{},
	}
}

func (p *ApplePay) Name() string { return "applepay" }

func (p *ApplePay) Fee(amount money.Cents) money.Cents {
	return amount*250/10000 + 20
}

func (p *ApplePay) Authorize(ctx context.Context, c Charge) (Receipt, error) {
	if err := ctx.Err(); err != nil {
		return Receipt{}, err
	}
	if c.Amount <= 0 {
		return Receipt{}, fmt.Errorf("applepay: invalid amount %d", c.Amount)
	}
	if c.Amount > applepayLimit {
		return Receipt{}, fmt.Errorf("applepay: amount %d is over the %d limit", c.Amount, applepayLimit)
	}
	p.mu.Lock()
	defer p.mu.Unlock()
	p.seq++
	id := fmt.Sprintf("ap_%06d", p.seq)
	p.auth[id] = c.Amount
	p.state[id] = "authorized"
	return Receipt{ID: id, Provider: p.Name(), Amount: c.Amount}, nil
}

func (p *ApplePay) Capture(ctx context.Context, id string) error {
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
		return fmt.Errorf("applepay: unknown charge %q", id)
	default:
		return fmt.Errorf("applepay: charge %q is %s, not authorized", id, p.state[id])
	}
}

func (p *ApplePay) Refund(ctx context.Context, id string, amount money.Cents) error {
	if err := ctx.Err(); err != nil {
		return err
	}
	if amount <= 0 {
		return errors.New("applepay: refund must be positive")
	}
	p.mu.Lock()
	defer p.mu.Unlock()
	if p.state[id] != "captured" {
		return fmt.Errorf("applepay: charge %q is %q, not captured", id, p.state[id])
	}
	if p.refunded[id]+amount > p.auth[id] {
		return fmt.Errorf("applepay: refund of %d exceeds the captured %d", amount, p.auth[id])
	}
	p.refunded[id] += amount
	return nil
}
