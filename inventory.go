package inventory

import "errors"

type Item struct {
	name        string
	description string
	icon        int
	count       uint
	maxCount    uint
	kind        int
}

type Inventory []Item

func (inv *Inventory) CountKind(kind int) uint {
	counter := uint(0)

	for i := range *inv {
		if (*inv)[i].kind == kind {
			counter++
		}
	}

	return counter
}

func (inv *Inventory) CountNegativeIcons() uint {
	counter := uint(0)

	for i := range *inv {
		if (*inv)[i].icon < 0 {
			counter++
		}
	}

	return counter
}

func (inven *Inventory) IncreaseCount(index int, amount uint) error {
	inv := *inven

	if index < 0 || index >= len(inv) {
		return errors.New("Inventory - IncreaseCount: Index out of bounds")
	}

	if inv[index].count+amount >= inv[index].maxCount {
		return errors.New("Inventory - IncreaseCount: cannot increase this slot")
	}

	inv[index].count += amount
	return nil
}
