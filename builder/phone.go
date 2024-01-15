package builder

type Phone struct {
	brand_name   string
	price        string
	product_type string
	quantity     int
}

func (phone *Phone) SetBrandName(name string) {
	phone.brand_name = name
}

func (phone *Phone) SetPrice(price string) {
	phone.price = price
}

func (phone *Phone) SetProductType(pd_type string) {
	phone.product_type = pd_type
}

func (phone *Phone) SetQuantity(quantity int) {
	phone.quantity = quantity
}

func (phone *Phone) GetProduct() Product {
	return Product{
		brand_name:   phone.brand_name,
		price:        phone.price,
		product_type: phone.product_type,
		quantity:     phone.quantity,
	}
}
