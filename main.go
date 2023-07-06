package main

import "fmt"

type Cart struct {
	CartID string
}

type CartStorage interface {
	GetCartByID(cartID string) *Cart
	PutCart(cart *Cart)
}

type CartStorageImpl struct {
}

func (c *CartStorageImpl) GetCartByID(cartID string) *Cart {
	return nil
}

func (c *CartStorageImpl) PutCart(cart *Cart) {

}
func main() {
	cartStorage := &CartStorageImpl{}

	cartID := "exampleID"
	cart := cartStorage.GetCartByID(cartID)

	if cart != nil {
		fmt.Println("cart found")
	} else {
		fmt.Println("cart not found")
	}
	//put cart in the data base
	newCart := &Cart{CartID: "newCartID"}
	cartStorage.PutCart(newCart)
}
