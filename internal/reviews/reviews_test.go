package reviews

import (
	"errors"
	"testing"

	"github.com/stackq-sandbox/shop/internal/money"
)

func TestTierFor(t *testing.T) {
	cases := []struct {
		spend money.Cents
		want  Tier
	}{{0, Bronze}, {49999, Bronze}, {50000, Silver}, {200000, Gold}, {5000000, Platinum}}
	for _, c := range cases {
		if got := TierFor(c.spend).Tier; got != c.want {
			t.Errorf("TierFor(%d) = %v, want %v", c.spend, got, c.want)
		}
	}
}

func TestEarnUsesTheTierHeldBeforeTheOrder(t *testing.T) {
	l := NewLedger()
	a, _, err := l.Earn("c1", "o1", 60000) // bronze rate: 600 points
	if err != nil || a.Points != 600 {
		t.Fatalf("first earn = %+v, %v", a, err)
	}
	a, _, _ = l.Earn("c1", "o2", 10000) // silver now: 100 * 2
	if a.Points != 800 {
		t.Fatalf("points = %d, want 800", a.Points)
	}
}

func TestEarnIsIdempotentPerOrder(t *testing.T) {
	l := NewLedger()
	if _, _, err := l.Earn("c1", "o1", 1000); err != nil {
		t.Fatal(err)
	}
	if _, _, err := l.Earn("c1", "o1", 1000); !errors.Is(err, ErrDuplicate) {
		t.Fatalf("second earn err = %v, want ErrDuplicate", err)
	}
}

func TestRedeemNeedsPoints(t *testing.T) {
	l := NewLedger()
	l.Earn("c1", "o1", 1000) // 10 points
	if _, _, err := l.Redeem("c1", "o2", 50); !errors.Is(err, ErrInsufficient) {
		t.Fatalf("err = %v, want ErrInsufficient", err)
	}
	a, _, err := l.Redeem("c1", "o2", 4)
	if err != nil || a.Points != 6 {
		t.Fatalf("redeem = %+v, %v", a, err)
	}
}

func TestRefundNeverGoesBelowZero(t *testing.T) {
	l := NewLedger()
	l.Earn("c1", "o1", 10000) // 100 points
	l.Redeem("c1", "o2", 90)
	a, e, err := l.Refund("c1", "o1", 10000)
	if err != nil || a.Points != 0 || e.Delta != -10 {
		t.Fatalf("refund = %+v %+v %v", a, e, err)
	}
	if a.LifetimeCents != 0 {
		t.Fatalf("lifetime = %d", a.LifetimeCents)
	}
}

func TestEntriesPerCustomer(t *testing.T) {
	l := NewLedger()
	l.Earn("c1", "o1", 1000)
	l.Earn("c2", "o2", 1000)
	if n := len(l.Entries("c1")); n != 1 {
		t.Fatalf("entries = %d", n)
	}
}
