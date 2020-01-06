package models

import (
	"fmt"
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

var item = Item{
	ID:           id,
	Name:         name,
	CurrencyCode: currency,
	Price:        price,
}

func TestPrintItem(t *testing.T) {
	expectedListing := fmt.Sprintf("ID: %s\n"+
		"Name: %s\nCurrency Code: %s\nPrice: %.2f\n", item.ID, item.Name, item.CurrencyCode, item.Price)
	listing := item.PrintItem()
	if listing != expectedListing {
		t.Errorf("item listing incorrect, got: %v, want: %v", listing, expectedListing)
	}
}

func TestChangeID(t *testing.T) {
	item.ChangeID(newId)
	if item.ID != newId {
		t.Errorf("item id incorrect, got: %v, want: %s", item.ID, newId)
	}
}

func TestChangeName(t *testing.T)  {
	item.ChangeName(newName)
	if item.Name != newName {
		t.Errorf("item name incorrect, got: %v, want: %s", item.Name, newName)
	}
}

func TestItemChangePrice(t *testing.T) {

	item.ChangePrice(newPrice, newCurrency)

	if item.Price != newPrice {
		t.Errorf("item price incorrect, got: %v, want: %.2f", item.Price, newPrice)
	}

	if item.CurrencyCode != newCurrency {
		t.Errorf("item currency incorrect, got: %v, want: %s", item.CurrencyCode, newCurrency)
	}
}

func BenchmarkChangeItemName(b *testing.B) {
	// run the Change Name function b.N times
	for n := 0; n < b.N; n++ {
		item.ChangeName("Toast")
	}
}