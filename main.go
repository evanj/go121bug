package main

import (
	"fmt"

	"github.com/evanj/go121bug/pkg"
)

func main() {
	out := &pkg.Router{}
	fmt.Println(out.Route)
}
