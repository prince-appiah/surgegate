package inventory

import (
	"errors"
	"testing"
)

func TestPurchase(t *testing.T) {
	inventory := &Inventory{Available: 1}

	err := inventory.Purchase()
	if err != nil {
		t.Fatalf("expected first purchase to succeed, got %v", err)
	}

	if inventory.Available != 0 {
		t.Errorf(
			"expected Available to be 0 after first purchase, got %d",
			inventory.Available,
		)
	}

	err = inventory.Purchase()
	if !errors.Is(err, ErrOutOfStock) {
		t.Errorf("expected ErrOutOfStock, got %v", err)
	}

	if inventory.Available != 0 {
		t.Errorf(
			"expected Available to remain 0 after failed purchase, got %d",
			inventory.Available,
		)
	}
}
