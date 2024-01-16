package builder

type Phone struct {
	BrandName   string
	Price       string
	ProductType string
	Quantity    int
}

func (phone *Phone) GetInstance(brand_name string, price string, quantity int) {
	phone.BrandName = brand_name
	phone.Price = price
	phone.Quantity = quantity
	phone.ProductType = "Phone"
	// return Product{
	// 	BrandName:   brand_name,
	// 	Price:       price,
	// 	ProductType: "Phone",
	// 	Quantity:    quantity,
	// }
}

func (phone *Phone) GetProductInfo() Product {
	return Product(*phone) //that is fucking cool
}
