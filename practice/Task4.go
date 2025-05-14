package main

import "fmt"

type Product struct {
	ID       int
	Name     string
	Price    float64
	Quantity int
}

type Inventory struct {
	Products []*Product
}

func (in *Inventory) AddProduct(p *Product) {
	in.Products = append(in.Products, p)
}

func (in *Inventory) GetProductByID(id int) *Product {
	for _, product := range in.Products {
		if product.ID == id {
			return product
		}
	}
	return nil
}

func (in *Inventory) UpdateQuantity(id, newQuantity int) bool {
	for _, product := range in.Products {
		if product.ID == id {
			product.Quantity = newQuantity
			return true
		}
	}
	return false
}

func (in *Inventory) GetTotalValue() float64 {
	var sum float64
	for _, pro := range in.Products {
		sum += pro.Price * float64(pro.Quantity)
	}
	return sum
}

func (in *Inventory) GetLowStockProducts(threshold int) []*Product {
	var lowStock []*Product
	for _, product := range in.Products {
		if product.Quantity < threshold {
			lowStock = append(lowStock, product)
		}
	}
	return lowStock
}

func main() {
	invent := Inventory{}

	invent.AddProduct(&Product{
		ID:       1,
		Name:     "Milk",
		Price:    30,
		Quantity: 50,
	})

	invent.AddProduct(&Product{
		ID:       2,
		Name:     "Apple",
		Price:    20,
		Quantity: 10,
	})

	fmt.Println("Стоимость: ", invent.GetTotalValue())

	product := invent.GetProductByID(2)
	if product != nil {
		fmt.Println("найден товар ", product.Name)
	}

	flag := invent.UpdateQuantity(2, 150)
	if flag {
		fmt.Println("товар обновлен")
	}

	lowStock := invent.GetLowStockProducts(60)
	fmt.Println("Малый остаток:")
	for _, p := range lowStock {
		fmt.Println("-", p.Name, p.Quantity)
	}

}
