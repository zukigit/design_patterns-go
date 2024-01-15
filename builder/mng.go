package builder

type Register struct {
	BrandName    string
	Price        string
	Product_Type string
	Quantity     int
	Builder      Builder
}

func (rgstr *Register) Register() {
	rgstr.Builder.SetBrandName(rgstr.BrandName)
	rgstr.Builder.SetPrice(rgstr.Price)
	rgstr.Builder.SetProductType(rgstr.Product_Type)
	rgstr.Builder.SetQuantity(rgstr.Quantity)
}

func (rgstr *Register) GetProudct() Product {
	return rgstr.GetProudct()
}
