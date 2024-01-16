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
	mng := new(builder.Mng)
	bldr := new(builder.Phone)

	bldr.GetInstance("Sony", "1000000", 10)
	mng.Register(bldr)
	fmt.Println(mng.GetProduct())
}
