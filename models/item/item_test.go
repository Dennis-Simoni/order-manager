package item_test

import (
	"fmt"
	"order-manager/models/item"
	"testing"
)

const (
	id          = "item-123"
	name        = "Pizza"
	price       = float64(19.990)
	currency    = "GBP"
	newId       = "test-123"
	newName     = "Burger"
	newPrice    = float64(21.89081)
	newCurrency = "USD"
)

var item1 = item.Item{
	ID:           id,
	Name:         name,
	CurrencyCode: currency,
	Price:        price,
}

func TestPrintItem(t *testing.T) {
	expectedListing := fmt.Sprintf("ID: %s\n"+
		"Name: %s\nCurrency Code: %s\nPrice: %.2f\n", item1.ID, item1.Name, item1.CurrencyCode, item1.Price)
	listing := item1.PrintItem()
	if listing != expectedListing {
		t.Errorf("item listing incorrect, got: %v, want: %v", listing, expectedListing)
	}
}

func TestChangeID(t *testing.T) {
	item1.ChangeID(newId)
	if item1.ID != newId {
		t.Errorf("item id incorrect, got: %v, want: %s", item1.ID, newId)
	}
}

func TestChangeName(t *testing.T) {
	item1.ChangeName(newName)
	if item1.Name != newName {
		t.Errorf("item name incorrect, got: %v, want: %s", item1.Name, newName)
	}
}

func TestItemChangePrice(t *testing.T) {

	item1.ChangePrice(newPrice, newCurrency)

	if item1.Price != newPrice {
		t.Errorf("item price incorrect, got: %v, want: %.2f", item1.Price, newPrice)
	}

	if item1.CurrencyCode != newCurrency {
		t.Errorf("item currency incorrect, got: %v, want: %s", item1.CurrencyCode, newCurrency)
	}
}

func BenchmarkChangeItemName(b *testing.B) {
	// run the Change Name function b.N times
	for n := 0; n < b.N; n++ {
		item1.ChangeName("Toast")
	}
}
