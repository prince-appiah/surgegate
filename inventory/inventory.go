package inventory

import (
	"errors"
)

type Inventory struct {
	Available int
}

var ErrOutOfStock = errors.New("item is out of stock")

const UnitsPerPurchase = 1

func (inventory *Inventory) Purchase() (int, error) {

	if inventory.Available == 0 {
		return inventory.Available, ErrOutOfStock
	}

	inventory.Available -= UnitsPerPurchase
	return inventory.Available, nil
}
