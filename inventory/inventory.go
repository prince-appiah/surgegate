package inventory

import (
	"errors"
)

type Inventory struct {
	Available int
}

var ErrOutOfStock = errors.New("item is out of stock")

func (inventory *Inventory) Purchase() error {

	if inventory.Available == 0 {
		return ErrOutOfStock
	}

	inventory.Available--
	return nil
}
