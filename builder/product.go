package builder

type Product struct {
	BrandName   string
	Price       string
	ProductType string
	Quantity    int
}

type Builder interface {
	GetInstance(brand_name string, price string, quantity int)
	GetProductInfo() Product
}
