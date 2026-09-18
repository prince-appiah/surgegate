package inventory

import (
	"errors"
	"testing"
)

func TestPurchase(t *testing.T) {
	tests := []struct {
		name              string
		startingAvailable int
		wantRemaining     int
		wantErr           error
	}{
		{
			name:              "purchase with available inventory",
			startingAvailable: 3,
			wantRemaining:     2,
			wantErr:           nil,
		},
		{
			name:              "purchase final item",
			startingAvailable: 1,
			wantRemaining:     0,
			wantErr:           nil,
		},
		{
			name:              "reject purchase when sold out",
			startingAvailable: 0,
			wantRemaining:     0,
			wantErr:           ErrOutOfStock,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			inventory := &Inventory{Available: test.startingAvailable}

			remaining, err := inventory.Purchase()

			if !errors.Is(err, test.wantErr) {
				t.Errorf("expected error %v, got %v", test.wantErr, err)
			}

			if remaining != test.wantRemaining {
				t.Errorf("expected remaining %d, got %d", test.wantRemaining, remaining)
			}

			if inventory.Available != test.wantRemaining {
				t.Errorf("expected stored inventory %d, got %d", test.wantRemaining, inventory.Available)
			}

		})

	}
}
