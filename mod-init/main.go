package main
import (
	"fmt"
	"mod-init/store"
	. "mod-init/fmt"
	"mod-init/store/cart"
)
func main(){

	var product *store.Product = &store.Product{
		Name: "Kayak",
		Category: "WaterSports",
	}
	product.SetPrice(100)
	 
	var cart *cart.Cart = &cart.Cart{
		CustomerName: "Giorno",
		Products: []store.Product{*product},
	}

	fmt.Println("Name:",cart.CustomerName)
	fmt.Println("Price:",ToCurrency(cart.GetTotal()))
}