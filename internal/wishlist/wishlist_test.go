package wishlist

import (
	"reflect"
	"testing"
)

func TestAddRemove(t *testing.T) {
	var l List
	l.Add("b")
	l.Add("a")
	l.Add("a")
	l.Remove("b")
	if got := l.Items(); !reflect.DeepEqual(got, []string{"a"}) {
		t.Fatalf("items = %v", got)
	}
}
