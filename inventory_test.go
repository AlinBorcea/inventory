package inventory

import "testing"

func initInventory(inv *Inventory) {
	*inv = make(Inventory, 10)
	for i := 0; i < 10; i++ {
		(*inv)[i] = Item{name: "An item", description: "", icon: i*i - 3*i, count: uint(i), maxCount: uint(64), kind: i}
	}
}

func TestCountKind(t *testing.T) {
	var inv Inventory
	initInventory(&inv)

	counter := inv.CountKind(1)
	if counter != 1 {
		t.Fatalf("The counter is %d not 1", counter)
	}
}

func TestCountNegativeIcons(t *testing.T) {
	var inv Inventory
	initInventory(&inv)

	exp := uint(2)
	act := inv.CountNegativeIcons()
	if act != exp {
		t.Fatalf("expected %d not %d", exp, act)
	}

}

func TestIncreaseCountNegativeIndex(t *testing.T) {
	var inv Inventory
	initInventory(&inv)

	index := -1
	amount := 0
	err := inv.IncreaseCount(index, uint(amount))
	if err == nil {
		t.Fatalf("Err is nill")
	}
}

func TestIncreaseCounBigAmount(t *testing.T) {
	var inv Inventory
	initInventory(&inv)

	index := 0
	amount := 64
	err := inv.IncreaseCount(index, uint(amount))
	if err == nil {
		t.Fatalf("Err is nill")
	}
}
