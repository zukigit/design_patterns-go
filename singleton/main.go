package main

import (
	"fmt"

	"github.com/zukigit/design_patterns-go/pkg"
)

func main() {
	pkg1 := pkg.GetInstance()
	pkg2 := pkg.GetInstance()

	pkg1.AddOne()
	pkg2.AddOne()

	fmt.Println(pkg1.Count, pkg2.Count)
}
