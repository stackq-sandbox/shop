package coupons

import (
	"errors"
	"sync"

	"github.com/stackq-sandbox/shop/internal/money"
)

// Account is one customer's standing.
type Account struct {
	CustomerID    string
	Points        int64
	LifetimeCents money.Cents
}

// Entry is one change to an account.
type Entry struct {
	ID         string
	CustomerID string
	OrderID    string
	Delta      int64
	Reason     string
}

// ErrDuplicate means the order already earned (or was refunded for) points.
var ErrDuplicate = errors.New("coupons: order already recorded")

// ErrInsufficient means the account has fewer points than a redemption needs.
var ErrInsufficient = errors.New("coupons: not enough points")

// Ledger keeps accounts and their entries. It is safe for concurrent use.
type Ledger struct {
	mu       sync.Mutex
	accounts map[string]*Account
	entries  []Entry
	seen     map[string]bool
}

// NewLedger returns an empty ledger.
func NewLedger() *Ledger {
	return &Ledger{accounts: map[string]*Account{}, seen: map[string]bool{}}
}

func (l *Ledger) account(id string) *Account {
	a, ok := l.accounts[id]
	if !ok {
		a = &Account{CustomerID: id}
		l.accounts[id] = a
	}
	return a
}

func (l *Ledger) record(a *Account, orderID, reason string, delta int64) (Entry, error) {
	key := orderID + "/" + reason
	if l.seen[key] {
		return Entry{}, ErrDuplicate
	}
	l.seen[key] = true
	e := Entry{ID: key, CustomerID: a.CustomerID, OrderID: orderID, Delta: delta, Reason: reason}
	l.entries = append(l.entries, e)
	a.Points += delta
	return e, nil
}

// Earn credits points for an order. The tier is the one the customer held
// before this order, so a big order does not pay itself a bonus.
func (l *Ledger) Earn(customer, orderID string, total money.Cents) (Account, Entry, error) {
	l.mu.Lock()
	defer l.mu.Unlock()
	a := l.account(customer)
	rate := TierFor(a.LifetimeCents).PointsPer
	e, err := l.record(a, orderID, "earn", int64(total/100)*int64(rate))
	if err != nil {
		return *a, Entry{}, err
	}
	a.LifetimeCents += total
	return *a, e, nil
}

// Redeem spends points on an order.
func (l *Ledger) Redeem(customer, orderID string, points int64) (Account, Entry, error) {
	l.mu.Lock()
	defer l.mu.Unlock()
	a := l.account(customer)
	if a.Points < points {
		return *a, Entry{}, ErrInsufficient
	}
	e, err := l.record(a, orderID, "redeem", -points)
	return *a, e, err
}

// Refund takes back what an order earned, never below zero.
func (l *Ledger) Refund(customer, orderID string, total money.Cents) (Account, Entry, error) {
	l.mu.Lock()
	defer l.mu.Unlock()
	a := l.account(customer)
	var earned int64
	for _, e := range l.entries {
		if e.OrderID == orderID && e.Reason == "earn" {
			earned = e.Delta
		}
	}
	if earned > a.Points {
		earned = a.Points
	}
	e, err := l.record(a, orderID, "refund", -earned)
	if err != nil {
		return *a, Entry{}, err
	}
	a.LifetimeCents -= total
	if a.LifetimeCents < 0 {
		a.LifetimeCents = 0
	}
	return *a, e, nil
}

// Entries returns a customer's entries, oldest first.
func (l *Ledger) Entries(customer string) []Entry {
	l.mu.Lock()
	defer l.mu.Unlock()
	var out []Entry
	for _, e := range l.entries {
		if e.CustomerID == customer {
			out = append(out, e)
		}
	}
	return out
}
