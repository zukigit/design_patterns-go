package main

import (
	"fmt"

	"github.com/zukigit/design_patterns-go/builder"
	"github.com/zukigit/design_patterns-go/singleton/pkg"
)

func main() {
	//singleton
	pkg1 := pkg.GetInstance()
	pkg2 := pkg.GetInstance()

	pkg1.AddOne()
	pkg2.AddOne()

	fmt.Println(pkg1.Count, pkg2.Count)

	//builder
	phone := &builder.Phone{}
	register := &builder.Register{
		BrandName:    "sony",
		Price:        "1000000",
		Product_Type: "Phone",
		Quantity:     100,
		Builder:      phone,
	}

	fmt.Println(register)
	// register.GetProudct()
}
