package builder

type Mng struct {
	bldr Builder
}

func (mng *Mng) Register(bldr Builder) {
	mng.bldr = bldr
}

func (mng *Mng) GetProduct() Product {
	return mng.bldr.GetProductInfo()
}
