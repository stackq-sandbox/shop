package events

import "testing"

func TestPublishReachesOnlyTheTopicsSubscribersInOrder(t *testing.T) {
	b := NewBus()
	var got []string
	b.Subscribe("stock.low", func(e Event) { got = append(got, "a:"+e.Payload.(string)) })
	b.Subscribe("stock.low", func(e Event) { got = append(got, "b:"+e.Payload.(string)) })
	b.Subscribe("order.placed", func(Event) { got = append(got, "other") })

	if n := b.Publish(Event{Topic: "stock.low", Payload: "sku-1"}); n != 2 {
		t.Fatalf("delivered to %d subscribers, want 2", n)
	}
	if len(got) != 2 || got[0] != "a:sku-1" || got[1] != "b:sku-1" {
		t.Fatalf("got %v", got)
	}
	if n := b.Publish(Event{Topic: "nobody"}); n != 0 {
		t.Fatalf("delivered to %d subscribers of an unknown topic", n)
	}
}
