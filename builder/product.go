package builder

type Product struct {
	brand_name   string
	price        string
	product_type string
	quantity     int
}

type Builder interface {
	SetBrandName(name string)
	SetPrice(price string)
	SetProductType(pd_type string)
	SetQuantity(quantity int)

	GetProduct() Product
}
