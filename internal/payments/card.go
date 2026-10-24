package payments

import (
	"context"
	"errors"
	"fmt"
	"sync"

	"github.com/stackq-sandbox/shop/internal/money"
)

const cardLimit money.Cents = 5000000

// Card is the card provider.
type Card struct {
	mu       sync.Mutex
	seq      int
	auth     map[string]money.Cents
	state    map[string]string
	refunded map[string]money.Cents
}

func NewCard() *Card {
	return &Card{
		auth:     map[string]money.Cents{},
		state:    map[string]string{},
		refunded: map[string]money.Cents{},
	}
}

func (p *Card) Name() string { return "card" }

func (p *Card) Fee(amount money.Cents) money.Cents {
	return amount*290/10000 + 30
}

func (p *Card) Authorize(ctx context.Context, c Charge) (Receipt, error) {
	if err := ctx.Err(); err != nil {
		return Receipt{}, err
	}
	if c.Amount <= 0 {
		return Receipt{}, fmt.Errorf("card: invalid amount %d", c.Amount)
	}
	if c.Amount > cardLimit {
		return Receipt{}, fmt.Errorf("card: amount %d is over the %d limit", c.Amount, cardLimit)
	}
	p.mu.Lock()
	defer p.mu.Unlock()
	p.seq++
	id := fmt.Sprintf("ch_%06d", p.seq)
	p.auth[id] = c.Amount
	p.state[id] = "authorized"
	return Receipt{ID: id, Provider: p.Name(), Amount: c.Amount}, nil
}

func (p *Card) Capture(ctx context.Context, id string) error {
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
		return fmt.Errorf("card: unknown charge %q", id)
	default:
		return fmt.Errorf("card: charge %q is %s, not authorized", id, p.state[id])
	}
}

func (p *Card) Refund(ctx context.Context, id string, amount money.Cents) error {
	if err := ctx.Err(); err != nil {
		return err
	}
	if amount <= 0 {
		return errors.New("card: refund must be positive")
	}
	p.mu.Lock()
	defer p.mu.Unlock()
	if p.state[id] != "captured" {
		return fmt.Errorf("card: charge %q is %q, not captured", id, p.state[id])
	}
	if p.refunded[id]+amount > p.auth[id] {
		return fmt.Errorf("card: refund of %d exceeds the captured %d", amount, p.auth[id])
	}
	p.refunded[id] += amount
	return nil
}
