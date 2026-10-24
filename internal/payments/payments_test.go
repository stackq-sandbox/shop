package payments

import (
	"context"
	"testing"

	"github.com/stackq-sandbox/shop/internal/money"
)

func TestLifecycle(t *testing.T) {
	for _, name := range Names() {
		t.Run(name, func(t *testing.T) {
			p, err := Get(name)
			if err != nil {
				t.Fatal(err)
			}
			if p.Name() != name {
				t.Fatalf("Name() = %q, want %q", p.Name(), name)
			}
			ctx := context.Background()
			r, err := p.Authorize(ctx, Charge{Order: "o1", Amount: 5000, Currency: "USD"})
			if err != nil {
				t.Fatal(err)
			}
			if err := p.Refund(ctx, r.ID, 100); err == nil {
				t.Error("refund before capture should fail")
			}
			if err := p.Capture(ctx, r.ID); err != nil {
				t.Fatal(err)
			}
			if err := p.Capture(ctx, r.ID); err == nil {
				t.Error("capturing twice should fail")
			}
			if err := p.Refund(ctx, r.ID, 2000); err != nil {
				t.Fatal(err)
			}
			if err := p.Refund(ctx, r.ID, 4000); err == nil {
				t.Error("refunding more than was captured should fail")
			}
			if err := p.Refund(ctx, r.ID, 3000); err != nil {
				t.Errorf("refunding the remainder: %v", err)
			}
		})
	}
}

func TestAuthorizeRejects(t *testing.T) {
	for _, name := range Names() {
		p, _ := Get(name)
		ctx := context.Background()
		if _, err := p.Authorize(ctx, Charge{Amount: 0}); err == nil {
			t.Errorf("%s: zero amount should fail", name)
		}
		if _, err := p.Authorize(ctx, Charge{Amount: money.Cents(1 << 40)}); err == nil {
			t.Errorf("%s: over-limit amount should fail", name)
		}
		cctx, cancel := context.WithCancel(ctx)
		cancel()
		if _, err := p.Authorize(cctx, Charge{Amount: 100}); err == nil {
			t.Errorf("%s: cancelled context should fail", name)
		}
	}
}

func TestFeeIsPositive(t *testing.T) {
	for _, name := range Names() {
		p, _ := Get(name)
		if fee := p.Fee(10000); fee <= 0 || fee >= 10000 {
			t.Errorf("%s: Fee(10000) = %d, want between 0 and 10000", name, fee)
		}
	}
}

func TestGetUnknown(t *testing.T) {
	if _, err := Get("barter"); err == nil {
		t.Fatal("expected an error for an unknown provider")
	}
}
