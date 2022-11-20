package main

import (
	// "fmt"
	_ "mod-init/data"
	. "mod-init/fmt"
	"mod-init/store"
	"mod-init/store/cart"

	"github.com/fatih/color"
)

func main() {

	var product *store.Product = &store.Product{
		Name:     "Kayak",
		Category: "Watersports",
	}
	product.SetPrice(400)

	var cart *cart.Cart = &cart.Cart{
		CustomerName: "Giorno",
		Products:     []store.Product{*product},
	}

	color.Green("Name: " + cart.CustomerName)
	color.Cyan("Price: " + ToCurrency(cart.GetTotal()))
}
